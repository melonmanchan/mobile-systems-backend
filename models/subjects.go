package models

// Subject ...
type Subject struct {
	ID   int64  `json:"id" db:"id"`
	Type string `json:"type" db:"type"`
}

// GetSubjects ...
func (c Client) GetSubjects() []Subject {
	subjects := []Subject{}
	c.DB.Select(&subjects, "SELECT * FROM subjects;")
	return subjects
}

// GetUserSubjects ...
func (c Client) GetUserSubjects(user *User) error {
	subjects := []Subject{}

	err := c.DB.Select(&subjects, `
	SELECT subjects.* FROM subjects
	WHERE subjects.id IN (
		SELECT user_to_subject.subject_id FROM user_to_subject
		WHERE user_to_subject.user_id = $1
	);`, user.ID)

	if err != nil {
		return err
	}

	user.Subjects = subjects

	return nil
}
