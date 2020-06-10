package stub

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/Analyse4/jrpc-client/protocol"
)

const maxBufferSize = 1024

// Send xxx
func Send(addr, id, msg string) (string, error) {
	bm := new(protocol.BaseMsg)
	bm.ID = id
	bm.Msg = []byte(msg)

	jmsg, err := json.Marshal(bm)
	if err != nil {
		return "", err
	}
	raddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return "", err
	}
	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	n, err := conn.Write(jmsg)
	if err != nil {
		return "", err
	}
	fmt.Printf("%v  datagram written: bytes: %d\n", time.Now(), n)

	buffer := make([]byte, maxBufferSize)
	n, saddr, err := conn.ReadFrom(buffer)
	if err != nil {
		return "", err
	}
	fmt.Printf("%v  datagram received: bytes: %d, from: %s\n", time.Now(), n, saddr.String())

	//var ack string
	//err = json.Unmarshal(buffer, &ack)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}
