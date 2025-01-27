package enum

type Tariff string

const (
	TariffMonth       Tariff = "month"
	TariffThreeMonths Tariff = "threeMonths"
	TariffYear        Tariff = "year"
)

func (t Tariff) IsValid() bool {
	return t == TariffMonth || t == TariffThreeMonths || t == TariffYear
}
