package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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
	log.Print("Starting Wallet Core")

	// Load environment variables
	err := loadEnv()
	if err != nil {
		log.Fatal("Error loading environment variables" + err.Error())
	}

	// Connect to database
	db, err := connectToDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create unit of work
	uow := createUow(db)

	// Create Kafka producer
	kafkaProducer := createKafkaProducer()

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

func loadEnv() error {
	log.Print("Loading environment variables")

	env := os.Getenv("WALLET_ENV")
	if "" == env {
		env = "development"
	}

	err := godotenv.Load(".env." + env + ".local")
	if err != nil {
		log.Print("Error loading .env." + env + ".local file: " + err.Error())
	}

	if "test" != env {
		err = godotenv.Load(".env.local")
		if err != nil {
			log.Print("Error loading .env.local file: " + err.Error())
		}
	}

	err = godotenv.Load(".env." + env)
	if err != nil {
		log.Print("Error loading .env." + env + " file: " + err.Error())
	}

	err = godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file" + err.Error())
		return err
	} else {
		log.Print("Environment variables loaded")
	}

	return nil
}

func connectToDatabase() (*sql.DB, error) {
	log.Print("Connecting to database")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")

	log.Print("Creating database connection")
	log.Print("- Host: " + host)
	log.Print("- Port: " + port)
	log.Print("- User: " + user)
	log.Printf("- Password: %s", "********")
	log.Print("- Database: " + dbName)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", user, password, host, port, dbName)
	db, err := sql.Open("mysql", connectionString)
	return db, err
}

func createUow(db *sql.DB) *uow.Uow {
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

	return uow
}

func createKafkaProducer() *kafka.Producer {
	log.Print("Creating Kafka producer")

	kafkaHost := os.Getenv("KAFKA_HOST")
	kafkaPort := os.Getenv("KAFKA_PORT")
	groupId := os.Getenv("KAFKA_GROUP_ID")

	log.Print("- Kafka Host: " + kafkaHost)
	log.Print("- Kafka Port: " + kafkaPort)
	log.Print("- Kafka Group ID: " + groupId)

	boostrapServers := kafkaHost + ":" + kafkaPort
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": boostrapServers,
		"group.id":          groupId,
	}
	kafkaProducer := kafka.NewProducer(&configMap)
	return kafkaProducer
}
