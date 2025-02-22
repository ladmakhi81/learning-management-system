package basetype

type PaginationParam struct {
	Limit int
	Page  int
}

func NewPaginationParam(limit, page int) PaginationParam {
	return PaginationParam{
		Limit: limit,
		Page:  page,
	}
}

type PaginationMetadata struct {
	CurrentPage int `json:"currentPage"`
	TotalPage   int `json:"totalPage"`
	TotalCount  int `json:"totalCount"`
}

func NewPaginationMetadata(currentPage, totalPage, totalCount int) *PaginationMetadata {
	return &PaginationMetadata{
		CurrentPage: currentPage,
		TotalPage:   totalPage,
		TotalCount:  totalCount,
	}
}
