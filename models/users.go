package models

import "database/sql"

// User ...
type User struct {
	ID        int64          `json:"id" db:"id"`
	FirstName string         `json:"first_name" db:"first_name"`
	LastName  string         `json:"last_name" db:"last_name"`
	Email     string         `json:"email" db:"email"`
	Password  sql.NullString `json:"-" db:"password"`
	UserType
	AuthenticationMethod
}

// AuthenticationMethod ...
type AuthenticationMethod struct {
	ID   int64
	Type string
}

// UserType ...
type UserType struct {
	ID   int64
	Type string
}

// GetUserByEmail ...
func (c Client) GetUserByEmail(email string) (*User, error) {
	user := User{}
	err := c.DB.Get(&user, "SELECT users.id, users.first_name, users.last_name, users.email, users.password, user_types.id, user_types.type, authentication_methods.id, authentication_methods.type from users inner join user_types on users.user_type = user_types.id inner join authentication_methods on users.auth_method = authentication_methods.id WHERE users.email = $1;", email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
