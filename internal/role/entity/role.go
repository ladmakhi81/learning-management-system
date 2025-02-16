package roleentity

import (
	"gorm.io/gorm"
)

type Role struct {
	Name        string
	CreatedByID *uint
	Lock        bool
	Permissions []Permission

	gorm.Model
}
