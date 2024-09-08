package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/josenaldo/fc-walletcore/internal/database"
	"github.com/josenaldo/fc-walletcore/internal/usecase/create_account"
	"github.com/josenaldo/fc-walletcore/internal/usecase/create_client"
	"github.com/josenaldo/fc-walletcore/internal/usecase/create_transaction"
	"github.com/josenaldo/fc-walletcore/internal/web"
	"github.com/josenaldo/fc-walletcore/internal/web/webserver"
	"github.com/josenaldo/fc-walletcore/pkg/events"
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

	fmt.Println("Creating event dispatcher")
	eventDispatcher := events.NewEventDispatcher()

	// eventDispatcher.Register("TransactionCreated", handler)

	fmt.Println("Creating use cases")
	clientDb := database.NewClientDb(db)
	accountDb := database.NewAccountDb(db)
	transactionDb := database.NewTransactionDb(db)

	createClientUsecase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUsecase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUsecase := create_transaction.NewCreateTransactionUseCase(transactionDb, accountDb, *eventDispatcher)

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
