package user

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const True string = "y"
const False string = "n"
const TokenExpirationHour = 12

type UserRepository struct {
	db *gorm.DB
}

type Bool struct {
	True string
}

type Claims struct {
	UserID int32 `json:"user_id"`
	jwt.StandardClaims
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func generateRandomKey(length int) (string, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(key), nil
}

func (r *UserRepository) AddUser(user Users) (Users, error) {
	fmt.Println("dblayer AddUser")

	var hasUserCount int64

	r.db.Table("users").
		Where(&Users{Email: user.Email}).
		Or(&Users{Nickname: user.Nickname}).
		Count(&hasUserCount)

	if hasUserCount >= 1 {
		fmt.Println(user)
		fmt.Printf("hasUserCount 값: %v", hasUserCount)
		return user, fmt.Errorf("이미 유저가 있습니다.")
	}

	user.IsActive = True
	user.LoggedIn = True
	hashPassword(&user.Password)
	fmt.Println(user)
	return user, r.db.Omit("register_number").Create(&user).Error
}

func hashPassword(password *string) error {
	if password == nil {
		return errors.New("Reference provided for hashing password is nil")
	}

	passwordBytes := []byte(*password)
	hashedBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	*password = string(hashedBytes[:])
	return nil
}

func (r *UserRepository) SignInUser(userRequestDto UserRequestDto) (UserResponseDto, error) {
	var user Users
	var userCount int64

	result := r.db.Table("users").Where(&Users{Email: userRequestDto.Email}).Find(&user)
	if result.Error != nil {
		return UserResponseDto{}, result.Error
	}

	result.Count(&userCount)
	if userCount == 0 {
		return UserResponseDto{}, errors.New("User not Founded")
	}

	if !checkPassword(user.Password, userRequestDto.Password) {
		return UserResponseDto{}, errors.New("Invalid password")
	}

	token, err := GenerateTokens(user.RegisterNumber)
	if err != nil {
		return UserResponseDto{}, err
	}

	userResponseDto := UserResponseDto{
		RegisterNumber: user.RegisterNumber,
		Email:          user.Email,
		Nickname:       user.Nickname,
		Company:        user.Company,
		AccessToken:    token.AccessToken,
		RefreshToken:   token.RefreshToken,
	}

	return userResponseDto, nil
}

func checkPassword(existingHash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(existingHash), []byte(password)) == nil
}

func (r *UserRepository) ChangePassword(userRequestDto UserRequestDto) (UserResponseDto, error) {
	var user Users
	var userCount int64

	result := r.db.Table("users").Where(&Users{Email: userRequestDto.Email}).Find(&user)
	if result.Error != nil {
		return UserResponseDto{}, result.Error
	}

	result.Count(&userCount)
	if userCount == 0 {
		return UserResponseDto{}, errors.New("User not Founded")
	}

	hashPassword(&user.Password)

	err := r.db.Model(&user).
		Where(&Users{Email: userRequestDto.Email}).
		Update("password", user.Password).Error
	if err != nil {
		return UserResponseDto{}, err
	}

	userResponseDto := UserResponseDto{
		RegisterNumber: user.RegisterNumber,
		Email:          user.Email,
		Nickname:       user.Nickname,
		Company:        user.Company,
	}

	return userResponseDto, nil
}
