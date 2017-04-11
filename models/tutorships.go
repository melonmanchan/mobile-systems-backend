package models

// Tutorship ...
type Tutorship struct {
	ID    int64 `json:"id" db:"id"`
	Tutor User  `json:"tutor" db:"tutor"`
	Tutee User  `json:"tutee" db:"tutee"`
}

// CreateTutoship ...
func (c Client) CreateTutorship(tutor *User, tutee *User) error {
	_, err := c.DB.Exec(`
	INSERT INTO tutorships (tutor_id, tutee_id)
	VALUES($1, $2);
	`, tutor.ID, tutee.ID)

	return err
}
