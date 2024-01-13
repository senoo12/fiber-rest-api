package infra

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@/db_go_fiber")
	if err != nil {
		return nil, err
	}

	DB = db

	fmt.Println("Kontol")
	return db, nil
}