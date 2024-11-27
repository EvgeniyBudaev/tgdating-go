package response

import "github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/entity"

type NavigatorResponseDto struct {
	SessionId string              `json:"sessionId"`
	Location  *entity.PointEntity `json:"location"`
}
