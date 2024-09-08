package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/josenaldo/fc-walletcore/internal/database"
	"github.com/josenaldo/fc-walletcore/internal/usecase/create_account"
	"github.com/josenaldo/fc-walletcore/internal/usecase/create_client"
	"github.com/josenaldo/fc-walletcore/internal/usecase/create_transaction"
	"github.com/josenaldo/fc-walletcore/pkg/events"
)

func main() {

	// Open up our database connection. Using a connection string, in this case, to connect to a MySQL database.
	user := "root"
	password := "password"
	host := "localhost"
	port := "3306"
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, password, host, port)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()

	// eventDispatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDb(db)
	accountDb := database.NewAccountDb(db)
	transactionDb := database.NewTransactionDb(db)

	createClientUsecase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUsecase := create_account.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUsecase := create_transaction.NewCreateTransactionUseCase(transactionDb, accountDb, *eventDispatcher)

	println("Wallet Core started services: ")
	println(fmt.Sprintf("CreateClientUseCase: %v", createClientUsecase))
	println(fmt.Sprintf("CreateAccountUseCase: %v", createAccountUsecase))
	println(fmt.Sprintf("CreateTransactionUseCase: %v", createTransactionUsecase))
}
