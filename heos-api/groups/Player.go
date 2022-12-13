package groups

var PlayerGroup = "player"
var GetPlayers = "get_players"

type Player struct {
	Gid     string `json:"gid"`
	Ip      string `json:"ip"`
	Lineout string `json:"lineout"`
	Model   string `json:"model"`
	Name    string `json:"name"`
	Network string `json:"network"`
	Pid     string `json:"pid"`
	Serial  string `json:"serial"`
	Version string `json:"version"`
}
