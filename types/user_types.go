package types
import "../models"

// DeviceRegisterRequest ...
type DeviceRegisterRequest struct {
	Token string `json:"token"`
}

type UpdateUserRequest struct {
	User models.User `json:"user"`
}

// IsValid ...
func (req DeviceRegisterRequest) IsValid() (bool, []APIError) {
	var errs []APIError

	if req.Token == "" {
		errs = append(errs, RequiredError("Token is required"))
	}

	return len(errs) == 0, errs
}
