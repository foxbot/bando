package bando

import "net"

type Bot struct {
	Lower int
	Upper int
}

type SummonsState struct {
	Origin net.Conn
	Responses int
	Guilds map[int]int
	Voice map[int]int
}

type State struct {
	Bots map[net.Conn]Bot
	Summons map[string]SummonsState
}