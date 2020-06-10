package main

import (
	"fmt"

	"github.com/Analyse4/jrpc-client/stub"
)

func main() {
	ack, err := stub.Send(":4242", "jrpcserver.ack", "req")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ack)
}
