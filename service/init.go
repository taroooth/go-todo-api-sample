package service

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func init() {
	loadErr := godotenv.Load(".env")

	if loadErr != nil {
		fmt.Printf("Failed to load env: %v", loadErr)
	}
	
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER_NAME"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "todo_sample",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
