package types

import "github.com/melonmanchan/mobile-systems-backend/models"

// CreateTutorShipRequest ...
type CreateTutorShipRequest struct {
	TutorID int64 `json:"id"`
}

type TutorshipsResponse struct {
	Tutors []models.User `json:"tutors"`
	Tutees []models.User `json:"tutees"`
}
