package bando

const (
	OpcodeIdentify               = 0
	OpcodeAuthenticated          = 1
	OpcodeAuthenticationRejected = 2
	OpcodeSummons                = 3
	OpcodeStatusRequest          = 6
	OpcodeStatusResponse         = 7
	OpcodeStatusAnswer           = 8
	OpcodeRestartAllShards       = 50
	OpcodeRestartSuccess         = 51
	OpcodeRestartFailure         = 52
	OpcodeRollingRestart         = 53
	OpcodeInvalidRequest         = 99
	OpcodeUnknown                = -1
)

type Message struct {
	Op   int         `json:"op"`
	Data interface{} `json:"data"`
}

// receive
type Identify struct {
	Key string `json:"key"`
	Bot bool   `json:"bot"`
	Min uint   `json:"min"`
	Max uint   `json:"max"`
}

type Summons struct {
	Op   int         `json:"op"`
	Data interface{} `json:"data"`
}

type StatusResp struct {
	Id     string        `json:"id"`
	Guilds map[uint]uint `json:"guilds"`
	Voice  map[uint]uint `json:"voice"`
	State  uint          `json:"state"`
}

// send

type StatusReq struct {
	Id string `json:"id"`
}
