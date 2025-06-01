package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env variables:%v", err)
		return
	}
}

func main() {
	e := echo.New()
	// All the endpoints will be defined here
	e.Logger.Fatal(e.Start(os.Getenv("APP_PORT")))
}
