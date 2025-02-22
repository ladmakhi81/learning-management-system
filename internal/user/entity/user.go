package userentity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	FirstName          string
	LastName           string
	Phone              string
	Password           string
	PasswordChangeBy   *uint
	PasswordChangeDate *time.Time
	ProfileImage       string
	IsBlock            bool
	BlockByID          *uint
	BlockDate          *time.Time
	BlockReason        string
	LastLoginDate      *time.Time
	RoleID             *uint

	gorm.Model
	Teacher
}

func (User) TableName() string {
	return "_users"
}

func NewUser(firstName, lastName, phone, password string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Password:  password,
	}
}
