package response

import "github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/entity"

type NavigatorResponseDto struct {
	TelegramUserId string              `json:"telegramUserId"`
	Location       *entity.PointEntity `json:"location"`
}
