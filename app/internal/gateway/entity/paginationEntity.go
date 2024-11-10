package entity

type PaginationEntity struct {
	HasPrevious   bool   `json:"hasPrevious"`
	HasNext       bool   `json:"hasNext"`
	Page          uint64 `json:"page"`
	Size          uint64 `json:"size"`
	TotalEntities uint64 `json:"totalEntities"`
	TotalPages    uint64 `json:"totalPages"`
}
