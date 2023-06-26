package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

)

var DB *sql.DB

func ConnectDB() {
	//db, err := sql.Open("mysql", "user:password@/dbname")
	// tambahkan parsetime true supaya parsetime bisa di gunakkan
	db, err := sql.Open("mysql", "root:@/go-lang-db?parseTime=true")
	if err != nil {
		panic(err)
	}

	log.Println("Database Connected")
	DB = db

}
