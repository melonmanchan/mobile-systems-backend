package main

import (
	"log"
	"net/http"
	"time"

	"./app"
	"./config"
	"./handlers"
	"./middleware"
	"./models"

	"github.com/albrow/negroni-json-recovery"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {

	config := config.ParseTuteeConfig()

	db, err := models.ConnectToDatabase(config.PgConf)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	log.Println("Performing migrations...")

	errs := models.PerformPendingMigrations(config.PgConf)

	if errs != nil {
		log.Fatal(errs)
	}

	log.Println("Migrations performed succesfully!")

	// Set up JSON error response middleware
	recovery.Formatter = func(errMsg string, stack []byte, file string, line int, fullMessages bool) interface{} {
		return map[string]string{
			"error": errMsg,
		}
	}

	n := negroni.New(
		negroni.NewLogger(),
		middleware.SetContentType(),
		recovery.JSONRecovery(true),
	)

	app := app.App{Client: *db, Config: config}

	mainRouter := mux.NewRouter()

	authRouter := mainRouter.PathPrefix("/auth").Subrouter().StrictSlash(true)
	handlers.AuthHandler(app, authRouter)

	mainRouter.NotFoundHandler = http.HandlerFunc(middleware.NotFoundHandler)
	n.UseHandler(mainRouter)

	server := &http.Server{
		Handler:      n,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server listening on port 8080")
	log.Fatal(server.ListenAndServe())
}
