package bot

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/rs/zerolog/log"
	"strings"
)

func (s *DiabetesSnapshot) placeholders() map[string]string {
	return map[string]string{
		"BG_MGDL":            s.Glucose.Display(),
		"BG_MMOL":            s.Glucose.ToMmol().Display(),
		"DELTA_MGDL":         s.Delta.DisplayDelta(),
		"DELTA_MMOL":         s.Delta.ToMmol().DisplayDelta(),
		"IOB":                s.Iob,
		"COB":                s.Cob,
		"DIRECTION_NAME":     s.Trend.Direction,
		"DIRECTION_ARROW":    s.Trend.Symbol,
		"TIMESTAMP_LOCAL":    s.Timestamp.Format("15:04"),
		"TIMESTAMP_UTC":      s.Timestamp.UTC().Format("15:04"),
		"RELATIVE_TIMESTAMP": humanize.Time(s.Timestamp),
	}
}

func fillPlaceholders(s *DiabetesSnapshot, msg string) string {
	final := msg
	for key, val := range s.placeholders() {
		key = fmt.Sprintf("${%s}", key)
		final = strings.ReplaceAll(final, key, val)
	}

	log.Debug().Str("message", final).Msg("Constructed message with placeholders")
	return final
}
