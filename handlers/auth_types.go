package handlers

import (
	"time"

	"../models"
)

// LoginRequest ...
type LoginRequest struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse ...
type LoginResponse struct {
	User      *models.User `json:"user"`
	Token     string       `json:"token"`
	ExpiresAt time.Time    `json:"expires_at"`
	CreatedAt time.Time    `json:"created_at"`
}

// GetErrors ...
func (req LoginRequest) GetErrors() []error {
	return nil
}
