package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s.db", os.Getenv("DATABASE_NAME")))
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
	res, err := db.Exec("CREATE TABLE `users` (`id` INTEGER PRIMARY KEY AUTOINCREMENT,`username` VARCHAR(64) NULL);")
	return res, err
}


func CreateDatabase() error {
	_,err := os.Create(fmt.Sprintf("%s.db", os.Getenv("DATABASE_NAME")))
	return err
}


func DeleteDatabase() error {
	return os.Remove(fmt.Sprintf("%s.db", os.Getenv("DATABASE_NAME")))
}
