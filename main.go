package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/makepostivechange/cats-echo-project/database"
	"github.com/makepostivechange/cats-echo-project/handler"
	"github.com/makepostivechange/cats-echo-project/models"
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
	db, err := database.MySQLConn(context.Background(), os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("DATABASE_NAME"))
	if err != nil {
		log.Printf("Error while connecting to the database:%v", err)
		return
	}
	var cat_model models.CatInfo
	err = db.AutoMigrate(&cat_model)
	if err != nil {
		log.Printf("Error while creating table in database:%v", err)
	}
	h := handler.Handler{ // Don't forget to declare variable
		DB: db,
	}
	e.GET("/health", h.HealthCheck)
	e.GET("/cats", h.GetCats)
	e.GET("/cats/breed_name/:breed_name", h.GetCat)
	e.Logger.Fatal(e.Start(os.Getenv("APP_PORT")))
}
