package Server

import (
	"errors"
	"net"
)

type server struct {
	Address  *net.TCPAddr
	Protocol string
	Listener *net.TCPListener
}

func (s server) GetListner() *net.TCPListener {
	return s.Listener
}

func (s *server) Start() error {
	if s.Listener != nil {
		return errors.New("Server already started.")
	}
	listener, err := net.ListenTCP(s.Protocol, s.Address)
	s.Listener = listener
	return err
}

func NewServer(address *net.TCPAddr, protocol string) (Server, error) {
	newServer := new(server)
	newServer.Address = address
	newServer.Protocol = protocol
	err := newServer.Start()
	return newServer, err
}

type Server interface {
	Start() error
	GetListner() *net.TCPListener
}
