package main

import (
	"bufio"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		panic(err)
	}

	addr := ln.Addr().String()
	log.Printf("listening on %s", addr)

	processRequest := func(conn net.Conn) {
		r := bufio.NewReader(conn)
		req, err := http.ReadRequest(r)
		if err != nil {
			panic(err)
		}
		log.Printf("got request %q", req.Method)

		w := bufio.NewWriter(conn)
		w.WriteString("HTTP/1.1 500 Whoa!?\r\n\r\n")
		w.Flush()
		conn.Close()
	}

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				panic(err)
			}
			log.Printf(
				"accepted connection: %s -> %s",
				conn.LocalAddr(),
				conn.RemoteAddr(),
			)
			go processRequest(conn)
		}
	}()

	time.Sleep(time.Hour)
}
