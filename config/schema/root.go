package schema

type ConfigFile struct {
	Twitch   TwitchLogin         `json:"twitch"`
	Log      Log                 `json:"log"`
	Channels map[string]*Channel `json:"channels"`
}
