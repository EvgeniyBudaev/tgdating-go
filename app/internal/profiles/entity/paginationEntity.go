package entity

type PaginationEntity struct {
	HasPrevious    bool   `json:"hasPrevious"`
	HasNext        bool   `json:"hasNext"`
	Page           uint64 `json:"page"`
	Size           uint64 `json:"size"`
	NumberEntities uint64 `json:"numberEntities"`
	TotalPages     uint64 `json:"totalPages"`
}

func GetPagination(page, size, numberEntities uint64) *PaginationEntity {
	return &PaginationEntity{
		HasPrevious:    page > 1,
		HasNext:        (page * size) < numberEntities,
		Page:           page,
		Size:           size,
		NumberEntities: numberEntities,
		TotalPages:     getTotalPages(size, numberEntities),
	}
}

func getTotalPages(size, numberEntities uint64) uint64 {
	return (numberEntities + size - 1) / size
}
