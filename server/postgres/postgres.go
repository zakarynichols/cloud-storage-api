package postgres

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

// TODO: Would like support for different sql databases syntax like MySQL
type DB struct {
	DB      *sql.DB
	connStr string
	now     func() time.Time
}

func NewDB(connStr string) *DB {
	db := &DB{
		connStr: connStr,
		now:     time.Now,
	}
	return db
}

func (db *DB) Open() error {

	var err error

	if db.DB, err = sql.Open("postgres", db.connStr); err != nil {
		return err
	}

	err = db.DB.Ping()

	if err != nil {
		return err
	}

	return nil
}

func (db *DB) Close() error {
	err := db.DB.Close()

	if err != nil {
		return err
	}

	return nil
}
