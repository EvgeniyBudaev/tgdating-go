package response

type NavigatorDetailResponseDto struct {
	CountryName *string  `json:"countryName"`
	City        *string  `json:"city"`
	Distance    *float64 `json:"distance"`
}
