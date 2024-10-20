package sqlite

import (
	"crudoperation/internal/config"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS studets(

	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	age INTEGER,
	email TEXT

	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil

}
