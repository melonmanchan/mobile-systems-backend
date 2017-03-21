package app

import (
	"../config"
	"../models"
	"../notifications"
)

// App ...
type App struct {
	Client   models.Client
	Config   config.Config
	Firebase notifications.Firebase
}
