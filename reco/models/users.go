package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name            string
	UserDescriptors []UserDescriptor
}

type UserDescriptor struct {
	gorm.Model

	UserID     uint
	Photo      string
	Descriptor []byte
}

func FindAllUsers(db *gorm.DB) (*[]User, error) {
	users := []User{}
	err := db.Preload("UserDescriptors").Find(&users).Error

	if err != nil {
		return &[]User{}, err
	}

	return &users, err
}

func FindUserById(db *gorm.DB, id int) (*User, error) {
	user := User{}
	err := db.Where("id = ?", id).First(&user).Error

	return &user, err
}

func FindDescriptorById(db *gorm.DB, id int) (*UserDescriptor, error) {
	descriptor := UserDescriptor{}
	err := db.Where("id = ?", id).First(&descriptor).Error

	return &descriptor, err
}
