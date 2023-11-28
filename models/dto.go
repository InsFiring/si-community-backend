package models

import "time"

type UserRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponseDto struct {
	RegisterNumber int32     `json:"register_number"`
	Email          string    `json:"email"`
	Nickname       string    `json:"nickname"`
	Company        string    `json:"company"`
	TokenId        int64     `json:"token_id"`
	Token          string    `json:"token"`
	ExpirationTime time.Time `json:"expiration_time"`
}
