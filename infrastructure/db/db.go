package db

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB dbの接続保持
var DB *gorm.DB

func Open() {
	var err error

	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")

	connection := "host=localhost port=5432 dbname=sample_database user=" + DBUser + " password=" + DBPass + " sslmode=disable"

	DB, err = gorm.Open("postgres", connection)
	if err != nil {
		panic(err)
	}

	DB.DB().SetMaxIdleConns(10)

	// SetMaxOpenConnsは接続済みのデータベースコネクションの最大数を設定します
	DB.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetimeは再利用され得る最長時間を設定します
	DB.DB().SetConnMaxLifetime(time.Hour)

	fmt.Println("data base ok")
}

func Close() {
	DB.Close()
}
