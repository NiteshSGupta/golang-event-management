package main

import (
	"database/sql"
	"log"
	_ "reast-api-in-gin/docs"
	"reast-api-in-gin/internal/database"
	"reast-api-in-gin/internal/env"

	_ "github.com/joho/godotenv/autoload" // _ in front will initalize else go will copmain that we are not useing it function : . Blank identifier import (_) ,
	_ "github.com/mattn/go-sqlite3"       // this is databse driver for go to talk with real SQLite databse
)

// @title Go Gin Rest API
// @version 1.0
// @description A rest API in Go using Gin framework.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token in the format **Bearer &lt;token&gt;**

// Apply the security definition to your endpoints
// @security BearerAuth

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models := database.NewModels(db)
	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "some-secret-1234567"),
		models:    models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
