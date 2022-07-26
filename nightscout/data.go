package nightscout

import (
	"time"
	"twitch-nightscout/diabetes"
)

type Entry struct {
	ID        string               `json:"_id"`
	Date      time.Time            `json:"dateString"`
	Sgv       diabetes.GlucoseMgdl `json:"sgv"`
	Delta     diabetes.GlucoseMgdl `json:"delta"`
	Direction string               `json:"direction"`
	Type      string               `json:"type"`
}

type Settings struct {
	Title      string
	Units      diabetes.GlucoseUnit
	Thresholds Thresholds
}

type Thresholds struct {
	High         diabetes.GlucoseMgdl
	TargetTop    diabetes.GlucoseMgdl
	TargetBottom diabetes.GlucoseMgdl
	Low          diabetes.GlucoseMgdl
}
