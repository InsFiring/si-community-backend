package dblayer

import "si-community/models"

type DBlayer interface {
	AddUser(models.Users) (models.Users, error)
	// GetUsersByEmailAndNickname(models.Users) (models.Users, error)
	SignInUser(userRequestDto models.UserRequestDto) (models.UserResponseDto, error)
	// SignOutUser(int) error
}
