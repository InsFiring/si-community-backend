package dblayer

import (
	"errors"
	"fmt"
	"si-community/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const True string = "y"
const False string = "n"

type DBORM struct {
	*gorm.DB
}

type Bool struct {
	True string
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
	// result := db.Table("users").Where(newUser{Email: user.Email}).Or(newUser{Nickname: user.Nickname}).First(&oldUser)
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

// func (db *DBORM) GetUsersByEmailAndNickname(user models.Users) (models.Users, error) {
// 	fmt.Println("dblayer AddUser")
// 	result := db.Table("Customers").Where(&models.Customer{Email: email})

// 	user.IsActive = True
// 	user.LoggedIn = True
// 	hashPassword(&user.Password)
// 	fmt.Println(user)
// 	return user, db.Omit("register_number").Create(&user).Error
// }
