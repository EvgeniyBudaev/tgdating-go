package response

type TelegramResponseDto struct {
	UserId          string `json:"userId"`
	UserName        string `json:"username"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	LanguageCode    string `json:"languageCode"`
	AllowsWriteToPm bool   `json:"allowsWriteToPm"`
	QueryId         string `json:"queryId"`
}
