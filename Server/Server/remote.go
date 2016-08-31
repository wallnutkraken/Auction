package Server

import (
	"errors"
	"net"
)

type client struct {
	connection *net.TCPConn
	echo       bool
}

func (c *client) GetConnection() *net.TCPConn {
	return c.connection
}

func (c *client) InvertEcho() bool {
	c.echo = !c.echo
	return c.echo
}

func (c *client) GetEcho() bool {
	return c.echo
}

func (c *client) Send(reply Reply) error {
	if bytes, err := reply.JSON(); err == nil {
		if _, err := c.connection.Write(bytes); err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}

func NewClient(connection *net.TCPConn) (Client, error) {
	if connection == nil {
		return nil, errors.New("Connection cannot be null")
	}
	cl := new(client)
	cl.connection = connection
	return cl, nil
}

type Client interface {
	GetConnection() *net.TCPConn
	InvertEcho() bool
	GetEcho() bool
	Send(Reply) error
}
