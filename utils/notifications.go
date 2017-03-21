package utils

import (
	"encoding/json"

	"../config"
	"../types"
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

func sendEvent(cfg config.Config, devices []string, data map[string]string) error {

	c := fcm.NewFCM(cfg.FirebaseServerKey)

	_, err := c.Send(fcm.Message{
		RegistrationIDs:  devices,
		ContentAvailable: true,
		Priority:         fcm.PriorityHigh,
		Data:             data,
	})

	return err
}

// SendGeneralMessage ...
func SendGeneralMessage(cfg config.Config, message types.GeneralMessage, devices []string) error {
	encoded, err := json.Marshal(message)

	if err != nil {
		data := make(map[string]string)
		data["type"] = "general"
		data["payload"] = string(encoded)

		err = sendEvent(cfg, devices, data)
	}

	return err
}
