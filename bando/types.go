package bando

type message struct {
	op int
	d interface{}
}

// receive
type identify struct {
	key string
}

// send