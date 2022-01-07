package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")

	db, err := sql.Open(
		"postgres",
		"host=localhost port=5432 dbname=sample_database user="+DBUser+" password="+DBPass+" sslmode=disable")

	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(300 * time.Second)

	fmt.Println("data base ok")

	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	engine.Run(":3000")
}
