package response

type ImageResponseDto struct {
	Id             uint64 `json:"id"`
	TelegramUserId string `json:"telegramUserId"`
	Name           string `json:"name"`
	Url            string `json:"url"`
}
