package response

import "github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/shared/enum"

type SettingsResponseDto struct {
	Measurement enum.Measurement `json:"measurement"`
}
