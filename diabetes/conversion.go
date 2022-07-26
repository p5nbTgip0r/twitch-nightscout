package diabetes

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type GlucoseUnit int

const (
	MGDL GlucoseUnit = iota
	MMOL
)

func InterpretGlucoseUnit(str string) (GlucoseUnit, error) {
	switch strings.ToLower(str) {
	case "mg/dl":
	case "mgdl":
	case "mg":
		return MGDL, nil
	case "mmol/l":
	case "mmol":
		return MMOL, nil
	}
	return 0, fmt.Errorf("cannot parse glucose unit %s", str)
}

type GlucoseMgdl float64
type GlucoseMmol float64

const ConversionFactor = 18.0156

func prefixSign(input string) string {
	if !strings.HasPrefix(input, "-") {
		return "+" + input
	}
	return input
}

func formatFloat(input float64) string {
	return strconv.FormatFloat(input, 'f', -1, 64)
}

// MGDL

func MgdlToMmol(mgdl GlucoseMgdl) GlucoseMmol {
	return GlucoseMmol(mgdl / ConversionFactor)
}

func (m GlucoseMgdl) ToMmol() GlucoseMmol {
	return MgdlToMmol(m)
}

func (m GlucoseMgdl) Display() string {
	return formatFloat(math.Round(float64(m)))
}

func (m GlucoseMgdl) DisplayDelta() string {
	out := m.Display()
	return prefixSign(out)
}

// MMOL

func (m GlucoseMmol) ToMgdl() GlucoseMgdl {
	return GlucoseMgdl(m * ConversionFactor)
}

func (m GlucoseMmol) Display() string {
	return formatFloat(math.Round(float64(m*10)) / 10)
}

func (m GlucoseMmol) DisplayDelta() string {
	out := m.Display()
	return prefixSign(out)
}
