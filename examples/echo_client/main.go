package main

import (
	"encoding/binary"
	"fmt"
	"github.com/catinred2/link"
)

// This is an echo client demo work with the echo_server.
// usage:
//     go run echo_client/main.go
func main() {
	protocol := link.PacketN(2, binary.BigEndian, 123)

	client, err := link.Dial("tcp", "127.0.0.1:10010", protocol, link.LittleEndian)
	if err != nil {
		panic(err)
	}
	go client.ReadLoop(func(msg link.InBuffer) {
		println("message:", string(msg.Get()))
	})

	for {
		var input string
		if _, err := fmt.Scanf("%s\n", &input); err != nil {
			break
		}
		client.Send(link.Binary(input))
	}

	client.Close(nil)

	println("bye")
}
