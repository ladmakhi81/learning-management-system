package basetype

type PaginationMetadata struct {
	CurrentPage uint `json:"currentPage"`
	TotalPage   uint `json:"totalPage"`
	TotalCount  uint `json:"totalCount"`
}

func NewPaginationMetadata(currentPage, totalPage, totalCount uint) *PaginationMetadata {
	return &PaginationMetadata{
		CurrentPage: currentPage,
		TotalPage:   totalPage,
		TotalCount:  totalCount,
	}
}
