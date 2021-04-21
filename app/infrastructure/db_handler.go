package infrastructure

import (
	"database/sql"

	"github.com/Lettuce222/RecipeDealer/interface/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbHandler struct {
	Db *gorm.DB
}

func NewDbHandler() *DbHandler {
	dsn := "host=postgres port=5432 user=postgres password=example dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&database.IngredientRecord{},
		&database.ProcedureRecord{},
	)

	return &DbHandler{Db: db}
}

func (hander *DbHandler) Create(value interface{}) error {
	result := hander.Db.Create(value)
	return result.Error
}

func (hander *DbHandler) Update(identifier uint, value map[string]interface{}, tableName string) error {
	result := hander.Db.Table(tableName).Where("ID = ?", identifier).Updates(value)
	return result.Error
}

func (hander *DbHandler) Delete(identifier uint, tableName string) error {
	result := hander.Db.Table(tableName).Where("id = ?", identifier).Delete(identifier) // Deleteの引数に消されたものが入る？
	return result.Error
}

func (hander *DbHandler) Show(tableName string, columnName string) (database.Rows, error) {
	rows, err := hander.Db.Table(tableName).Select(columnName).Rows()
	return SqlRows{Rows: rows}, err
}

func (hander *DbHandler) Find(identifier uint, tableName string, columnName string) (row database.Row) {
	row = hander.Db.Table(tableName).Where("id = ?", identifier).Select(columnName).Row()
	return
}

type SqlRows struct {
	Rows *sql.Rows
}

func (r SqlRows) Scan(dst ...interface{}) error {
	return r.Rows.Scan(dst...)
}

func (r SqlRows) Next() bool {
	return r.Rows.Next()
}

func (r SqlRows) Close() error {
	return r.Rows.Close()
}
