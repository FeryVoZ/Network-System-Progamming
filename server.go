package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func(clientConn net.Conn) {
			defer clientConn.Close()

			var msgSize uint32
			err := binary.Read(clientConn, binary.LittleEndian, &msgSize)
			if err != nil {
				panic(err)
			}

			msg := make([]byte, msgSize)
			clientConn.SetReadDeadline(time.Now().Add(20 * time.Second))
			_, err = clientConn.Read(msg)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(msg))
		}(clientConn)

	}

}
