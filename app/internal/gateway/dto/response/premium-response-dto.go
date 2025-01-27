package response

import "time"

type PremiumResponseDto struct {
	IsPremium      bool      `json:"isPremium"`
	AvailableUntil time.Time `json:"availableUntil"`
}
