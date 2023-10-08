package utils

func CalculateLimitAndOffset(pageNumber, pageSize int) (limit, offset int) {
	if pageNumber <= 0 {
		pageNumber = 1
	}
	if pageSize <= 0 {
		pageSize = 10 // Default page size if not specified
	}

	offset = (pageNumber - 1) * pageSize
	limit = pageSize

	return limit, offset
}
