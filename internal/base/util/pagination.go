package baseutil

import (
	"math"
	"strconv"

	basetype "github.com/ladmakhi81/learning-management-system/internal/base/type"
)

func CalculateTotalPage(totalCount, limit int) int {
	return int(math.Ceil(float64(totalCount) / float64(limit)))
}

func ExtraPaginationData(pageParam, limitParam string) basetype.PaginationParam {
	page := 0
	limit := 0
	page, pageErr := strconv.Atoi(pageParam)
	if pageErr != nil {
		page = 0
	}
	limit, limitErr := strconv.Atoi(limitParam)
	if limitErr != nil {
		limit = 10
	}
	return basetype.NewPaginationParam(limit, page)
}
