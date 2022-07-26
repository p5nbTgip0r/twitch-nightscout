package main

import (
	"flag"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"os"
	"twitch-nightscout/bot"
	"twitch-nightscout/config"
	"twitch-nightscout/core"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yaml", "config file")

	data, err := ioutil.ReadFile(configPath)
	if os.IsNotExist(err) {
		wrote, err := config.WriteExampleConfig(configPath)
		switch {
		case wrote:
			log.Info().Msgf("Wrote example config to '%s'", configPath)
		case err != nil:
			log.Error().Err(err).Msgf("Could not write example config to '%s' for error", configPath)
		case err == nil:
			log.Error().Msgf("Could not write example config to '%s' for an unknown reason", configPath)
		}
		return
	}
	if err != nil {
		log.Err(err).Str("file", configPath).Msg("Reading config file failed")
		return
	}

	cfg, err := config.ReadConfig(data)
	if err != nil {
		log.Err(err).Str("file", configPath).Msg("Parsing config file failed")
		return
	}

	core.LoggingInitialize(&cfg.Log)

	bot.Initialize(cfg)
}
