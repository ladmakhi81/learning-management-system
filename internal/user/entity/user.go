package userentity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	FirstName     string
	LastName      string
	Phone         string
	Password      string
	ProfileImage  string
	IsBlock       bool
	BlockByID     *uint
	BlockDate     *time.Time
	BlockReason   string
	LastLoginDate *time.Time
	RoleID        *uint

	gorm.Model
	Teacher
}
