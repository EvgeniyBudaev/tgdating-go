package response

import "github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"

type ProfileNavigatorResponseDto struct {
	SessionID string              `json:"sessionId"`
	Location  *entity.PointEntity `json:"location"`
}
