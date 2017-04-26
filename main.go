package main

import (
	"log"
	"net/http"

	"time"

	"github.com/melonmanchan/mobile-systems-backend/app"
	"github.com/melonmanchan/mobile-systems-backend/aws"
	"github.com/melonmanchan/mobile-systems-backend/config"
	"github.com/melonmanchan/mobile-systems-backend/handlers"
	"github.com/melonmanchan/mobile-systems-backend/middleware"
	"github.com/melonmanchan/mobile-systems-backend/models"
	"github.com/melonmanchan/mobile-systems-backend/notifications"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	// First off, parse configuration variables from the server environment
	config := config.ParseTuteeConfig()

	// Attempt to fetch a database connection
	db, err := models.ConnectToDatabase(config.DatabaseURL)

	fb := notifications.BuildFirebaseClient(config)

	aws := aws.BuildAWSUploader(config.S3Conf)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	log.Println("Performing migrations...")

	// Automatically migrate database models if there are any that are unapplied
	// We use github.com/mattes/migrate
	errs := models.PerformPendingMigrations(config.MigrationsPath, config.DatabaseURL)

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

	// Authentication routes
	authRouter := mux.NewRouter().PathPrefix("/auth").Subrouter().StrictSlash(true)

	handlers.AuthHandler(app, authRouter)

	mainRouter.PathPrefix("/auth").Handler(n.With(
		negroni.Wrap(authRouter),
	))

	// User routes
	userRouter := mux.NewRouter().PathPrefix("/user").Subrouter().StrictSlash(true)

	handlers.UserHandler(app, userRouter)

	mainRouter.PathPrefix("/user").Handler(n.With(
		negroni.HandlerFunc(middleware.CreateResolveUser(app)),
		negroni.Wrap(userRouter),
	))

	// Tutorships routes
	tutorshipRouter := mux.NewRouter().PathPrefix("/tutorship").Subrouter().StrictSlash(false)

	handlers.TutorshipHandler(app, tutorshipRouter)

	mainRouter.PathPrefix("/tutorship").Handler(n.With(
		negroni.HandlerFunc(middleware.CreateResolveUser(app)),
		negroni.Wrap(tutorshipRouter),
	))

	// Subjects routes
	subjectsRouter := mux.NewRouter().PathPrefix("/subject").Subrouter().StrictSlash(false)

	handlers.SubjectHandler(app, subjectsRouter)

	mainRouter.PathPrefix("/subject").Handler(n.With(
		negroni.Wrap(subjectsRouter),
	))

	// Messages routes
	messagesRouter := mux.NewRouter().PathPrefix("/message").Subrouter().StrictSlash(false)

	handlers.MessageHandler(app, messagesRouter)

	mainRouter.PathPrefix("/message").Handler(n.With(
		negroni.HandlerFunc(middleware.CreateResolveUser(app)),
		negroni.Wrap(messagesRouter),
	))

	// Events routes
	eventsRouter := mux.NewRouter().PathPrefix("/event").Subrouter().StrictSlash(false)

	handlers.EventHandler(app, eventsRouter)

	mainRouter.PathPrefix("/event").Handler(n.With(
		negroni.HandlerFunc(middleware.CreateResolveUser(app)),
		negroni.Wrap(eventsRouter),
	))

	mainRouter.NotFoundHandler = http.HandlerFunc(middleware.NotFoundHandler)

	n.UseHandler(mainRouter)

	server := &http.Server{
		Handler:      n,
		Addr:         ":" + config.Port,
		WriteTimeout: 35 * time.Second,
		ReadTimeout:  35 * time.Second,
	}

	log.Println("Server listening on port 8080")
	log.Fatal(server.ListenAndServe())
}
