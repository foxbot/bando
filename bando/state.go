package bando

import "net"

type Bot struct {
	Lower uint
	Upper uint
}

type SummonsState struct {
	Origin    net.Conn
	Responses int
	Guilds    map[uint]uint
	Voice     map[uint]uint
}

type State struct {
	Bots    map[net.Conn]Bot
	Summons map[string]SummonsState
}

func NewState() State {
	return State{
		Bots:    make(map[net.Conn]Bot),
		Summons: make(map[string]SummonsState),
	}
}
