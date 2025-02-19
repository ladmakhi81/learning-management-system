package rolehandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	baseerror "github.com/ladmakhi81/learning-management-system/internal/base/error"
	rolecontractor "github.com/ladmakhi81/learning-management-system/internal/role/contractor"
	rolerequestdto "github.com/ladmakhi81/learning-management-system/internal/role/dto/request"
	rolemapper "github.com/ladmakhi81/learning-management-system/internal/role/mapper"
)

type RoleHandler struct {
	roleSvc    rolecontractor.RoleService
	roleMapper rolemapper.RoleMapper
}

func NewRoleHandler(
	roleSvc rolecontractor.RoleService,
	roleMapper rolemapper.RoleMapper,
) RoleHandler {
	return RoleHandler{
		roleSvc:    roleSvc,
		roleMapper: roleMapper,
	}
}

func (h RoleHandler) CreateRole(ctx *gin.Context) error {
	dto := rolerequestdto.NewCreateRoleReqDTO()
	if err := ctx.Bind(dto); err != nil {
		return baseerror.NewClientErr("Invalid Request Body", http.StatusBadRequest)
	}
	role, roleErr := h.roleSvc.CreateRole(dto)
	if roleErr != nil {
		return roleErr
	}
	res := h.roleMapper.MapRoleToRoleResponseDTO(role)
	ctx.JSON(http.StatusCreated, gin.H{"data": res})
	return nil
}

func (h RoleHandler) GetRoles(ctx *gin.Context) error {
	pageParam := ctx.Query("page")
	limitParam := ctx.Query("limit")
	page, pageErr := strconv.Atoi(pageParam)
	if pageErr != nil {
		page = 0
	}
	limit, limitErr := strconv.Atoi(limitParam)
	if limitErr != nil {
		limit = 10
	}
	roles, rolesErr := h.roleSvc.GetRoles(page, limit)
	if rolesErr != nil {
		return rolesErr
	}
	res := h.roleMapper.MapRolesToRolesResponseDTO(roles)
	ctx.JSON(http.StatusOK, gin.H{"data": res})
	return nil
}

func (h RoleHandler) DeleteRoleById(ctx *gin.Context) error {
	roleIdParam := ctx.Param("id")
	roleId, roleIdErr := strconv.Atoi(roleIdParam)
	if roleIdErr != nil {
		return baseerror.NewClientErr("Invalid Role ID", http.StatusBadRequest)
	}
	if err := h.roleSvc.DeleteRoleById(uint(roleId)); err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete Successfully"})
	return nil
}
