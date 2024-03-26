package main

import (
	"encoding/binary"
	"net"
	"time"
)

func main() {

	serverConn, err := net.DialTimeout("tcp", "localhost:1234", 3*time.Second)
	if err != nil {
		panic(err)
	}
	defer serverConn.Close()

	msg := "Ping!"

	err = binary.Write(serverConn, binary.LittleEndian, uint32(len(msg)))
	if err != nil {
		panic(err)
	}

	_, err = serverConn.Write([]byte(msg))
	if err != nil {
		panic(err)
	}

}