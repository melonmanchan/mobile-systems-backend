package app

import (
	"../config"
	"../models"
)

// App ...
type App struct {
	Client models.Client
	Config config.Config
}
