package utils

import "math"

type PaginationOptions struct {
	Page        int64
	ItemPerPage int64
}

type Pagination struct {
	Items      interface{}
	Size       int64
	Total      int64
	TotalPages int64
}

func FormatPagination(items interface{}, size, total int64) *Pagination {
	var totalPages int64 = 1

	if total > 0 && size > 0 {
		totalPages = int64(math.Ceil(float64(total) / float64(size)))
	}

	return &Pagination{
		Items:      items,
		Size:       size,
		Total:      total,
		TotalPages: totalPages,
	}
}

func GetOffsetLimit(pagination *PaginationOptions) (offset, limit int64) {
	offset = -1
	limit = -1

	if pagination != nil {
		if pagination.ItemPerPage > 0 {
			limit = pagination.ItemPerPage

			if pagination.Page > 0 {
				offset = (pagination.Page - 1) * pagination.ItemPerPage
			}
		}
	}

	return offset, limit
}
