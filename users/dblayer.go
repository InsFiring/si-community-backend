package user

type DBlayer interface {
	AddUser(Users) (Users, error)
	// GetUsersByEmailAndNickname(models.Users) (models.Users, error)
	SignInUser(userRequestDto UserRequestDto) (UserResponseDto, error)
	ChangePassword(userRequestDto UserRequestDto) (UserResponseDto, error)
}
