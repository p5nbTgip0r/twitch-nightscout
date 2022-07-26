package core

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"twitch-nightscout/config/schema"
)

func LoggingInitialize(cfg *schema.Log) {
	var writers []io.Writer

	if cfg.Console.Enable {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if cfg.File.Enable {
		writers = append(writers, &lumberjack.Logger{
			Filename:   cfg.File.Filename,
			MaxSize:    cfg.File.MaxSize,
			MaxBackups: cfg.File.MaxFiles,
			MaxAge:     cfg.File.MaxAge,
		})
	}

	level, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		log.Panic().Msgf("Invalid log level '%s'", cfg.Level)
	}

	multi := io.MultiWriter(writers...)
	log.Logger = zerolog.New(multi).
		With().
		Timestamp().
		Str(zerolog.TimestampFieldName, "2006-01-02 15:04:05 -0700").
		Logger().
		Level(level)
}
