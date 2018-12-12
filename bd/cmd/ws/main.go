package main

import (
	"flag"
	"log"
	"net"
	"sync"

	"github.com/gobwas/ws"
)

var (
	addr = flag.String("addr", "127.0.0.1:8080", "addr to bind to")
)

type Chat struct {
	mu    sync.Mutex
	users map[string]*User
}

func NewChat() *Chat {
	return &Chat{
		users: make(map[string]*User),
	}
}

func (c *Chat) Register(u *User) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.users[u.Name()] = u
}

func (c *Chat) Remove(u *User) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.users, u.Name())
}

func (c *Chat) Broadcast(msg []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, user := range c.users {
		user.Send(msg)
	}
}

type User struct {
	conn  net.Conn
	sendq chan []byte
}

func NewUser(qsize int, conn net.Conn) *User {
	return &User{
		conn:  conn,
		sendq: make(chan []byte, qsize),
	}
}

func (u *User) Name() string {
	return u.conn.RemoteAddr().String()
}

func (u *User) Send(msg []byte) error {
	// FIXME: handle timeouts.
	u.sendq <- msg
	return nil
}

func (u *User) Recv() ([]byte, error) {
	f, err := ws.ReadFrame(u.conn)
	if err != nil {
		return nil, err
	}
	return f.Payload, nil
}

func (u *User) drainSendQ() {
	for msg := range u.sendq {
		// FIXME: handle errors.
		ws.WriteFrame(u.conn, ws.NewTextFrame(msg))
	}
}

func (u *User) Close() {
	u.conn.Close()
}

func main() {
	flag.Parse()

	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening at %s", ln.Addr())

	c := NewChat()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		u := NewUser(100, conn)
		c.Register(u)

		go func() {
			for {
				msg, err := u.Recv()
				if err != nil {
					c.Remove(u)
					u.Close()
					return
				}
				c.Broadcast(msg)
			}
		}()
		go u.drainSendQ()
	}
}
