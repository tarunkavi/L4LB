package main

import (
	"fmt"
	"net"
)

type client struct {
	conn net.Conn
	addr net.Addr
}

func newClient(conn net.Conn) *client {
	return &client{
		conn: conn,
		addr: conn.RemoteAddr(),
	}
}

func (c *client) readLoop() {
	buf := make([]byte, 1024)
	for {
		n, err := c.conn.Read(buf)
		if err != nil {
			return
		}
		data := buf[:n]
		fmt.Println(string(data))
	}
}
