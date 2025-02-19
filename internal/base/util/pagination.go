package baseutil

import "math"

func CalculateTotalPage(totalCount uint, limit uint) uint {
	return uint(math.Ceil(float64(totalCount) / float64(limit)))
}
