package groups

type HeosPlayerAnswer struct {
	Heos    Heos     `json:"heos"`
	Payload []Player `json:"payload"`
}

type Player struct {
	Name    string `json:"name"`
	Pid     int    `json:"pid"`
	Gid     int    `json:"gid"`
	Ip      string `json:"ip"`
	Lineout int    `json:"lineout"`
	Model   string `json:"model"`
	Network string `json:"network"`
	Serial  string `json:"serial"`
	Version string `json:"version"`
	Role    string `json:"role"`
}
