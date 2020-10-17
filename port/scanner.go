package port

import (
	"net"
	"sync"
	"time"
)

//Scanner -
type Scanner interface {
	DialTimeout(network, address string, timeout time.Duration) (net.Conn, error)
	LookupIP(host string) ([]net.IP, error)
}

type HostScanner interface {
	ScanPort(protocol, hostname, service string, port int, resultChannel chan Result, wg *sync.WaitGroup, fireWallDetectionOff bool)
	ScanPorts(hostname string, ports Range, threads int, fireWallDetectionOff bool) (ScanResult, error)
	DisplayScanResult(result ScanResult)
	GetOpenPorts(hostname string, ports Range, threads int)
}

type NetScanner struct {
}

func (p *NetScanner) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return net.DialTimeout(network, address, timeout)
}

func (p *NetScanner) LookupIP(host string) ([]net.IP, error) {
	return net.LookupIP(host)
}
