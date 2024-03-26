package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {

	clientConn, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}

	defer clientConn.Close()

	for {
		serverConn, err := clientConn.Accept()
		if err != nil {
			panic(err)
		}

		go func(serverConn net.Conn) {
			defer serverConn.Close()

			var msgSize uint32
			err := binary.Read(serverConn, binary.LittleEndian, &msgSize)
			if err != nil {
				panic(err)
			}

			msg := make([]byte, msgSize)
			_, err = serverConn.Read(msg)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(msg))
		}(serverConn)

	}

}
