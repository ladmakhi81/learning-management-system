package roleservice

import (
	"net/http"

	baseerror "github.com/ladmakhi81/learning-management-system/internal/base/error"
	basetype "github.com/ladmakhi81/learning-management-system/internal/base/type"
	baseutil "github.com/ladmakhi81/learning-management-system/internal/base/util"
	roleconstant "github.com/ladmakhi81/learning-management-system/internal/role/constant"
	rolecontractor "github.com/ladmakhi81/learning-management-system/internal/role/contractor"
	rolerequestdto "github.com/ladmakhi81/learning-management-system/internal/role/dto/request"
	roleentity "github.com/ladmakhi81/learning-management-system/internal/role/entity"
)

type RoleServiceImpl struct {
	roleRepo rolecontractor.RoleRepository
}

func NewRoleServiceImpl(
	roleRepo rolecontractor.RoleRepository,
) RoleServiceImpl {
	return RoleServiceImpl{
		roleRepo: roleRepo,
	}
}

func (svc RoleServiceImpl) CreateRole(createdById uint, dto *rolerequestdto.CreateRoleReqDTO) (*roleentity.Role, error) {
	duplicatedName, duplicatedNameErr := svc.roleRepo.FindRoleByName(dto.Name)
	if duplicatedNameErr != nil {
		return nil, duplicatedNameErr
	}
	if duplicatedName != nil {
		return nil, baseerror.NewClientErr(roleconstant.ROLE_DUPLICATE_NAME, http.StatusConflict)
	}
	role := roleentity.NewRole(
		dto.Name,
		&createdById,
		dto.Permissions,
	)
	if createErr := svc.roleRepo.CreateRole(role); createErr != nil {
		return nil, createErr
	}
	return role, nil
}

func (svc RoleServiceImpl) DeleteRoleById(id uint) error {
	role, roleErr := svc.roleRepo.FindRoleById(id)
	if roleErr != nil {
		return roleErr
	}
	if deleteErr := svc.roleRepo.DeleteRoleById(role.ID); deleteErr != nil {
		return deleteErr
	}
	return nil
}

func (svc RoleServiceImpl) FindRoleById(id uint) (*roleentity.Role, error) {
	role, roleErr := svc.roleRepo.FindRoleById(id)
	if roleErr != nil {
		return nil, roleErr
	}
	if role == nil {
		return nil, baseerror.NewClientErr(
			roleconstant.ROLE_NOT_FOUND_ID,
			http.StatusNotFound,
		)
	}
	return role, nil
}

func (svc RoleServiceImpl) FindRoleByName(name string) (*roleentity.Role, error) {
	role, roleErr := svc.roleRepo.FindRoleByName(name)
	if roleErr != nil {
		return nil, roleErr
	}
	return role, nil
}

func (svc RoleServiceImpl) GetRoles(page, limit int) ([]roleentity.Role, error) {
	roles, rolesErr := svc.roleRepo.GetRoles(page, limit)
	if rolesErr != nil {
		return nil, rolesErr
	}
	return roles, nil
}

func (svc RoleServiceImpl) GetRolesPaginationMetadata(currentPage, limit int) (*basetype.PaginationMetadata, error) {
	totalCount, totalCountErr := svc.roleRepo.GetRolesCount()
	if totalCountErr != nil {
		return nil, totalCountErr
	}
	return basetype.NewPaginationMetadata(
		currentPage,
		baseutil.CalculateTotalPage(totalCount, limit),
		totalCount,
	), nil
}
