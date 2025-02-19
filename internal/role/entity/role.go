package roleentity

import (
	"gorm.io/gorm"
)

type Role struct {
	Name        string
	CreatedByID *uint
	Lock        bool
	Permissions Permissions `gorm:"type:text[]"`

	gorm.Model
}

func (Role) TableName() string {
	return "_roles"
}

func NewRole(
	name string,
	createdById *uint,
	permissions []Permission,
) *Role {
	return &Role{
		Name:        name,
		CreatedByID: createdById,
		Lock:        false,
		Permissions: permissions,
	}
}
