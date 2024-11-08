package database

import (
	"database/sql"

	"github.com/josenaldo/fc-walletcore/internal/entity"
)

type ClientDb struct {
	*BaseRepository
}

func NewClientDb(db *sql.DB) *ClientDb {
	return &ClientDb{
		BaseRepository: &BaseRepository{
			db: db,
		},
	}
}

func NewClientDbWithTx(tx *sql.Tx) *ClientDb {
	return &ClientDb{
		BaseRepository: &BaseRepository{
			tx: tx,
		},
	}
}

// retorna um slice de clientes ou um erro. Se não houver clientes, retorna um slice vazio
func (repo *ClientDb) GetAll() ([]*entity.Client, error) {
	query := "SELECT id, created_at, updated_at, name, email FROM clients"

	stmt, err := repo.prepareQuery(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// criar um slice de clientes e preencher com cada cliente, ou retornar um slice vazio
	clients := make([]*entity.Client, 0)
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		client := &entity.Client{}
		err = rows.Scan(&client.ID, &client.CreatedAt, &client.UpdatedAt, &client.Name, &client.Email)
		if err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	return clients, nil
}

func (repo *ClientDb) Get(id string) (*entity.Client, error) {
	query := "SELECT id, created_at, updated_at, name, email FROM clients WHERE id = ?"

	stmt, err := repo.prepareQuery(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	client := &entity.Client{}
	row := stmt.QueryRow(id)
	err = row.Scan(&client.ID, &client.CreatedAt, &client.UpdatedAt, &client.Name, &client.Email)
	if err != nil {
		return nil, err
	}

	return client, nil

}

func (repo *ClientDb) Save(client *entity.Client) error {
	query := "INSERT INTO clients (id, created_at, updated_at, name, email) VALUES (?, ?, ?, ?, ?)"

	stmt, err := repo.prepareQuery(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(client.ID, client.CreatedAt, client.UpdatedAt, client.Name, client.Email)
	if err != nil {
		return err
	}

	return nil
}
