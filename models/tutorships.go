package models

// Tutorship ...
type Tutorship struct {
	ID    int64 `json:"id" db:"id"`
	Tutor User  `json:"tutor" db:"tutor"`
	Tutee User  `json:"tutee" db:"tutee"`
}

// CreateTutoship ...
func (c Client) CreateTutorship(tutorID int64, tuteeID int64) error {
	_, err := c.DB.Exec(`
	INSERT INTO tutorships (tutor_id, tutee_id)
	VALUES($1, $2);
	`, tutorID, tuteeID)

	return err
}

// GetUserTutorships ...
func (c Client) GetUserTutorships(user *User) ([]Tutorship, error) {
	tutorships := []Tutorship{}
	users := []User{}

	err := c.DB.Select(&users, `
	SELECT users.id, users.first_name, users.last_name, users.email, users.avatar, users.device_tokens,
	users.description,	user_types.id as "user_type.id", user_types.type as "user_type.type",
	authentication_methods.id as "auth_method.id", authentication_methods.type as "auth_method.type"
	FROM users
	INNER JOIN user_types ON users.user_type = user_types.id
	INNER JOIN authentication_methods ON users.auth_method = authentication_methods.id
	WHERE users.id IN (
		SELECT tutorships.tutee_id, tutorships.tutor_id FROM tutorships
		WHERE tutorships.tutee_id = $1 OR tutorships.tutor_id = $1
	);`, user.ID)

	if err != nil {
		return nil, err
	}

	for _, u := range users {
		if user.UserType == TutorType && u.UserType == TuteeType {
			tutorships = append(tutorships, Tutorship{Tutor: *user, Tutee: u})
		} else if user.UserType == TuteeType && u.UserType == TutorType {
			tutorships = append(tutorships, Tutorship{Tutor: u, Tutee: *user})
		}
	}

	return tutorships, nil
}
