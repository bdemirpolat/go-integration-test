package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"os/exec"
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

func CreateDatabase() error {
	cmd := exec.Command("createdb", "integration_test")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func DeleteDatabase() error {
	cmd := exec.Command("dropdb", "-f","integration_test")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
