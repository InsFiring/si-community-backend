package user

import (
	"encoding/json"
)

type MockDBLayer struct {
	err   error
	users []Users
}

func NewMockDBLayer(users []Users) *MockDBLayer {
	return &MockDBLayer{
		users: users,
	}
}

func NewMockDBLayerWithData() *MockDBLayer {
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

	var users []Users
	json.Unmarshal([]byte(USERS), &users)

	return NewMockDBLayer(users)
}

func (mock *MockDBLayer) GetMockUserData() []Users {
	return mock.users
}

func (mock *MockDBLayer) SetError(err error) {
	mock.err = err
}

func (mock *MockDBLayer) AddUser(user Users) (Users, error) {
	if mock.err != nil {
		return Users{}, mock.err
	}

	mock.users = append(mock.users, user)
	return user, nil
}

func (mock *MockDBLayer) SignInUser(userRequestDto UserRequestDto) (UserResponseDto, error) {
	return UserResponseDto{}, nil
}

func (mock *MockDBLayer) ChangePassword(userRequestDto UserRequestDto) (UserResponseDto, error) {
	return UserResponseDto{}, nil
}
