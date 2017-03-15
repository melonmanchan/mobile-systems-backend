package types

// DeviceRegisterRequest ...
type DeviceRegisterRequest struct {
	Token string `json:"token"`
}

// IsValid ...
func (req DeviceRegisterRequest) IsValid() (bool, []APIError) {
	var errs []APIError

	if req.Token == "" {
		errs = append(errs, RequiredError("Email is required"))
	}

	return len(errs) == 0, errs
}
