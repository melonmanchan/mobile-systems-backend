package types

import (
	"fmt"
	"time"

	"../models"
	v "../validators"
	"github.com/guregu/null"
)

// LoginRequest ...
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterRequest ...
type RegisterRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	UserType  string `json:"user_type"`
}

// LoginResponse ...
type LoginResponse struct {
	User      *models.User `json:"user"`
	Token     string       `json:"token"`
	ExpiresAt time.Time    `json:"expires_at"`
	CreatedAt time.Time    `json:"created_at"`
}

// ToUser ...
func (req RegisterRequest) ToUser() models.User {
	user := models.User{
		ID:                   0,
		FirstName:            req.FirstName,
		LastName:             req.LastName,
		Email:                req.Email,
		Password:             null.StringFrom(req.Password),
		Description:          "",
		AuthenticationMethod: models.NormalAuth,
	}

	if req.UserType == "TUTOR" {
		user.UserType = models.TutorType
	} else if req.UserType == "TUTEE" {
		user.UserType = models.TuteeType
	}

	return user
}

// IsValid ...
func (req LoginRequest) IsValid() (bool, []APIError) {
	var errs []APIError

	if req.Email == "" {
		errs = append(errs, RequiredError("Email is required"))
	} else if !v.IsEmail(req.Email) {
		errs = append(errs, FormatError(fmt.Sprintf("%s is not a valid email address", req.Email)))
	}

	if req.Password == "" {
		errs = append(errs, RequiredError("Password is required"))
	}

	return len(errs) == 0, errs
}

// IsValid ...
func (req RegisterRequest) IsValid() (bool, []APIError) {
	var errs []APIError

	if req.Email == "" {
		errs = append(errs, RequiredError("Email is required"))
	} else if !v.IsEmail(req.Email) {
		errs = append(errs, FormatError(fmt.Sprintf("%s is not a valid email address", req.Email)))
	}

	if req.Password == "" {
		errs = append(errs, RequiredError("Password is required"))
	}

	if req.FirstName == "" {
		errs = append(errs, RequiredError("First name is required"))
	}

	if req.LastName == "" {
		errs = append(errs, RequiredError("Last name is required"))
	}

	if req.UserType == "" {
		errs = append(errs, RequiredError("User type is required"))
	} else if req.UserType != models.TutorType.Type && req.UserType != models.TuteeType.Type {
		errs = append(errs, FormatError(fmt.Sprintf("%s is an unknown user type", req.UserType)))
	}

	return len(errs) == 0, errs
}
