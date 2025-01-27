package enum

type SearchGender string

const (
	SearchGenderAll   SearchGender = "all"
	SearchGenderMan   SearchGender = "man"
	SearchGenderWoman SearchGender = "woman"
)

func (g SearchGender) IsValid() bool {
	return g == SearchGenderAll || g == SearchGenderMan || g == SearchGenderWoman
}
