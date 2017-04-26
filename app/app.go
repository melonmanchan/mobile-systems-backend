package app

import (
	"github.com/melonmanchan/mobile-systems-backend/aws"
	"github.com/melonmanchan/mobile-systems-backend/config"
	"github.com/melonmanchan/mobile-systems-backend/models"
	"github.com/melonmanchan/mobile-systems-backend/notifications"
)

// App ...
type App struct {
	Client   models.Client
	Config   config.Config
	Firebase notifications.Firebase
	Uploader aws.Uploader
}
