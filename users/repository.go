package user

import (
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

func (r *UserRepository) SignInUser(userRequestDto UserRequestDto) (*UserResponseDto, error) {
	var user Users
	var userCount int64

	result := r.db.Table("users").Where(&Users{Email: userRequestDto.Email}).Find(&user)
	if result.Error != nil {
		return &UserResponseDto{}, result.Error
	}

	result.Count(&userCount)
	if userCount == 0 {
		return &UserResponseDto{}, errors.New("User not Founded")
	}

	if !checkPassword(user.Password, userRequestDto.CurrentPassword) {
		return &UserResponseDto{}, errors.New("Invalid password")
	}

	userResponseDto := UserResponseDto{
		RegisterNumber: user.RegisterNumber,
		Email:          user.Email,
		Nickname:       user.Nickname,
		Company:        user.Company,
	}

	return &userResponseDto, nil
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

	if !checkPassword(user.Password, userRequestDto.CurrentPassword) {
		return UserResponseDto{}, errors.New("Invalid password")
	}

	hashPassword(&userRequestDto.NewPassword)

	err := r.db.Model(&user).
		Where(&Users{Email: userRequestDto.Email}).
		Update("password", userRequestDto.NewPassword).Error
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

func (r *UserRepository) HasEmail(userRequestDto UserRequestDto) (bool, error) {
	var user Users
	var userCount int64

	result := r.db.Table("users").Where(&Users{Email: userRequestDto.Email}).Find(&user)
	if result.Error != nil {
		return true, result.Error
	}

	result.Count(&userCount)
	if userCount == 0 {
		return false, nil
	}

	return true, nil
}

func (r *UserRepository) ChangeUserInfo(userModifyDto UserModifyDto) (Users, error) {
	var user Users

	result := r.db.Table("users").Where(&Users{Email: userModifyDto.Email}).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}

	user.Company = userModifyDto.Company
	user.Nickname = userModifyDto.Nickname

	r.db.Updates(user)

	return user, nil
}

func (r *UserRepository) SignOutUser(userSignOutDto UserSignOutDto) (string, error) {
	var user Users
	var count int64

	result := r.db.Table("users").
		Where(&Users{Email: userSignOutDto.Email}).
		Find(&user)

	result.Count(&count)
	err := result.Error
	if err != nil {
		return userSignOutDto.Email, err
	}

	if count == 0 {
		return userSignOutDto.Email, errors.New("회원이 없습니다.")
	}

	return userSignOutDto.Email, r.db.Where(&Users{Email: userSignOutDto.Email}).Delete(&Users{}).Error
}
