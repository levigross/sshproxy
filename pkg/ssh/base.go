package ssh

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/pkg/errors"
)

// Config is your ssh config
type Config struct {
	HostPort    int
	DstPort     int
	DstHostname string

	listen net.Listener

	fatalError chan error
}

func (c *Config) validate() (err error) {
	c.fatalError = make(chan error, 1)

	c.listen, err = net.Listen("tcp", fmt.Sprintf(":%d", c.HostPort))
	if err != nil {
		return errors.Wrap(err, "unable to create listner")
	}

	if c.DstHostname == "" {
		return errors.New("the Dest Hostname cannot be blank")
	}

	return nil
}

// Run ...
func (c *Config) Run() error {
	if err := c.validate(); err != nil {
		return errors.Wrap(err, "failed validation")
	}
	return c.run()
}

func (c *Config) run() error {
	go c.listener()
	go c.sender()
	return <-c.fatalError
}
func (c *Config) listener() {
	for {
		conn, err := c.listen.Accept()
		if err != nil {
			c.fatalError <- errors.Wrap(err, "unable to accept connection")
		}
		log.Println("Got connection", conn.LocalAddr())
		go c.handleConn(conn)
	}

}
func (c *Config) sender() {

}
func (c *Config) handleConn(conn net.Conn) {
	d, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.DstHostname, c.DstPort))
	if err != nil {
		c.fatalError <- errors.Wrap(err, "unable to dial out connection")
	}
	log.Println("Dialed", d.RemoteAddr())
	go c.copyTraffic(d, conn)
	go c.copyTraffic(conn, d)
}
func (c *Config) copyTraffic(dst net.Conn, src net.Conn) {
	for {
		if _, err := io.Copy(dst, src); err != nil {
			c.fatalError <- errors.Wrap(err, "unable to copy bytes ")
			return
		}
	}
}
