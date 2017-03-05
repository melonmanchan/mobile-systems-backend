package models

import "database/sql"

var (
	TutorType = UserType{
		ID:   1,
		Type: "TUTOR",
	}

	TuteeType = UserType{
		ID:   2,
		Type: "TUTEE",
	}

	NormalAuth = AuthenticationMethod{
		ID:   1,
		Type: "BASIC",
	}

	GoogleAuth = AuthenticationMethod{
		ID:   2,
		Type: "GOOGLE_OAUTH",
	}
)

// User ...
type User struct {
	ID                   int64                `json:"-" db:"id"`
	FirstName            string               `json:"first_name" db:"first_name"`
	LastName             string               `json:"last_name" db:"last_name"`
	Email                string               `json:"email" db:"email"`
	Password             sql.NullString       `json:"-" db:"password"`
	UserType             UserType             `json:"user_type" db:"user_type"`
	AuthenticationMethod AuthenticationMethod `json:"auth_method"  db:"auth_method"`
}

// AuthenticationMethod ...
type AuthenticationMethod struct {
	ID   int64  `json:"-" db:"id"`
	Type string `json:"type" db:"type"`
}

// UserType ...
type UserType struct {
	ID   int64  `json:"-" db:"id"`
	Type string `json:"type" db:"type"`
}

// CreateUser ...
func (c Client) CreateUser(user User) error {
	return nil
}

// GetUserByEmail ...
func (c Client) GetUserByEmail(email string, method AuthenticationMethod) (*User, error) {
	user := User{}
	err := c.DB.Get(&user, `
	SELECT users.id , users.first_name, users.last_name, users.email, users.password,
	user_types.id as "user_type.id", user_types.type as "user_type.type",
	authentication_methods.id as "auth_method.id", authentication_methods.type as "auth_method.type"
	FROM users
	INNER JOIN user_types ON users.user_type = user_types.id
	INNER JOIN authentication_methods ON users.auth_method = authentication_methods.id
	WHERE users.email = $1 AND authentication_methods.id = $2;`, email, method.ID)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
