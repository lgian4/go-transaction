package db

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(dbName string) {
	var err error
	DB, err = sql.Open("sqlite3", dbName)
	if err != nil {
		panic(errors.Join(errors.New("could not connect to database"), err))
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	CreateTransactionsTable()
}

func CreateTransactionsTable() {
	createEventTable := `create table if not exists transactions(
		id integer primary key autoincrement,
		date_time datetime not null, 
		description text not null, 
		debit NUMERIC,
		credit NUMERIC,
		current_balance NUMERIC
	)`
	_, err := DB.Exec(createEventTable)
	if err != nil {
		panic(errors.Join(errors.New("could not create transactions table"), err))
	}
}
