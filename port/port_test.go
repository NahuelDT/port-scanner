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
	portScanner := PortScanner{}
	netScannerMock := NetScannerMock{DialTimeoutMock: func(network, address string, timeout time.Duration) (net.Conn, error) {
		return &ConnMock{}, nil
	}}

	netScanner = &netScannerMock
	var wg sync.WaitGroup
	hostname := "stackoverflow.com"
	resultChannel := make(chan Result, 1)
	port := 80
	go portScanner.ScanPort("tcp", hostname, "", port, resultChannel, &wg)
	wg.Wait()

	result := <-resultChannel

	expected := Result{Port: 80, State: true, Service: ""}

	assert.Equal(t, expected, result)
}

func TestScanPortFail(t *testing.T) {
	portScanner := PortScanner{}
	netScannerMock := NetScannerMock{DialTimeoutMock: func(network, address string, timeout time.Duration) (net.Conn, error) {
		return nil, errors.New("Error")
	}}

	netScanner = &netScannerMock
	var wg sync.WaitGroup
	hostname := "stackoverflow.com"
	resultChannel := make(chan Result, 1)
	port := 80
	go portScanner.ScanPort("tcp", hostname, "", port, resultChannel, &wg)
	wg.Wait()

	result := <-resultChannel

	expected := Result{Port: 80, State: false, Service: ""}

	assert.Equal(t, expected, result)
}

func TestScanPortFailTooManyConns(t *testing.T) {
	count := 0
	portScanner := PortScanner{}
	netScannerMock := NetScannerMock{DialTimeoutMock: func(network, address string, timeout time.Duration) (net.Conn, error) {
		if count == 0 {
			count = count + 1
			return nil, errors.New("too many open files")
		}
		return &ConnMock{}, nil
	}}

	netScanner = &netScannerMock
	var wg sync.WaitGroup
	hostname := "stackoverflow.com"
	resultChannel := make(chan Result, 1)
	port := 80
	go portScanner.ScanPort("tcp", hostname, "", port, resultChannel, &wg)
	wg.Wait()

	result := <-resultChannel

	expected := Result{Port: 80, State: true, Service: ""}

	assert.Equal(t, expected, result)
}
