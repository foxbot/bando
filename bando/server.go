package bando

import (
	"encoding/json"
	"log"
	"net"
)

func serve() {
	ln, err := net.Listen("tcp", conf.Host)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close();
	log.Println("Listening on", conf.Host)
	for {
		sock, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go requestLoop(sock)
	}
}

func requestLoop(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			conn.Close()
			log.Println(err)
			return
		}
		var m message
		err = json.Unmarshal(buf, &m)
		if err != nil {
			conn.Close()
			log.Println(err)
			return
		}
		go handleMessage(conn, m)
	}
}