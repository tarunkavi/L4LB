package main

import (
	"math/rand"
	"net"
)

type backend struct {
	address     string
	healthy     bool
	activeConns int
	conn        net.Conn
}

type backendManager struct {
	backends []*backend
	length   uint32
}

func (m *backendManager) register(address string) {
	if _, ok := m.get(address); ok {
		return
	}
	m.backends = append(m.backends, newBackend(address))
	m.length++
}

func (m *backendManager) get(address string) (*backend, bool) {
	for _, b := range m.backends {
		if b.address == address {
			return b, true
		}
	}
	return nil, false
}
func (m *backendManager) distribute() *backend {
	n := 0
	for _, b := range m.backends {
		if b.healthy {
			n++
		}
	}
	if n == 0 {
		return nil
	}
	i := rand.Int() % n
	for _, b := range m.backends {
		if !b.healthy {
			continue
		}
		if i == 0 {
			return b
		}
		i--
	}
	return nil
}

func newBackend(address string) *backend {
	return &backend{
		address: address,
		healthy: true,
	}
}

func (b *backend) dial() (net.Conn, error) {
	return net.Dial("tcp", b.address)
}

func send(conn net.Conn, data []byte) error {
	_, err := conn.Write(data)
	return err
}
