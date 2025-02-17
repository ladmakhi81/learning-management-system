package rolehandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct{}

func NewRoleHandler() RoleHandler {
	return RoleHandler{}
}

func (h RoleHandler) CreateRole(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "done CreateRole"})
}

func (h RoleHandler) GetRoles(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "done GetRoles"})
}

func (h RoleHandler) DeleteRoleById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "done DeleteRoleById"})
}
