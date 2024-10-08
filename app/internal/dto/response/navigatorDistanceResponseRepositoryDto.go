package response

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"time"
)

type NavigatorDistanceResponseRepositoryDto struct {
	Id        uint64              `json:"id"`
	SessionId string              `json:"sessionId"`
	Location  *entity.PointEntity `json:"location"`
	IsDeleted bool                `json:"isDeleted"`
	CreatedAt time.Time           `json:"createdAt"`
	UpdatedAt time.Time           `json:"updatedAt"`
	Distance  float64             `json:"distance"`
}
