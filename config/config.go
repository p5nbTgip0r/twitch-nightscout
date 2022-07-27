package config

import (
	_ "embed"
	"fmt"
	"github.com/goccy/go-yaml"
	"os"
	"strings"
	"twitch-nightscout/config/schema"
	"twitch-nightscout/nightscout"
)

func ReadConfig(data []byte) (*schema.ConfigFile, error) {
	var config schema.ConfigFile
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	// used for converting channel names to lowercase
	replacement := make(map[string]*schema.Channel)

	for channel, chobj := range config.Channels {
		inst, err := nightscout.GetInstance(chobj.Nightscout.Url, chobj.Nightscout.Token)
		if err != nil {
			return nil, fmt.Errorf("could not parse NS URL for channel '%s': %w", channel, err)
		}
		chobj.NightscoutInstance = inst

		replacement[strings.ToLower(channel)] = chobj
	}

	config.Channels = replacement

	return &config, nil
}

//go:embed config.example.yaml
var example []byte

func WriteExampleConfig(path string) (created bool, err error) {
	_, err = os.Stat(path)
	if err == nil {
		return false, nil
	}
	if !os.IsNotExist(err) {
		return false, err
	}

	err = os.WriteFile(path, example, 0644)
	if err != nil {
		return false, err
	}

	return true, nil
}
