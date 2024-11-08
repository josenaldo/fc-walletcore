package testutils

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

type TestDBInterface interface {
	SetupTestDB(t *testing.T) *sql.DB
	TearDownTestDB(t *testing.T)
}

type TestDb struct {
	Db *sql.DB
}

// SetupTestDB initializes a new in-memory SQLite database and returns the connection.
// It also runs any necessary migrations or schema setups.
func SetupTestDB(t *testing.T) *TestDb {
	t.Helper() // Marks this function as a test helper function

	db, err := sql.Open("sqlite3", "file:dbtest?mode=memory&cache=shared")
	db.SetMaxOpenConns(1)
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}

	// Create tables
	createTableClientSQL := `
        CREATE TABLE clients (
            id VARCHAR(255) PRIMARY KEY,
            created_at DATETIME,
            updated_at DATETIME,
            name VARCHAR(255),
            email VARCHAR(255)
        );
    `

	createTableAccountsSQL := `
        CREATE TABLE accounts (
            id VARCHAR(255) PRIMARY KEY, 
            created_at DATETIME, 
            updated_at DATETIME, 
            balance DECIMAL, 
            client_id VARCHAR(255)
        );
    `

	createTableTransactionsSQL := `
        CREATE TABLE transactions (
            id VARCHAR(255) PRIMARY KEY, 
            created_at DATETIME, 
            account_from_id VARCHAR(255), 
            account_to_id VARCHAR(255),
            amount DECIMAL, 
            FOREIGN KEY(account_from_id) REFERENCES accounts(id),
            FOREIGN KEY(account_to_id) REFERENCES accounts(id)
        );
    `
	_, err = db.Exec(createTableClientSQL)
	if err != nil {
		t.Fatalf("Failed to create tables clients %v", err)
	}

	_, err = db.Exec(createTableAccountsSQL)
	if err != nil {
		t.Fatalf("Failed to create table accounts: %v", err)
	}

	_, err = db.Exec(createTableTransactionsSQL)
	if err != nil {
		t.Fatalf("Failed to create table transactions: %v", err)
	}

	return &TestDb{
		Db: db,
	}
}

// Close closes the database connection.
func (testDb *TestDb) TearDownTestDB() {
	defer testDb.Db.Close()
	testDb.Db.Exec("DROP TABLE transactions")
	testDb.Db.Exec("DROP TABLE accounts")
	testDb.Db.Exec("DROP TABLE clients")
}
