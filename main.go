package main

import (
	"log"
	"net/http"
	"os"

	"go-todo-test/infrastructure/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
		envLoad()
	}
	db.Init()
	// insert()

	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	engine.Run(":" + port)
}

// func insert() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		panic(err)
// 	}
// 	DBUser := os.Getenv("DB_USER")
// 	DBPass := os.Getenv("DB_PASS")

// 	dns := "host=localhost port=5432 dbname=sample_database user=" + DBUser + " password=" + DBPass + " sslmode=disable"

// 	db, err := sql.Open("postgres", dns)

// 	if err != nil {
// 		panic(err)
// 	}

// 	user := models.User{
// 		Email:          null.StringFrom("test@example.com"),
// 		PasswordDigest: null.StringFrom("digested-password"),
// 	}
// 	fmt.Printf("before user = %+v\n", user)
// 	user.Insert(context.Background(), db, boil.Infer())
// 	fmt.Printf("after user = %+v\n", user)
// }
