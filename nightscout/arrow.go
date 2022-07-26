package nightscout

import "strings"

type Trend struct {
	Direction string
	Symbol    string
}

var (
	Trends = []Trend{
		buildTrend("None", ""),
		buildTrend("DoubleUp", "↑↑"),
		buildTrend("SingleUp", "↑"),
		buildTrend("FortyFiveUp", "↗"),
		buildTrend("Flat", "→"),
		buildTrend("FortyFiveDown", "↘"),
		buildTrend("SingleDown", "↓"),
		buildTrend("DoubleDown", "↓↓"),
		buildTrend("NOT COMPUTABLE", "↮"),
		buildTrend("RATE OUT OF RANGE", "↺"),
	}
)

func buildTrend(direction string, symbol string) Trend {
	return Trend{
		Direction: direction,
		Symbol:    symbol,
	}
}

func GetTrendByDirection(direction string) *Trend {
	for _, trend := range Trends {
		if strings.EqualFold(trend.Direction, direction) {
			return &trend
		}
	}

	return nil
}
