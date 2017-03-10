package types

import (
	"database/sql"
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

// RegisterRequest ...
type RegisterRequest struct {
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Password    string `json:"password"`
	UserType    string `json:"user_type"`
	Description string `json:"description"`
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
		Password:             sql.NullString{String: req.Password, Valid: true},
		Description:          req.Description,
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
func (req LoginRequest) IsValid() (bool, []error) {
	var errs []error

	if req.Email == "" {
		errs = append(errs, fmt.Errorf("email address is required"))
	} else if !v.IsEmail(req.Email) {
		errs = append(errs, fmt.Errorf("%s is not a valid email address", req.Email))
	}

	if req.Password == "" {
		errs = append(errs, fmt.Errorf("password is required"))
	}

	return len(errs) == 0, errs
}

// IsValid ...
func (req RegisterRequest) IsValid() (bool, []error) {
	var errs []error

	if req.Email == "" {
		errs = append(errs, fmt.Errorf("email address is required"))
	} else if !v.IsEmail(req.Email) {
		errs = append(errs, fmt.Errorf("%s is not a valid email address", req.Email))
	}

	if req.Password == "" {
		errs = append(errs, fmt.Errorf("password is required"))
	}

	if req.FirstName == "" {
		errs = append(errs, fmt.Errorf("first name is required"))
	}

	if req.LastName == "" {
		errs = append(errs, fmt.Errorf("last name is required"))
	}

	if req.UserType == "" {
		errs = append(errs, fmt.Errorf("usertype is required"))
	} else if req.UserType != models.TutorType.Type && req.UserType != models.TuteeType.Type {
		errs = append(errs, fmt.Errorf("%s is an unknown user type", req.UserType))
	}

	if req.Description == "" && req.UserType == models.TutorType.Type {
		errs = append(errs, fmt.Errorf("tutors must have description"))
	}

	return len(errs) == 0, errs
}
