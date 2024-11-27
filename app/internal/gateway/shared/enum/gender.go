package enum

type Gender string

const (
	GenderMan   Gender = "man"
	GenderWoman Gender = "woman"
)

func (g Gender) IsValid() bool {
	return g == GenderMan || g == GenderWoman
}
