package response

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"time"
)

type NavigatorDistanceResponseRepositoryDto struct {
	Id        uint64              `json:"id"`
	SessionId string              `json:"sessionId"`
	Location  *entity.PointEntity `json:"location"`
	CreatedAt time.Time           `json:"createdAt"`
	UpdatedAt time.Time           `json:"updatedAt"`
	Distance  float64             `json:"distance"`
}
