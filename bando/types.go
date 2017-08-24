package bando

type Message struct {
	Op int `json:"op"`
	Data interface{} `json:"data"`
}

// receive
type Identify struct {
	Key string `json:"key"`
	Bot bool `json:"bot"`
	Min uint `json:"min"`
	Max uint `json:"max"`
}

type Summons struct {
	Op int `json:"op"`
	Data interface{} `json:"data"`
}

type StatusResp struct {
	Id string `json:"id"`
	Guilds map[uint]uint `json:"guilds"`
	Voice map[uint]uint `json:"voice"`
	State uint `json:"state"`
}

// send

type StatusReq struct {
	Id string `json:"id"`
}