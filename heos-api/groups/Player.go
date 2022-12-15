package groups

type Response interface {
	getCommand() string
}

type HeosPlayerAnswer struct {
	Heos    Heos     `json:"heos"`
	Payload []Player `json:"payload"`
}

type HeosPlayerGetNowPlayingResponse struct {
	Heos    Heos         `json:"heos"`
	Payload PlayingMedia `json:"payload"`
}

type PlayingMedia struct {
	Type      string `json:"type"`
	Song      string `json:"song"`
	Station   string `json:"station"`
	Album     string `json:"album"`
	Artist    string `json:"artist"`
	Image_url string `json:"image_url"`
	Mid       string `json:"mid"`
	Qid       string `json:"qid"`
	Sid       int    `json:"sid"`
	Album_id  string `json:"album_id"`
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

func (a HeosPlayerAnswer) getCommand() string {
	return a.Heos.Command
}
