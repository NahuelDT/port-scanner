package port

import (
	"net"
	"time"
)

type Scanner interface {
	DialTimeout(network, address string, timeout time.Duration) (net.Conn, error)
	LookupIP(host string) ([]net.IP, error)
}

type PortScanner struct {
}

func (p *PortScanner) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return net.DialTimeout(network, address, timeout)
}

func (p *PortScanner) LookupIP(host string) ([]net.IP, error) {
	return net.LookupIP(host)
}
