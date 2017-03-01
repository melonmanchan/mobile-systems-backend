package main

import (
	"log"
	"net/http"
	"time"

	"./models"

	"github.com/gorilla/mux"
)

func main() {
	mainRouter := mux.NewRouter()

	db, err := models.ConnectToDatabase()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	mainRouter.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello"))
	})

	server := &http.Server{
		Handler:      mainRouter,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
