package rolehandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	baseerror "github.com/ladmakhi81/learning-management-system/internal/base/error"
	basehandler "github.com/ladmakhi81/learning-management-system/internal/base/handler"
	roleconstant "github.com/ladmakhi81/learning-management-system/internal/role/constant"
	rolecontractor "github.com/ladmakhi81/learning-management-system/internal/role/contractor"
	rolerequestdto "github.com/ladmakhi81/learning-management-system/internal/role/dto/request"
	roleresponsedto "github.com/ladmakhi81/learning-management-system/internal/role/dto/response"
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

func (h RoleHandler) CreateRole(ctx *gin.Context) (*basehandler.Response, error) {
	dto := rolerequestdto.NewCreateRoleReqDTO()
	if err := ctx.Bind(dto); err != nil {
		return nil, baseerror.NewClientErr(roleconstant.INVALID_REQUEST_BODY, http.StatusBadRequest)
	}
	role, roleErr := h.roleSvc.CreateRole(dto)
	if roleErr != nil {
		return nil, roleErr
	}

	mappedRole := h.roleMapper.MapRoleToRoleResponseDTO(role)
	res := roleresponsedto.NewCreateRoleRes(mappedRole)
	return basehandler.NewResponse(res, http.StatusCreated), nil
}

func (h RoleHandler) GetRoles(ctx *gin.Context) (*basehandler.Response, error) {
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
		return nil, rolesErr
	}
	mappedRoles := h.roleMapper.MapRolesToRolesResponseDTO(roles)
	pagination, paginationErr := h.roleSvc.GetRolesPaginationMetadata(uint(page), uint(limit))
	if paginationErr != nil {
		return nil, paginationErr
	}
	res := roleresponsedto.NewGetRolesRes(mappedRoles, *pagination)
	return basehandler.NewResponse(res, http.StatusOK), nil
}

func (h RoleHandler) DeleteRoleById(ctx *gin.Context) (*basehandler.Response, error) {
	roleIdParam := ctx.Param("id")
	roleId, roleIdErr := strconv.Atoi(roleIdParam)
	if roleIdErr != nil {
		return nil, baseerror.NewClientErr(roleconstant.INVALID_ROLE_ID, http.StatusBadRequest)
	}
	if err := h.roleSvc.DeleteRoleById(uint(roleId)); err != nil {
		return nil, err
	}
	res := roleresponsedto.NewDeleteRoleRes("Delete Role Successfully")
	return basehandler.NewResponse(res, http.StatusOK), nil
}
