package entity

type PaginationEntity struct {
	HasPrevious   bool   `json:"hasPrevious"`
	HasNext       bool   `json:"hasNext"`
	Page          uint64 `json:"page"`
	Size          uint64 `json:"size"`
	TotalEntities uint64 `json:"totalEntities"`
	TotalPages    uint64 `json:"totalPages"`
}

func GetPagination(page, size, totalEntities uint64) *PaginationEntity {
	return &PaginationEntity{
		HasPrevious:   page > 1,
		HasNext:       (page * size) < totalEntities,
		Page:          page,
		Size:          size,
		TotalEntities: totalEntities,
		TotalPages:    getTotalPages(size, totalEntities),
	}
}

func getTotalPages(size, totalEntities uint64) uint64 {
	return (totalEntities + size - 1) / size
}
