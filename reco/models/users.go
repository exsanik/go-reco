// models/users.go

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
	PhotoPath  string
	Descriptor [512]byte
}
