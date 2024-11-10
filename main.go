package main

import (
	"log"
	"net/http"
	"os"

	"web-lab2/platform/database"
	"web-lab2/platform/middleware"
	"web-lab2/platform/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.New(os.Getenv("DB_CONN_STR"))
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err := db.Setup(); err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}

	csrfMid := middleware.NewCSRFMiddleware("secret123")

	rtr := router.New(db, csrfMid)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe("0.0.0.0:8080", rtr); err != nil {
		log.Fatal(err)
	}
}
