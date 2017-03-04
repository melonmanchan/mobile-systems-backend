package models

import "database/sql"

// User ...
type User struct {
	ID        int64          `json:"id" db:"user_id"`
	FirstName string         `json:"first_name" db:"user_first_name"`
	LastName  string         `json:"last_name" db:"user_last_name"`
	Email     string         `json:"email" db:"user_email"`
	Password  sql.NullString `json:"-" db:"user_password"`
	UserType
	AuthenticationMethod
}

// AuthenticationMethod ...
type AuthenticationMethod struct {
	ID   int64  `json:"-" db:"authentication_method_id"`
	Type string `json:"auth_type" db:"authentication_method_type"`
}

// UserType ...
type UserType struct {
	ID   int64  `json:"-" db:"user_type_id"`
	Type string `json:"user_type" db:"user_type_type"`
}

// GetUserByEmail ...
func (c Client) GetUserByEmail(email string) (*User, error) {
	user := User{}
	err := c.DB.Get(&user, `
	SELECT users.user_id , users.user_first_name, users.user_last_name, users.user_email, users.user_password,
	user_types.user_type_id, user_types.user_type_type,
	authentication_methods.authentication_method_id, authentication_methods.authentication_method_type
	FROM users
	INNER JOIN user_types ON users.user_user_type = user_types.user_type_id
	INNER JOIN authentication_methods ON users.user_auth_method = authentication_methods.authentication_method_id
	WHERE users.user_email = $1;`, email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
