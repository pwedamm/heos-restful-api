package groups

type HeosGroupAnswer struct {
	Heos    Heos    `json:"heos"`
	Payload []Group `json:"payload"`
}

type Group struct {
	Name    string   `json:"name"`
	Gid     float32  `json:"gid"`
	Players []Player `json:"players"`
}
