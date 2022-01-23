package main

import (
	"log"
	"os"

	"go-gorm-test/domain/repository"
	"go-gorm-test/infrastructure/db"
	"go-gorm-test/infrastructure/route"
	"go-gorm-test/interface/controllers"
	"go-gorm-test/usecase"
	"go-gorm-test/util"

	"github.com/go-playground/validator/v10"
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
	db.Open()
	defer db.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	// DI
	passwordUtil := util.NewPasswordUtil()
	authUtil := util.NewAuthUtil()
	userRepository := repository.NewUserRepository(db.DB)
	userUseCase := usecase.NewUserUseCase(userRepository, passwordUtil, authUtil)
	userController := controllers.NewUserController(userUseCase)
	userRouter := route.NewUserRouter(userController)
	indexRouter := route.NewIndexRouter(userRouter)

	indexRouter.Routing(e)

	e.Logger.Fatal(e.Start(":" + port))
}
