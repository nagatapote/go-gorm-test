package main

import (
	"context"
	"database/sql"
	"fmt"
	"go-todo-test/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	initDB()
	insert()

	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	engine.Run(":3000")
}

func initDB() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
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

func insert() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")

	dns := "host=localhost port=5432 dbname=sample_database user=" + DBUser + " password=" + DBPass + " sslmode=disable"

	db, err := sql.Open("postgres", dns)

	if err != nil {
		panic(err)
	}

	user := models.User{
		Email:          null.StringFrom("test@example.com"),
		PasswordDigest: null.StringFrom("digested-password"),
	}
	fmt.Printf("before user = %+v\n", user)
	user.Insert(context.Background(), db, boil.Infer())
	fmt.Printf("after user = %+v\n", user)
}
