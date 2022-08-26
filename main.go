package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go-ecommerce/router"
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
	e := echo.New()

	router.RegisterRoutes(e)

	e.Logger.Fatal(e.Start("0.0.0.0:8000"))
}
