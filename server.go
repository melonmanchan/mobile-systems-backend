package main

import (
	"log"
	"net/http"
	"time"

	"./config"
	"./handlers"
	"./models"

	"github.com/gorilla/mux"
)

func main() {

	config := config.ParseTuteeConfig()

	db, err := models.ConnectToDatabase(config.PgConf)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	mainRouter := mux.NewRouter()

	authRouter := mainRouter.PathPrefix("/auth").Subrouter()
	handlers.AuthHandler(db, authRouter)

	server := &http.Server{
		Handler:      mainRouter,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening on port 8080")
	log.Fatal(server.ListenAndServe())
}
