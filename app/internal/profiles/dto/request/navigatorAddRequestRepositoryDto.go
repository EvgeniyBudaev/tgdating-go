package request

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"time"
)

type NavigatorAddRequestRepositoryDto struct {
	SessionId string              `json:"sessionId"`
	Location  *entity.PointEntity `json:"location"`
	CreatedAt time.Time           `json:"createdAt"`
	UpdatedAt time.Time           `json:"updatedAt"`
}
