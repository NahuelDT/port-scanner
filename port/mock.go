package port

import (
	"net"
	"time"
)

type NetScannerMock struct {
	DialTimeoutMock func(network, address string, timeout time.Duration) (net.Conn, error)
	LookupIPMock    func(host string) ([]net.IP, error)
}

func (p *NetScannerMock) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return p.DialTimeoutMock(network, address, timeout)
}

//LookupIP -
func (p *NetScannerMock) LookupIP(host string) ([]net.IP, error) {
	return p.LookupIPMock(host)
}

type ConnMock struct {
	net.Conn
}

func (c *ConnMock) Close() error {
	return nil
}
