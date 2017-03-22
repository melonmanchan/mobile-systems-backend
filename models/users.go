package models

import (
	"database/sql"
	"database/sql/driver"
	"errors"

	"github.com/lib/pq"

	"golang.org/x/crypto/bcrypt"
)

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
	ID           int64          `json:"id" db:"id"`
	FirstName    string         `json:"first_name" db:"first_name"`
	LastName     string         `json:"last_name" db:"last_name"`
	Email        string         `json:"email" db:"email"`
	Description  string         `json:"description" db:"description"`
	Password     sql.NullString `json:"-" db:"password"`
	DeviceTokens pq.StringArray `json:"-" db:"device_tokens"`

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

func (t UserType) Value() (driver.Value, error) {
	return int64(t.ID), nil
}

func (t AuthenticationMethod) Value() (driver.Value, error) {
	return int64(t.ID), nil
}

// IsPasswordValid ...
func (u User) IsPasswordValid(password string) error {
	if !u.Password.Valid {
		return errors.New("User password is nil")
	}

	return bcrypt.CompareHashAndPassword([]byte(u.Password.String), []byte(password))
}

// AddTokenToUser ...
func (c Client) AddTokenToUser(user *User, token string) error {
	tx := c.DB.MustBegin()

	tx.MustExec(`
		UPDATE users
		SET device_tokens = array_append(device_tokens, $1)
		WHERE users.id = $2;
	`, token, user.ID)

	err := tx.Commit()

	return err
}

// CreateUser ...
func (c Client) CreateUser(user *User) error {

	if user.Password.Valid {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password.String), bcrypt.DefaultCost)

		if err != nil {
			return err
		}

		user.Password.String = string(hashedPassword)
	}

	res, err := c.DB.NamedExec(`
	INSERT INTO users (first_name, last_name, email, password, description, user_type, auth_method)
	VALUES(:first_name, :last_name, :email, :password, :description, :user_type, :auth_method)
	`, user)

	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	user.ID = id

	return nil
}

// GetUserByEmail ...
func (c Client) GetUserByEmail(email string, method AuthenticationMethod) (*User, error) {
	user := User{}

	// ID might not persist in JWT, resolve it this way
	if method.Type == NormalAuth.Type {
		method.ID = NormalAuth.ID
	} else if method.Type == GoogleAuth.Type {
		method.ID = GoogleAuth.ID
	}

	err := c.DB.Get(&user, `
	SELECT users.id, users.first_name, users.last_name, users.email, users.password, users.device_tokens,
	users.description,	user_types.id as "user_type.id", user_types.type as "user_type.type",
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

func (c Client) UpdateUserProfile(user *User) error {
	if user.ID == 0 {
		return errors.New("User is not valid!")
	}

	tx := c.DB.MustBegin()

	tx.MustExec(`
		UPDATE users
		SET first_name, last_name = VALUES(:first_name, :last_name)
		WHERE users.id = :id;
	`, user)

	err := tx.Commit()

	return err

}
