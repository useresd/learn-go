package storage

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	db *sql.DB
}

// NewMySQL connect to mysql and returl back the sql.DB
func NewMySQL() (*MySQL, error) {
	db, err := sql.Open("mysql", "root:password@/ledger")

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySQL{db: db}, nil
}

func (m *MySQL) DB() *sql.DB {
	return m.db
}
