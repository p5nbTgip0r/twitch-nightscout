package schema

import "twitch-nightscout/nightscout"

type Channel struct {
	Options            ChannelOptions `json:"options"`
	Nightscout         Nightscout     `json:"nightscout"`
	NightscoutInstance *nightscout.Instance
}

type Nightscout struct {
	Url   string `json:"url"`
	Token string `json:"token"`
}

type ChannelOptions struct {
	Aliases        []string `json:"aliases"`
	ResponseFormat string   `json:"response_format"`
}
