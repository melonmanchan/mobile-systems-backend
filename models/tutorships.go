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

func (c Client) GetUserTutors(user *User) ([]User, error) {
	tutors := []User{}

	err := c.DB.Select(&tutors, `
	SELECT users.id, users.first_name, users.last_name, users.email, users.avatar, users.device_tokens,
	users.description, users.price,	user_types.id as "user_type.id", user_types.type as "user_type.type",
	authentication_methods.id as "auth_method.id", authentication_methods.type as "auth_method.type"
	FROM users
	INNER JOIN user_types ON users.user_type = user_types.id
	INNER JOIN authentication_methods ON users.auth_method = authentication_methods.id
	WHERE users.id IN (
		SELECT  tutorships.tutor_id FROM tutorships
		WHERE tutorships.tutee_id = $1
	) AND users.id != $1;`, user.ID)

	if err != nil {
		return nil, err
	}

	return tutors, nil
}

func (c Client) GetUserTutees(user *User) ([]User, error) {
	tutees := []User{}

	err := c.DB.Select(&tutees, `
	SELECT users.id, users.first_name, users.last_name, users.email, users.avatar, users.device_tokens,
	users.description, users.price,	user_types.id as "user_type.id", user_types.type as "user_type.type",
	authentication_methods.id as "auth_method.id", authentication_methods.type as "auth_method.type"
	FROM users
	INNER JOIN user_types ON users.user_type = user_types.id
	INNER JOIN authentication_methods ON users.auth_method = authentication_methods.id
	WHERE users.id IN (
		SELECT tutorships.tutee_id FROM tutorships
		WHERE tutorships.tutor_id = $1
	) AND users.id != $1;`, user.ID)

	if err != nil {
		return nil, err
	}

	return tutees, nil
}
