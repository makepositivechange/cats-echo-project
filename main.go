package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/makepostivechange/cats-echo-project/database"
	"github.com/makepostivechange/cats-echo-project/handler"
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
	db, err := database.MySQLConn(context.Background())
	if err != nil {
		log.Printf("Error while connecting to the database:%v", err)
		return
	}
	_ = handler.Handler{ // Don't forget to declare variable
		DB: db,
	}
	e.Logger.Fatal(e.Start(os.Getenv("APP_PORT")))
}
