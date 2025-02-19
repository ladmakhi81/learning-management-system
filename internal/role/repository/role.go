package rolerepository

import (
	"errors"

	roleentity "github.com/ladmakhi81/learning-management-system/internal/role/entity"
	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	DB *gorm.DB
}

func NewRoleRepositoryImpl(
	DB *gorm.DB,
) RoleRepositoryImpl {
	return RoleRepositoryImpl{
		DB: DB,
	}
}

func (r RoleRepositoryImpl) CreateRole(role *roleentity.Role) error {
	duplicatedName, duplicatedNameErr := r.FindRoleByName(role.Name)
	if duplicatedNameErr != nil {
		return duplicatedNameErr
	}
	if duplicatedName != nil {
		return errors.New("Role Exist With This Name")
	}
	result := r.DB.Create(role)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("Database Can't Create Role With Provided Information")
	}
	return nil
}

func (r RoleRepositoryImpl) DeleteRoleById(id uint) error {
	role, roleErr := r.FindRoleById(id)
	if roleErr != nil {
		return roleErr
	}
	result := r.DB.Delete(&roleentity.Role{}, role.ID)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("Database Can't Delete Role By ID")
	}
	return nil
}

func (r RoleRepositoryImpl) FindRoleById(id uint) (*roleentity.Role, error) {
	role := new(roleentity.Role)
	result := r.DB.First(&role, id)
	if result.Error != nil {
		return nil, errors.New("Database Can't Return The Role With Provided ID")
	}
	if role == nil {
		return nil, errors.New("Role Not Found With This Provided ID")
	}
	return role, nil
}

func (r RoleRepositoryImpl) FindRoleByName(name string) (*roleentity.Role, error) {
	role := new(roleentity.Role)
	result := r.DB.Where("name = ?", name).First(&role)
	if result.Error != nil {
		return nil, errors.New("Database Can't Return The Role With Provided Name")
	}
	if role == nil {
		return nil, errors.New("Role Not Found With This Provided Name")
	}
	return role, nil
}

func (r RoleRepositoryImpl) GetRoles(page, limit int) ([]roleentity.Role, error) {
	roles := make([]roleentity.Role, 0)
	result := r.DB.Offset(page).Limit(limit).Find(&roles)
	if result.Error != nil {
		return nil, errors.New("Database Can't Find All Roles")
	}
	return roles, nil
}
