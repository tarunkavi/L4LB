package main

import (
	"fmt"
	"io"
	"net"
)

type client struct {
	conn net.Conn
	addr net.Addr
}

type clientManager struct {
	clients map[string]*client
}

func (m *clientManager) register(conn net.Conn) *client {
	if m.clients == nil {
		m.clients = make(map[string]*client)
	}
	c := newClient(conn)
	m.clients[c.addr.String()] = c
	return c
}

func (m *clientManager) get(address string) (*client, bool) {
	c, ok := m.clients[address]
	return c, ok
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

func (c *client) forward(b *backend) error {
	backendConn, err := b.dial()
	if err != nil {
		return err
	}
	_, err = io.Copy(backendConn, c.conn)
	return err
}
