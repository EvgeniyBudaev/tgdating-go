package enum

type Tariff string

const (
	TariffFree        Tariff = "free"
	TariffMonth       Tariff = "month"
	TariffThreeMonths Tariff = "threeMonths"
	TariffYear        Tariff = "year"
)

func (t Tariff) IsValid() bool {
	return t == TariffFree || t == TariffMonth || t == TariffThreeMonths || t == TariffYear
}
