package roleresponsedto

import basetype "github.com/ladmakhi81/learning-management-system/internal/base/type"

type GetRolesRes struct {
	Rows []RoleResDTO `json:"rows"`
	basetype.PaginationMetadata
}

func NewGetRolesRes(
	rows []RoleResDTO,
	pagination basetype.PaginationMetadata,
) GetRolesRes {
	return GetRolesRes{
		Rows: rows,
		PaginationMetadata: *basetype.NewPaginationMetadata(
			pagination.CurrentPage+1,
			pagination.TotalPage,
			pagination.TotalCount,
		),
	}
}
