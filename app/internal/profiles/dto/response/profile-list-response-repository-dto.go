package response

import "github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"

type ProfileListResponseRepositoryDto struct {
	*entity.PaginationEntity
	Content []*ProfileListItemResponseRepositoryDto `json:"content"`
}
