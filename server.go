package main

import "net"

type server struct {
	config   config
	listener net.Listener
	backends *backendManager
	clients  *clientManager
}

func newServer(cfg config) *server {
	m := &backendManager{}
	for _, bc := range cfg.Backends {
		m.register(bc.Address)
	}
	return &server{
		config:   cfg,
		backends: m,
		clients:  &clientManager{},
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
		b := s.backends.distribute()
		if b == nil {
			conn.Close()
			continue
		}
		c := s.clients.register(conn)
		go c.forward(b)
	}
}
