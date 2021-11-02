package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_NAME"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func CreateTable(db *sql.DB) (sql.Result, error) {
	res, err := db.Exec("create table users(id serial constraint users_pk primary key,username varchar);")
	return res, err
}
