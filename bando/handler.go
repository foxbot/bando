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
			resp := Message{Op: 1 }
			sendResp(conn, resp)
			log.Println("accepted connection from", conn.RemoteAddr().String())
			state.Bots[conn] = Bot{Lower: i.Min, Upper: i.Max }
		}
	case 3:
		var s Summons
		err := mapstructure.Decode(m.Data, &s)
		if err != nil {
			closeErr(err, conn)
			return
		}
		go doSummons(conn, s)
	case 6:
		var r StatusResp
		err := mapstructure.Decode(m.Data, &r)
		if err != nil {
			closeErr(err, conn)
			return
		}
		handleStatusResp(r)
	default:
		resp := Message { Op: 99 }
		sendResp(conn, resp)
	}
}

func doSummons(conn net.Conn, s Summons) {
	key := randString(10)
	var req interface{}
	switch s.Op{
	case 6:
		req = StatusReq{
			Id: key,
		}
	default:
		return
	}
	state.Summons[key] = SummonsState{Origin:conn}

	for k := range state.Bots{
		m := Message{
			Op: s.Op,
			Data: req,
		}
		sendResp(k, m)
	}
}

func handleStatusResp(resp StatusResp) {
	summons := state.Summons[resp.Id]
	summons.Responses++

	for shard, guilds := range resp.Guilds {
		summons.Guilds[shard] = guilds
	}
	for shard, voice := range resp.Voice {
		summons.Voice[shard] = voice
	}

	if summons.Responses == len(state.Bots) {
		r := StatusResp{
			Guilds: summons.Guilds,
			Voice: summons.Voice,
		}
		m := Message{
			Op: 8,
			Data: r,
		}
		sendResp(summons.Origin, m)
	}
}

func sendResp(conn net.Conn, resp Message) {
	bytes, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		conn.Close()
	}
	_, err = conn.Write([]byte(string(bytes) + "\n"))
	if err != nil {
		log.Println(err)
		conn.Close()
	}
}

func closeErr(e error, conn net.Conn) {
	log.Println(e)
	conn.Close()
}
