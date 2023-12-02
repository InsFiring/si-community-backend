package dblayer

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"si-community/models"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const True string = "y"
const False string = "n"
const TokenExpirationHour = 12

type DBORM struct {
	*gorm.DB
}

type Bool struct {
	True string
}

type Claims struct {
	UserID int32 `json:"user_id"`
	jwt.StandardClaims
}

func generateRandomKey(length int) (string, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(key), nil
}

func DBConnection() (*DBORM, error) {
	connection := "test:test1234@tcp(127.0.0.1:3306)/si_community?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	return &DBORM{
		DB: db,
	}, err
}

func (db *DBORM) AddUser(user models.Users) (models.Users, error) {
	fmt.Println("dblayer AddUser")

	var hasUserCount int64

	db.Table("users").
		Where(&models.Users{Email: user.Email}).
		Or(&models.Users{Nickname: user.Nickname}).
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
	return user, db.Omit("register_number").Create(&user).Error
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

func (db *DBORM) addToken(registerNumber int32) (models.Tokens, error) {
	token, err := GenerateToken(registerNumber)
	if err != nil {
		return token, err
	}

	return token, db.Omit("token_id").Create(&token).Error
}

func (db *DBORM) getToken(registerNumber int32) (models.Tokens, error) {
	var token models.Tokens
	result := db.Table("user_tokens").
		Where(&models.Tokens{RegisterNumber: registerNumber}).
		Order("token_id DESC").
		Limit(1).
		Find(&token)

	err := result.Error
	if err != nil {
		return token, err
	}

	// fmt.Printf("now : %s\n", time.Now().Format("2006-01-02 15:04:05"))
	// fmt.Printf("token : %s\n", token.ExpirationTime.Format("2006-01-02 15:04:05"))
	if time.Now().Before(token.ExpirationTime) {
		return token, err
	}

	token, err = db.addToken(registerNumber)
	if err != nil {
		return token, err
	}

	return token, err
}

func GenerateToken(registerNumber int32) (models.Tokens, error) {
	expirationTime := time.Now().Add(TokenExpirationHour * time.Hour)

	claims := &Claims{
		UserID: registerNumber,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey, err := generateRandomKey(32)
	if err != nil {
		return models.Tokens{}, err
	}

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return models.Tokens{}, err
	}

	return models.Tokens{
		RegisterNumber: registerNumber,
		Token:          signedToken,
		ExpirationTime: expirationTime,
	}, nil
}

func (db *DBORM) SignInUser(userRequestDto models.UserRequestDto) (models.UserResponseDto, error) {
	var user models.Users
	var userCount int64

	result := db.Table("Users").Where(&models.Users{Email: userRequestDto.Email}).Find(&user)
	if result.Error != nil {
		return models.UserResponseDto{}, result.Error
	}

	result.Count(&userCount)
	if userCount == 0 {
		return models.UserResponseDto{}, errors.New("User not Founded")
	}

	if !checkPassword(user.Password, userRequestDto.Password) {
		return models.UserResponseDto{}, errors.New("Invalid password")
	}

	token, err := db.getToken(user.RegisterNumber)
	if err != nil {
		return models.UserResponseDto{}, err
	}

	userResponseDto := models.UserResponseDto{
		RegisterNumber: user.RegisterNumber,
		Email:          user.Email,
		Nickname:       user.Nickname,
		Company:        user.Company,
		TokenId:        token.TokenId,
		Token:          token.Token,
		ExpirationTime: token.ExpirationTime,
	}

	return userResponseDto, nil
}

func checkPassword(existingHash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(existingHash), []byte(password)) == nil
}

func (db *DBORM) ChangePassword(userRequestDto models.UserRequestDto) (models.UserResponseDto, error) {
	var user models.Users
	var userCount int64

	result := db.Table("Users").Where(&models.Users{Email: userRequestDto.Email}).Find(&user)
	if result.Error != nil {
		return models.UserResponseDto{}, result.Error
	}

	result.Count(&userCount)
	if userCount == 0 {
		return models.UserResponseDto{}, errors.New("User not Founded")
	}

	hashPassword(&user.Password)

	err := db.Model(&user).
		Where(&models.Users{Email: userRequestDto.Email}).
		Update("password", user.Password).Error
	if err != nil {
		return models.UserResponseDto{}, err
	}

	userResponseDto := models.UserResponseDto{
		RegisterNumber: user.RegisterNumber,
		Email:          user.Email,
		Nickname:       user.Nickname,
		Company:        user.Company,
	}

	return userResponseDto, nil
}

// func (db *DBORM) GetUsersByEmailAndNickname(user models.Users) (models.Users, error) {
// 	fmt.Println("dblayer AddUser")
// 	result := db.Table("Customers").Where(&models.Customer{Email: email})

// 	user.IsActive = True
// 	user.LoggedIn = True
// 	hashPassword(&user.Password)
// 	fmt.Println(user)
// 	return user, db.Omit("register_number").Create(&user).Error
// }
