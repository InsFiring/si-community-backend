package dblayer

import (
	"encoding/json"
	"si-community/models"
)

type MockDBLayer struct {
	err    error
	users  []models.Users
	tokens []models.Tokens
}

func NewMockDBLayer(
	users []models.Users,
	tokens []models.Tokens) *MockDBLayer {
	return &MockDBLayer{
		users:  users,
		tokens: tokens,
	}
}

func NewMockDBLayerWithData() *MockDBLayer {
	TOKEN := `[
		{
			"token_id" : 1,
			"register_number" : 1,
			"token" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MDE1Njg3NzF9.Ii4AGCsgmjqGGVqENxe0cSzSn3v-Pb1VieAStvuhxSc",
			"expiration_time" : "2023-12-03 10:59:31",
			"created_at" : "2023-12-02 22:59:31",
			"updated_at" : "2023-12-02 22:59:31",
			"deleted_at" : null
		}
	]
	`
	USERS := `[
		{
			"register_number" : 1,
			"email" : "test@gmail.com",
			"password" : "$2a$10$CCPv9RqbH057GLWOXXJC5.pjVvvVMoK43Pogijzwst4A3YSgP4ZqW",
			"nickname" : "test",
			"company" : "keke",
			"is_active" : "Y",
			"loggedin" : "Y",
			"created_at" : "2023-11-19 15:41:27",
			"updated_at" : "2023-12-02 20:08:33",
			"deleted_at" : null
		},
		{
			"register_number" : 4,
			"email" : "test2@gmail.com",
			"password" : "$2a$10$3Xn2P58cdpHWWq2X3CqOtu8X4RWgr37Gd\/a6Xgn4ekOvz8JfKUs8y",
			"nickname" : "test2",
			"company" : "keke",
			"is_active" : "y",
			"loggedin" : "y",
			"created_at" : "2023-11-19 21:16:15",
			"updated_at" : "2023-11-19 21:16:15",
			"deleted_at" : null
		}
	]
	`

	var users []models.Users
	var tokens []models.Tokens
	json.Unmarshal([]byte(USERS), &users)
	json.Unmarshal([]byte(TOKEN), &tokens)

	return NewMockDBLayer(users, tokens)
}

func (mock *MockDBLayer) GetMockUserData() []models.Users {
	return mock.users
}

func (mock *MockDBLayer) GetMockTokenData() []models.Tokens {
	return mock.tokens
}

func (mock *MockDBLayer) SetError(err error) {
	mock.err = err
}

func (mock *MockDBLayer) AddUser(user models.Users) (models.Users, error) {
	if mock.err != nil {
		return models.Users{}, mock.err
	}

	mock.users = append(mock.users, user)
	return user, nil
}

func (mock *MockDBLayer) SignInUser(userRequestDto models.UserRequestDto) (models.UserResponseDto, error) {
	return models.UserResponseDto{}, nil
}

func (mock *MockDBLayer) ChangePassword(userRequestDto models.UserRequestDto) (models.UserResponseDto, error) {
	return models.UserResponseDto{}, nil
}
