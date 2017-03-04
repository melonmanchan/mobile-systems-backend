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
	Type string `db:"auth_method_desc"`
}

// UserType ...
type UserType struct {
	ID   int64
	Type string `db:"user_type_desc"`
}

// GetUserByEmail ...
func (c Client) GetUserByEmail(email string) (*User, error) {
	user := User{}
	err := c.DB.Get(&user, `
	SELECT users.id, users.first_name, users.last_name, users.email, users.password,
	user_types.id, user_types.user_type_desc,
	authentication_methods.id, authentication_methods.auth_method_desc
	FROM users
	INNER JOIN user_types ON users.user_type = user_types.id
	INNER JOIN authentication_methods ON users.auth_method = authentication_methods.id
	WHERE users.email = $1;`, email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
