package port

import (
	"errors"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestScanPort(t *testing.T) {
	netScannerMock := NetScannerMock{DialTimeoutMock: func(network, address string, timeout time.Duration) (net.Conn, error) {
		time.Sleep(1 * time.Second)
		return &ConnMock{}, nil
	}}

	netScanner = &netScannerMock
	portScanner := PortScanner{}
	var wg sync.WaitGroup
	resultChannel := make(chan Result, 1)

	go portScanner.ScanPort("tcp", "stackoverflow.com", "", 80, resultChannel, &wg)

	wg.Wait()
	result := <-resultChannel

	expected := Result{Port: 80, State: true, Service: ""}
	assert.Equal(t, expected, result)
}

func TestScanPortFail(t *testing.T) {
	netScannerMock := NetScannerMock{DialTimeoutMock: func(network, address string, timeout time.Duration) (net.Conn, error) {
		time.Sleep(1 * time.Second)
		return nil, errors.New("Error")
	}}

	netScanner = &netScannerMock
	portScanner := PortScanner{}
	var wg sync.WaitGroup
	resultChannel := make(chan Result, 1)

	go portScanner.ScanPort("tcp", "stackoverflow.com", "", 80, resultChannel, &wg)

	wg.Wait()
	result := <-resultChannel

	expected := Result{Port: 80, State: false, Service: ""}
	assert.Equal(t, expected, result)
}

func TestScanPortFailTooManyConns(t *testing.T) {
	count := 0
	netScannerMock := NetScannerMock{DialTimeoutMock: func(network, address string, timeout time.Duration) (net.Conn, error) {
		time.Sleep(1 * time.Second)
		if count == 0 {
			count = count + 1
			return nil, errors.New("too many open files")
		}
		return &ConnMock{}, nil
	}}

	netScanner = &netScannerMock
	portScanner := PortScanner{}
	var wg sync.WaitGroup
	resultChannel := make(chan Result, 1)

	go portScanner.ScanPort("tcp", "stackoverflow.com", "", 80, resultChannel, &wg)

	wg.Wait()
	result := <-resultChannel

	expected := Result{Port: 80, State: true, Service: ""}
	assert.Equal(t, expected, result)
}

func TestScanPorts(t *testing.T) {
	netScannerMock := NetScannerMock{}
	netScannerMock.LookupIPMock = func(host string) ([]net.IP, error) {
		return []net.IP{}, nil
	}
	netScannerMock.DialTimeoutMock = func(network, address string, timeout time.Duration) (net.Conn, error) {
		time.Sleep(1 * time.Second)
		return &ConnMock{}, nil
	}

	netScanner = &netScannerMock
	portScanner := PortScanner{}

	results, err := portScanner.ScanPorts("stackoverflow.com", Range{Start: 0, End: 2}, 5)

	expected := ScanResult{hostname: "stackoverflow.com", ip: []net.IP{}, results: nil}

	assert.Equal(t, expected, results)
	assert.NoError(t, err)
}

// func TestScanPortsFailLookupIP(t *testing.T) {
// 	netScannerMock := NetScannerMock{}
// 	netScannerMock.LookupIPMock = func(host string) ([]net.IP, error) {
// 		return []net.IP{}, errors.New("Error")
// 	}

// 	netScanner = &netScannerMock
// 	portScanner := PortScanner{}

// 	results, _ := portScanner.ScanPorts("stackover", Range{Start: 0, End: 2}, 5)

// 	expected := ScanResult{hostname: "", ip: nil, results: nil}

// 	assert.Equal(t, expected, results)
// }
