package utils

import (
	"../config"
	"github.com/maddevsio/fcm"
)

// SendNotification ...
func SendNotification(cfg config.Config, devices []string, payload fcm.Notification) error {

	c := fcm.NewFCM(cfg.FirebaseServerKey)

	_, err := c.Send(fcm.Message{
		RegistrationIDs:  devices,
		ContentAvailable: true,
		Priority:         fcm.PriorityHigh,
		Notification:     payload,
	})

	return err
}
