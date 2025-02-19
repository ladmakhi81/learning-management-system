package rolerepository

import (
	baseerror "github.com/ladmakhi81/learning-management-system/internal/base/error"
	basestorage "github.com/ladmakhi81/learning-management-system/internal/base/storage"
	roleentity "github.com/ladmakhi81/learning-management-system/internal/role/entity"
	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	storage *basestorage.Storage
}

func NewRoleRepositoryImpl(
	storage *basestorage.Storage,
) RoleRepositoryImpl {
	return RoleRepositoryImpl{
		storage: storage,
	}
}

func (r RoleRepositoryImpl) CreateRole(role *roleentity.Role) error {
	result := r.storage.DB.Create(role)
	if result.Error != nil || result.RowsAffected == 0 {
		return baseerror.NewServerErr(result.Error.Error(), "RoleRepositoryImpl.CreateRole")
	}
	return nil
}

func (r RoleRepositoryImpl) DeleteRoleById(id uint) error {
	result := r.storage.DB.Delete(&roleentity.Role{}, id)
	if result.Error != nil || result.RowsAffected == 0 {
		return baseerror.NewServerErr(result.Error.Error(), "RoleRepositoryImpl.DeleteRoleById")
	}
	return nil
}

func (r RoleRepositoryImpl) FindRoleById(id uint) (*roleentity.Role, error) {
	role := new(roleentity.Role)
	result := r.storage.DB.First(&role, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, baseerror.NewServerErr(result.Error.Error(), "RoleRepositoryImpl.FindRoleById")
	}
	return role, nil
}

func (r RoleRepositoryImpl) FindRoleByName(name string) (*roleentity.Role, error) {
	role := new(roleentity.Role)
	result := r.storage.DB.Where("name = ?", name).First(&role)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, baseerror.NewServerErr(result.Error.Error(), "RoleRepositoryImpl.FindRoleByName")
	}
	return role, nil
}

func (r RoleRepositoryImpl) GetRoles(page, limit int) ([]roleentity.Role, error) {
	roles := make([]roleentity.Role, 0)
	result := r.storage.DB.Unscoped().Offset(page).Limit(limit).Order("created_at desc").Find(&roles)
	if result.Error != nil {
		return nil, baseerror.NewServerErr(result.Error.Error(), "RoleRepositoryImpl.GetRoles")
	}
	return roles, nil
}
