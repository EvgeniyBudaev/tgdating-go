package response

type NavigatorResponseRepositoryDto struct {
	CountryName *string  `json:"countryName"`
	City        *string  `json:"city"`
	Distance    *float64 `json:"distance"`
}
