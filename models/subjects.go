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
