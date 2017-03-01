package main

import (
	"log"
	"net/http"
	"time"

	"./handlers"
	"./models"

	"github.com/gorilla/mux"
)

func main() {

	db, err := models.ConnectToDatabase()

	if err != nil {
		log.Fatal(err)
	}

	mainRouter := mux.NewRouter()

	authRouter := mainRouter.PathPrefix("/auth").Subrouter()
	handlers.AuthHandler(db, authRouter)

	defer db.Close()

	server := &http.Server{
		Handler:      mainRouter,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening on port 8080")
	log.Fatal(server.ListenAndServe())
}
