package utilities

import "math"

func CalculatePageSize(totalRows int64, limit int) int {
	var pagesize int
	if limit == -1 {
		pagesize = 1
	} else {
		pagesize = int(math.Ceil(float64(totalRows) / float64(limit)))
	}
	return pagesize
}

func CalculateOffset(page, limit int) int {
	var offset int
	if limit == -1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}
	return offset
}
