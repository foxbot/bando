package bando

import (
	"encoding/json"
	"log"
	"net"

	"github.com/mitchellh/mapstructure"
)

func handleMessage(conn net.Conn, m Message) {
	switch m.Op {
	case 0:
		var i Identify
		err := mapstructure.Decode(m.Data, &i)
		if err != nil {
			log.Println(err)
			conn.Close()
			log.Println("rejected connection from", conn.RemoteAddr().String(), "-", err)
		}
		if i.Key != conf.Key {
			resp := Message { Op: 2 }
			sendResp(conn, resp)
			conn.Close()
			log.Println("rejected connection from", conn.RemoteAddr().String(), "- invalid key")
		} else if i.Key == conf.Key {
			resp := Message { Op: 1 }
			sendResp(conn, resp)
			log.Println("accepted connection from", conn.RemoteAddr().String())
		}
	default:
		resp := Message { Op: 3 }
		sendResp(conn, resp)
	}
}

func sendResp(conn net.Conn, resp Message) {
	json, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		conn.Close()
	}
	_, err = conn.Write(json)
	if err != nil {
		log.Println(err)
		conn.Close()
	}
}
