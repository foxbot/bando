package bando

type Message struct {
	Op int `json:"op"`
	Data interface{} `json:"data"`
}

// receive
type Identify struct {
	Key string
}

// send
