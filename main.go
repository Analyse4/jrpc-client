package main

import (
	"encoding/json"
	"fmt"

	"github.com/Analyse4/jrpc-client/protocol"
	"github.com/Analyse4/jrpc-client/stub"
)

func main() {
	ack, err := stub.Send(":4241", "jrpc.simplehandler", "req")
	if err != nil {
		fmt.Println(err)
	}
	sa := new(protocol.SimpleAck)
	err = json.Unmarshal(ack.Msg, sa)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sa.Content)
}
