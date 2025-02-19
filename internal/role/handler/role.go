package rolehandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	rolecontractor "github.com/ladmakhi81/learning-management-system/internal/role/contractor"
	rolerequestdto "github.com/ladmakhi81/learning-management-system/internal/role/dto/request"
)

type RoleHandler struct {
	roleSvc rolecontractor.RoleService
}

func NewRoleHandler(
	roleSvc rolecontractor.RoleService,
) RoleHandler {
	return RoleHandler{
		roleSvc: roleSvc,
	}
}

func (h RoleHandler) CreateRole(ctx *gin.Context) {
	dto := rolerequestdto.NewCreateRoleReqDTO()
	if err := ctx.Bind(dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Body"})
		return
	}
	role, roleErr := h.roleSvc.CreateRole(dto)
	if roleErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": roleErr.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": role})
}

func (h RoleHandler) GetRoles(ctx *gin.Context) {
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
		ctx.JSON(http.StatusBadRequest, gin.H{"message": rolesErr})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": roles})
}

func (h RoleHandler) DeleteRoleById(ctx *gin.Context) {
	roleIdParam := ctx.Param("id")
	roleId, roleIdErr := strconv.Atoi(roleIdParam)
	if roleIdErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Role Id"})
		return
	}
	if err := h.roleSvc.DeleteRoleById(uint(roleId)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete Successfully"})
}
