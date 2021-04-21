package database

import (
	"database/sql"
	"time"
)

type DbHandler interface {
	Create(value interface{}) error
	Update(identifier uint, value map[string]interface{}, tableName string) error
	Delete(identifier uint, tableName string) error

	Show(tableName string, columnName string) (Rows, error)
	Find(identifier uint, tableName string, columnName string) Row
}

type Model struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

type Rows interface {
	Close() error
	Next() bool
	Scan(dest ...interface{}) error
}

type Row interface {
	Scan(...interface{}) error
}
