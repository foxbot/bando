package bando

type Message struct {
	Op int
	Data interface{}
}

// receive
type Identify struct {
	Key string
	Bot bool
	Min int
	Max int
}

type Summons struct {
	Op int
	Data interface{}
}

type StatusResp struct {
	Id string
	Guilds map[int]int
	Voice map[int]int
}

// send

type StatusReq struct {
	Id string
}