package bando

import (
	"encoding/json"
	"log"
	"net"
	"time"
)

var state State

func serve() {
	ln, err := net.Listen("tcp", conf.Host)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close();
	log.Println("Listening on", conf.Host)
	state = State{}
	for {
		sock, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("new conn from", sock.RemoteAddr().String())
		go requestLoop(sock)
	}
}

func requestLoop(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if n == 0 {
			// no data in socket, check later
			time.Sleep(500 * time.Millisecond)
			continue
		}
		if err != nil {
			conn.Close()
			log.Println("read err")
			log.Println(err)
			return
		}
		var m Message
		err = json.Unmarshal(buf[:n], &m)
		if err != nil {
			conn.Close()
			log.Println("unmarshal err ;", string(buf))
			log.Println(err)
			return
		}
		go handleMessage(conn, m)
	}
}
