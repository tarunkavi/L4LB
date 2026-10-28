package main

import "net"

type server struct {
	config   config
	listener net.Listener
}

func newServer(address string) *server {
	return &server{
		config: config{Address: address},
	}
}

func (s *server) listen() error {
	l, err := net.Listen("tcp", s.config.Address)
	if err != nil {
		return err
	}
	s.listener = l
	return s.acceptLoop()
}

func (s *server) acceptLoop() error {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}
		c := newClient(conn)
		c.readLoop()
	}
}
