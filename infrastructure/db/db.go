package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/volatiletech/sqlboiler/boil"
)

// DB dbの接続保持
var DB *sql.DB

func Init() {
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")

	dns := "host=localhost port=5432 dbname=sample_database user=" + DBUser + " password=" + DBPass + " sslmode=disable"

	db, err := sql.Open("postgres", dns)

	if err != nil {
		panic(err)
	}

	// connection pool settings
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(300 * time.Second)

	// global connection setting
	boil.SetDB(db)
	boil.DebugMode = true

	fmt.Println("data base ok")
}
