package dblayer

import "si-community/models"

type DBlayer interface {
	AddUser(models.Users) (models.Users, error)
	// GetUsersByEmailAndNickname(models.Users) (models.Users, error)
	// SignInUser(email, password string) (models.Users, error)
	// SignOutUser(int) error
}
