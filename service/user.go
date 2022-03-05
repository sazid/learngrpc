package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// User contains user's information
type User struct {
	Username       string
	HashedPassword string
	Role           string
}

func NewUser(username, password, role string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}

	user := &User{
		Username:       username,
		HashedPassword: string(hashedPassword),
		Role:           role,
	}

	return user, nil
}

func (u *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
	return err == nil
}

// Clone clones the user and returns a new pointer to the clone
func (u *User) Clone() *User {
	return &User{
		Username:       u.Username,
		HashedPassword: u.HashedPassword,
		Role:           u.Role,
	}
}
