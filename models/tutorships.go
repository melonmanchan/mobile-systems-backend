package models

// Tutorship ...
type Tutorship struct {
	ID      int64 `json:"id" db:"id"`
	TutorID int64 `json:"tutor_id" db:"tutor_id"`
	TuteeID int64 `json:"tutee_id" db:"tutee_id"`
}

// CreateTutoship ...
func (c Client) CreateTutorship(tutorID int64, tuteeID int64) (Tutorship, error) {
	_, err := c.DB.Exec(`
	INSERT INTO tutorships (tutor_id, tutee_id)
	VALUES($1, $2);
	`, tutorID, tuteeID)

	return Tutorship{TutorID: tutorID, TuteeID: tuteeID}, err
}
