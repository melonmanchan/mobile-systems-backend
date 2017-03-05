package handlers

import (
	"fmt"
	"time"

	"../models"
	v "../validators"
)

// LoginRequest ...
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse ...
type LoginResponse struct {
	User      *models.User `json:"user"`
	Token     string       `json:"token"`
	ExpiresAt time.Time    `json:"expires_at"`
	CreatedAt time.Time    `json:"created_at"`
}

// IsValid ...
func (req LoginRequest) IsValid() (bool, []error) {
	var errs []error

	if req.Email == "" {
		errs = append(errs, fmt.Errorf("Email address is required!"))
	} else if !v.IsEmail(req.Email) {
		errs = append(errs, fmt.Errorf("%s is not a valid email address!", req.Email))
	}

	if req.Password == "" {
		errs = append(errs, fmt.Errorf("Password is required!"))
	}

	return len(errs) == 0, errs
}
