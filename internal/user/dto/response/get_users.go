package userresponsedto

import basetype "github.com/ladmakhi81/learning-management-system/internal/base/type"

type GetUsersRes struct {
	Rows []UserResDTO `json:"rows"`
	basetype.PaginationMetadata
}

func NewGetUsersRes(
	rows []UserResDTO,
	pagination basetype.PaginationMetadata,
) GetUsersRes {
	return GetUsersRes{
		Rows: rows,
		PaginationMetadata: *basetype.NewPaginationMetadata(
			pagination.CurrentPage+1,
			pagination.TotalPage,
			pagination.TotalCount,
		),
	}
}
