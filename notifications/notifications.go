package notifications

import (
	"encoding/json"

	"../config"
	"../types"
	"github.com/maddevsio/fcm"
)

// Firebase ...
type Firebase struct {
	fb fcm.FCM
}

// BuildFirebaseClient ...
func BuildFirebaseClient(cfg config.Config) Firebase {
	fcm := fcm.NewFCM(cfg.FirebaseServerKey)
	return Firebase{fb: *fcm}
}

// SendNotification ...
func (f Firebase) SendNotification(devices []string, payload fcm.Notification) error {
	_, err := f.fb.Send(fcm.Message{
		RegistrationIDs:  devices,
		ContentAvailable: true,
		Priority:         fcm.PriorityHigh,
		Notification:     payload,
	})

	return err
}

func (f Firebase) sendEvent(devices []string, data map[string]string) error {

	_, err := f.fb.Send(fcm.Message{
		RegistrationIDs:  devices,
		ContentAvailable: true,
		Priority:         fcm.PriorityHigh,
		Data:             data,
	})

	return err
}

// SendGeneralMessage ...
func (f Firebase) SendGeneralMessage(message types.GeneralMessage, devices []string) error {
	encoded, err := json.Marshal(message)

	if err != nil {
		data := make(map[string]string)
		data["type"] = "general"
		data["payload"] = string(encoded)

		err = f.sendEvent(devices, data)
	}

	return err
}