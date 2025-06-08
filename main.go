package main

import (
	"context"
	"log"
	"os"
	"time"

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
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	db, err := database.MySQLConn(ctx, os.
		Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("DATABASE_NAME"))
	if err != nil {
		log.Printf("Error while connecting to the database:%v", err)
		return
	}
	var cat_model models.CatInfo
	err = db.AutoMigrate(&cat_model)
	if err != nil {
		log.Printf("Error while creating table in database:%v", err)
		return
	}
	h := handler.Handler{
		DB: db,
	}
	// All the endpoints will be defined here
	e.GET("/health", h.HealthCheck)
	e.GET("/cats", h.GetCats)
	e.PUT("updatecat/breed_name/:breed_name", h.UpdateCatInfo)
	e.GET("/cats/breed_name/:breed_name", h.GetCat)
	e.POST("/newcat", h.AddNewCatToDB)
	e.DELETE("/deletecat/breed_name/:breed_name", h.RemoveCatFromDB)
	// The below command will start the application
	e.Logger.Fatal(e.Start(os.Getenv("APP_PORT")))
}
