package main

import (
	"log"
	"net/http"
	"time"

	"./app"
	"./aws"
	"./config"
	"./handlers"
	"./middleware"
	"./models"
	"./notifications"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {

	// First off, parse configuration variables from the server environment
	config := config.ParseTuteeConfig()

	// Attempt to fetch a database connection
	db, err := models.ConnectToDatabase(config.PgConf)

	fb := notifications.BuildFirebaseClient(config)

	aws := aws.BuildAWSUploader(config.S3Conf)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	log.Println("Performing migrations...")

	// Automatically migrate database models if there are any that are unapplied
	// We use github.com/mattes/migrate
	errs := models.PerformPendingMigrations(config.MigrationsPath, config.PgConf)

	if errs != nil {
		log.Fatal(errs)
	}

	log.Println("Migrations performed succesfully!")

	// These are common middleware to be used for every route
	n := negroni.New(
		negroni.NewLogger(),
		middleware.SetContentType(),
		middleware.JSONRecovery(),
	)

	// App variable holding the database connection and configuration that we can inject into handy places!
	app := app.App{Client: *db, Config: config, Firebase: fb, Uploader: aws}

	// Instantiating the actual routes
	mainRouter := mux.NewRouter()

	authRouter := mux.NewRouter().PathPrefix("/auth").Subrouter().StrictSlash(true)

	handlers.AuthHandler(app, authRouter)

	mainRouter.PathPrefix("/auth").Handler(n.With(
		negroni.Wrap(authRouter),
	))

	userRouter := mux.NewRouter().PathPrefix("/user").Subrouter().StrictSlash(true)

	handlers.UserHandler(app, userRouter)

	mainRouter.PathPrefix("/user").Handler(n.With(
		negroni.HandlerFunc(middleware.CreateResolveUser(app)),
		negroni.Wrap(userRouter),
	))

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
