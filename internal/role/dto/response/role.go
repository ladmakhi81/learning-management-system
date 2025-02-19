package roleresponsedto

import (
	"time"

	roleentity "github.com/ladmakhi81/learning-management-system/internal/role/entity"
)

type RoleResDTO struct {
	ID          uint                   `json:"id"`
	CreatedAt   *time.Time             `json:"createdAt"`
	UpdatedAt   *time.Time             `json:"updatedAt"`
	Name        string                 `json:"name"`
	CreatedByID *uint                  `json:"createdById"`
	Lock        bool                   `json:"lock"`
	Permissions roleentity.Permissions `json:"permissions"`
}
