package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"urlShortener/database"
	"urlShortener/http_handlers"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	checkEnvVars()

	database.CheckConnection()

	r := mux.NewRouter()
	r.HandleFunc("/", http_handlers.HomeHandler).Methods(http.MethodGet)
	r.HandleFunc("/shortener", http_handlers.LinksHandler).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/shortener/{link}", http_handlers.LinksShowHandler).Methods(http.MethodGet)
	r.HandleFunc("/shortener/{link}", http_handlers.LinksDeleteHandler).Methods(http.MethodDelete)
	r.HandleFunc("/stats/{link}", http_handlers.LinksStatHandler).Methods(http.MethodGet)

	http.Handle("/", r)

	err = http.ListenAndServe(":"+os.Getenv("APP_PORT"), r)
	if err != nil {
		return
	}
}

func checkEnvVars() {
	envVars := []string{
		"APP_PORT",
		"DB_PORT",
		"DB_DATABASE",
		"DB_USERNAME",
		"DB_PASSWORD",
	}
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("%s in ENV not set", envVar)
		}
	}
}
