package response

import "github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"

type NavigatorResponseDto struct {
	Location *entity.PointEntity `json:"location"`
}
