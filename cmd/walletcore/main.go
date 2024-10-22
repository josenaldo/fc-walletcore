package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/josenaldo/fc-walletcore/internal/database"
	"github.com/josenaldo/fc-walletcore/internal/event/handler"
	"github.com/josenaldo/fc-walletcore/internal/usecase/create_account"
	"github.com/josenaldo/fc-walletcore/internal/usecase/create_client"
	"github.com/josenaldo/fc-walletcore/internal/usecase/create_transaction"
	"github.com/josenaldo/fc-walletcore/internal/web"
	"github.com/josenaldo/fc-walletcore/internal/web/webserver"
	"github.com/josenaldo/fc-walletcore/pkg/events"
	"github.com/josenaldo/fc-walletcore/pkg/kafka"
	"github.com/josenaldo/fc-walletcore/pkg/uow"
)

func main() {

	// Open up our database connection. Using a connection string, in this case, to connect to a MySQL database.
	fmt.Println("Starting Wallet Core")

	fmt.Println("Connecting to database")
	user := "root"
	password := "root"
	host := "localhost"
	port := "3306"
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/wallet?parseTime=true", user, password, host, port)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Creating unit of work")
	ctx := context.Background()
	uow := uow.NewUow(ctx, db)
	uow.Register("ClientDB", func(tx *sql.Tx) interface{} {
		return database.NewClientDbWithTx(tx)
	})

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDbWithTx(tx)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDbWithTx(tx)
	})

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9094",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewProducer(&configMap)

	transactionCreatedKafkaHandler := handler.NewTransactionCreatedKafkaHandler(kafkaProducer)

	fmt.Println("Creating event dispatcher")
	eventDispatcher := events.NewEventDispatcher()

	eventDispatcher.Register("TransactionCreated", transactionCreatedKafkaHandler)

	fmt.Println("Creating use cases")
	clientDb := database.NewClientDb(db)
	accountDb := database.NewAccountDb(db)

	createClientUsecase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUsecase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUsecase := create_transaction.NewCreateTransactionUseCase(uow, *eventDispatcher)

	fmt.Println("Creating web server")
	webserver := webserver.NewWebServer(":8000")

	clientHandler := web.NewWebClientHandler(*createClientUsecase)
	accountHandler := web.NewWebAccountHandler(*createAccountUsecase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUsecase)

	fmt.Println("Adding handlers")
	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Starting web server")
	webserver.Start()
}
