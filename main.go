package main

import (
	"log"
	"net/http"
	"os"

	"go-sqlboiler-test/domain/repository"
	"go-sqlboiler-test/infrastructure/db"
	"go-sqlboiler-test/infrastructure/route"
	"go-sqlboiler-test/interface/controllers"
	"go-sqlboiler-test/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userRepository := repository.NewUserRepository(db.DB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userController := controllers.NewUserController(userUseCase)
	userRouter := route.NewUserRouter(userController)
	indexRouter := route.NewIndexRouter(userRouter)

	indexRouter.Routing(e)

	engine.Run(":" + port)
}
