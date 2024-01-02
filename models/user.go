package models

import (
	"errors"
	UserDto "events-api/models/dtos/userDto"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name          string    `json:"name" binding:"required"`
	Email         string    `json:"email" binding:"required"`
	Password      []byte    `json:"password" binding:"required"`
	Events        []Event
	Registrations []Registration
}

var user User

func Save(userdto UserDto.UserDto, db *gorm.DB) (*User, error) {
	var user User

	user.Email = userdto.Email
	user.Name = userdto.Name

	encryptedPassword, err1 := bcrypt.GenerateFromPassword([]byte(userdto.Password), 14)

	if err1 != nil {
		return &User{}, errors.New("error occurred while hashing password")
	}

	user.Password = encryptedPassword
	user.ID = uuid.New()
	err := db.Create(&user)
	if err.Error != nil {
		return &User{}, errors.New("error creating user")
	}

	return &user, nil
}

func CheckIfEmailExists(userdto UserDto.UserDto, db *gorm.DB) bool {
	var user User
	db.Where("email = ?", userdto.Email).First(&user)

	if user.Name != "" {
		return true
	}

	return false
}

func (user *User) PopulateUser(id uuid.UUID, db *gorm.DB) {
	db.Where("id = ?", id).First(&user)
}

func ValidateCredentials(userDto *UserDto.LoginDto, db *gorm.DB) error {
	var user User
	db.Where("email = ?", userDto.Email).First(&user)

	//encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), 14)

	//if err != nil {
	//	return errors.New("Error encrypting password")
	//}

	userDto.Id = user.ID

	err := bcrypt.CompareHashAndPassword(user.Password, []byte(userDto.Password))

	if err != nil {
		return errors.New("Wrong Password")
	}

	return nil

}

func GetAllUserEvents(id uuid.UUID, db *gorm.DB) (User, error) {
	var user User

	result := db.Preload("Events").First(&user, id)

	if result.Error != nil {
		panic(result.Error)
	}

	return user, nil
}
