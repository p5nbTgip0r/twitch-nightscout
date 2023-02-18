package bot

import (
	"errors"
	"fmt"
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/rs/zerolog/log"
	"strings"
	"twitch-nightscout/config/schema"
	"twitch-nightscout/core"
)

var cfg *schema.ConfigFile
var client *twitch.Client

func Initialize(c *schema.ConfigFile) {
	//goland:noinspection GoBoolExpressions
	if core.Version != "" {
		log.Info().Msgf("Initializing bot (%s)..", core.Version)
	} else {
		log.Info().Msgf("Initializing bot..")
	}

	cfg = c
	client = twitch.NewClient(cfg.Twitch.Username, cfg.Twitch.OAuth)

	client.OnPrivateMessage(handleMessage)
	client.OnConnect(func() {
		log.Info().Msg("Connected")
	})

	for c := range cfg.Channels {
		log.Trace().Str("channel", c).Msgf("Joining channel %s", c)
		client.Join(c)
	}

	err := client.Connect()
	if err != nil {
		log.Panic().Err(err).Msg("Twitch connection failed")
	}
}

func handleMessage(message twitch.PrivateMessage) {
	log.Trace().
		Str("message_id", message.ID).
		Str("message_author", message.User.Name).
		Str("message", message.Message).
		Msg("Received Twitch message")
	channel := cfg.Channels[strings.ToLower(message.Channel)]
	if channel == nil {
		return
	}

	for _, alias := range channel.Options.Aliases {
		if strings.EqualFold(message.Message, alias) {
			log.Trace().Str("message_id", message.ID).
				Msg("Message is a valid command")
			goto command
		}
	}
	return

command:

	var reply string
	bg, err := getPebble(channel.NightscoutInstance)
	if err != nil {
		log.Err(err).
			Str("channel", message.Channel).
			Str("author", message.User.Name).
			Str("message", message.Message).
			Msg("Error when retrieving NS data")

		reply = "An error occurred"

		var presErr *PresentableError
		if errors.As(err, &presErr) {
			reply = fmt.Sprintf("%s: %s", reply, presErr.Msg)
		}
	} else {
		reply = fillPlaceholders(bg, channel.Options.ResponseFormat)
	}

	client.Reply(message.Channel, message.ID, reply)
	log.Debug().
		Str("message_id", message.ID).
		Str("response", reply).
		Msgf("Replied to %s", message.User.Name)
}
