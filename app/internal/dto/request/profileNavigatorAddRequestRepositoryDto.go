package request

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"time"
)

type ProfileNavigatorAddRequestRepositoryDto struct {
	SessionID string              `json:"sessionId"`
	Location  *entity.PointEntity `json:"location"`
	IsDeleted bool                `json:"isDeleted"`
	CreatedAt time.Time           `json:"createdAt"`
	UpdatedAt time.Time           `json:"updatedAt"`
}
