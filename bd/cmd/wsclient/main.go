package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/gobwas/ws"
)

var (
	dest = flag.String("dest", "127.0.0.1:8080", "addr to connect to")
)

func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", *dest)
	if err != nil {
		log.Fatal(err)
	}

	incoming := make(chan []byte, 100)
	go func() {
		for msg := range incoming {
			fmt.Printf("<- %s\n", msg)
		}
	}()
	go func() {
		for {
			f, err := ws.ReadFrame(conn)
			if err != nil {
				log.Fatal(err)
			}
			incoming <- f.Payload
		}
	}()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		err := ws.WriteFrame(conn, ws.NewTextFrame(s.Bytes()))
		if err != nil {
			log.Fatal(err)
		}
	}
}
