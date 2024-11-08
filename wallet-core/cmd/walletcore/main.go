package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/josenaldo/fc-walletcore/internal/database"
	"github.com/josenaldo/fc-walletcore/internal/event"
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
	log.Print("Starting Wallet Core")

	log.Print("Connecting to database")
	user := "root"
	password := "root"
	host := "mysql"
	port := "3306"
	dbName := "wallet"
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", user, password, host, port, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	log.Print("Creating unit of work")
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

	log.Print("Creating Kafka producer")
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewProducer(&configMap)

	transactionCreatedKafkaHandler := handler.NewTransactionCreatedKafkaHandler(kafkaProducer)
	updateBalanceKafkaHandler := handler.NewUpdateBalanceKafkaHandler(kafkaProducer)

	log.Print("Creating event dispatcher")
	eventDispatcher := events.NewEventDispatcher()

	eventDispatcher.Register(event.TRANSACTION_CREATED_NAME, transactionCreatedKafkaHandler)
	eventDispatcher.Register(event.BALANCE_UPDATED_NAME, updateBalanceKafkaHandler)

	log.Print("Creating use cases")
	clientDb := database.NewClientDb(db)
	accountDb := database.NewAccountDb(db)

	createClientUsecase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUsecase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUsecase := create_transaction.NewCreateTransactionUseCase(uow, *eventDispatcher)

	log.Print("Creating web server")
	webserver := webserver.NewWebServer(":8080")

	clientHandler := web.NewWebClientHandler(*createClientUsecase)
	accountHandler := web.NewWebAccountHandler(*createAccountUsecase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUsecase)

	log.Print("Adding web handlers")
	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	log.Print("Starting web server")
	webserver.Start()

	log.Print("Web server Stopped")
}
