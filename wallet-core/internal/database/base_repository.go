package database

import (
	"database/sql"
)

type BaseRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (base *BaseRepository) prepareQuery(query string) (*sql.Stmt, error) {
	var stmt *sql.Stmt
	var err error

	if base.tx != nil {
		stmt, err = base.tx.Prepare(query)
	} else {
		stmt, err = base.db.Prepare(query)
	}

	if err != nil {
		return nil, err
	}

	return stmt, nil
}
