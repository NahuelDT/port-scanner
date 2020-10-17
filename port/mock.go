package port

import (
	"net"
	"sync"
	"time"
)

type PortScannerMock struct {
	ScanPortMock          func(protocol, hostname, service string, port int, resultChannel chan Result, wg *sync.WaitGroup, fireWallDetectionOff bool)
	ScanPortsMock         func(hostname string, ports Range, threads int, fireWallDetectionOff bool) (ScanResult, error)
	DisplayScanResultMock func(result ScanResult)
	GetOpenPortsMock      func(hostname string, ports Range, threads int)
}

func (p *PortScannerMock) ScanPort(protocol, hostname, service string, port int, resultChannel chan Result, wg *sync.WaitGroup, fireWallDetectionOff bool) {
	p.ScanPortMock(protocol, hostname, service, port, resultChannel, wg, fireWallDetectionOff)
}

func (p *PortScannerMock) ScanPorts(hostname string, ports Range, threads int, fireWallDetectionOff bool) (ScanResult, error) {
	return p.ScanPortsMock(hostname, ports, threads, fireWallDetectionOff)
}

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
