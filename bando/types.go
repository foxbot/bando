package bando

type Message struct {
	Op int
	Data interface{} `json:"d"`
}

// receive
type Identify struct {
	Key string
	Bot bool
	Min uint
	Max uint
}

type Summons struct {
	Op int
	Data interface{}
}

type StatusResp struct {
	Id string
	Guilds map[uint]uint
	Voice map[uint]uint
	State uint
}

// send

type StatusReq struct {
	Id string
}