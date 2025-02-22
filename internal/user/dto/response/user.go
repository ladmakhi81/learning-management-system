package userresponsedto

import (
	"time"

	"gorm.io/gorm"
)

type UserResDTO struct {
	ID        uint           `json:"id"`
	CreatedAt *time.Time     `json:"createdAt"`
	UpdatedAt *time.Time     `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Phone     string         `json:"phone"`
}

func NewUserResDTO() UserResDTO {
	return UserResDTO{}
}
