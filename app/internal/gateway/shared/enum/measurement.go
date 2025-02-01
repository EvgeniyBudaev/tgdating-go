package enum

type Measurement string

const (
	GenderMetric   Measurement = "metric"
	GenderAmerican Measurement = "american"
)

func (m Measurement) IsValid() bool {
	return m == GenderMetric || m == GenderAmerican
}
