package bando

import (
	"encoding/json"
	"log"
	"net"

	"github.com/mitchellh/mapstructure"
)

func handleMessage(conn net.Conn, m message) {
	switch m.op {
	case 0:
		var i identify
		err := mapstructure.Decode(m.d, &i)
		if err != nil {
			log.Println(err)
			conn.Close()
		}
		if i.key != conf.Key {
			resp := message { op: 2 }
			sendResp(conn, resp)
			conn.Close()
		} else if i.key == conf.Key {
			resp := message { op: 1 }
			sendResp(conn, resp)
		}
	default: 
		resp := message { op: 3 }
		sendResp(conn, resp)
	}
}

func sendResp(conn net.Conn, resp message) {
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