package utils

import (
	"log"

	"../config"
	"github.com/NaySoftware/go-fcm"
)

// SendNotification ...
func SendNotification(cfg config.Config, device string, payload fcm.NotificationPayload) error {
	data := map[string]string{
		"msg": "Hello World1",
	}

	log.Println(cfg.FirebaseServerKey)
	log.Println(device)

	c := fcm.NewFcmClient(cfg.FirebaseServerKey)

	c.SetPriority(fcm.Priority_HIGH)

	c.SetNotificationPayload(&payload)

	c.NewFcmMsgTo(device, data)

	_, err := c.Send()

	return err
}
