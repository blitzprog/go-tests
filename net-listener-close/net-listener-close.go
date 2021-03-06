package main

import (
	"fmt"
	"net"
)

func main() {
	i := 0

	for {
		listener, err := net.Listen("tcp", ":3000")

		if err != nil {
			fmt.Printf("failure @ iteration %d | %s\n", i, err)
			i++
			continue
		}

		go accept(listener)
		// time.Sleep(100 * time.Millisecond)
		err = listener.Close()

		if err != nil {
			fmt.Printf("close failure @ iteration %d | %s\n", i, err)
		}

		i++
	}
}

func accept(listener net.Listener) {
	for {
		_, err := listener.Accept()

		if err != nil {
			break
		}
	}
}
