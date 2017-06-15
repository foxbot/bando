package bando

import "log"

var conf config

func Run() {
	log.Println("bando up")
	conf = loadConfig()
	serve()
	log.Println("bando down")
}