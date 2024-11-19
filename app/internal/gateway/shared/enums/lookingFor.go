package enums

type LookingFor string

const (
	LookingForAll          LookingFor = "all"
	LookingForBusiness     LookingFor = "business"
	LookingForChat         LookingFor = "chat "
	LookingForDates        LookingFor = "dates"
	LookingForFriendship   LookingFor = "friendship"
	LookingForRelationship LookingFor = "relationship"
	LookingForSex          LookingFor = "sex"
)

func (g LookingFor) IsValid() bool {
	return g == LookingForAll || g == LookingForBusiness || g == LookingForChat || g == LookingForDates || g == LookingForFriendship ||
		g == LookingForRelationship || g == LookingForSex
}
