package stub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/Analyse4/jrpc-client/protocol"
)

const maxBufferSize = 1024

// Send xxx
func Send(addr, id, msg string) (*protocol.BaseMsg, error) {
	data := new(protocol.SimpleReq)
	data.Content = msg
	tdata, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	bm := new(protocol.BaseMsg)
	bm.Msg = make([]byte, 0)
	bm.ID = id
	bm.Msg = append(bm.Msg, tdata...)
	fmt.Println(bm)

	jmsg, err := json.Marshal(bm)
	if err != nil {
		return nil, err
	}
	fmt.Println(len(jmsg))
	raddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	n, err := conn.Write(jmsg)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v  datagram written: bytes: %d\n", time.Now(), n)

	buffer := make([]byte, maxBufferSize)
	n, saddr, err := conn.ReadFrom(buffer)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v  datagram received: bytes: %d, from: %s\n", time.Now(), n, saddr.String())

	buffer = bytes.Trim(buffer, "\x00")
	err = json.Unmarshal(buffer, bm)
	if err != nil {
		return nil, err
	}
	return bm, nil
}
