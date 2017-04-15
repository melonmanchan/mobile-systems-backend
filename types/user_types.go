package types

import "../models"

// DeviceRegisterRequest ...
type DeviceRegisterRequest struct {
	Token string `json:"token"`
}

// RegisterTutorExtraRequest ...
type RegisterTutorExtraRequest struct {
	Description string           `json:"description"`
	Price       int64            `json:"price"`
	Subjects    []models.Subject `json:"subjects"`
}

// UpdateUserRequest ...
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

// IsValid ...
func (req RegisterTutorExtraRequest) IsValid() (bool, []APIError) {
	var errs []APIError

	if req.Description == "" {
		errs = append(errs, RequiredError("Description is required"))
	}

	if req.Price < 0 {
		errs = append(errs, RequiredError("Price must not be less than 0"))
	} else if req.Price > 150 {
		errs = append(errs, RequiredError("Price must not be higher than 150"))
	}

	if len(req.Subjects) == 0 {
		errs = append(errs, RequiredError("At least one subject must be selected"))
	}

	return len(errs) == 0, errs
}
