package app

import (
	"../aws"
	"../config"
	"../models"
	"../notifications"
)

// App ...
type App struct {
	Client   models.Client
	Config   config.Config
	Firebase notifications.Firebase
	Uploader aws.Uploader
}
