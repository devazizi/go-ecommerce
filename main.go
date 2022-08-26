package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go-ecommerce/adapter/db"
	"go-ecommerce/router"
	"os"
)

func loadEnvironmentVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
}

func main() {

	fmt.Println("start an ecommerce project ...")

	loadEnvironmentVariables()

	database := db.NewDB(os.Getenv("DB_DSN"))

	e := echo.New()

	router.RegisterRoutes(e, database)

	e.Logger.Fatal(e.Start("0.0.0.0:8000"))
}
