package response

import "github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/entity"

type ProfileListResponseDto struct {
	*entity.PaginationEntity
	Content []*ProfileListItemResponseDto `json:"content"`
}
