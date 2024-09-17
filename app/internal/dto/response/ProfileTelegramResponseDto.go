package response

type ProfileTelegramResponseDto struct {
	SessionID       string `json:"sessionId"`
	UserID          uint64 `json:"userId"`
	Username        string `json:"username"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	LanguageCode    string `json:"languageCode"`
	AllowsWriteToPm bool   `json:"allowsWriteToPm"`
	QueryID         string `json:"queryId"`
	ChatID          uint64 `json:"chatId"`
}
