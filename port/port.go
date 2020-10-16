package port

import (
	"fmt"
	"net"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

//Result port scan results
type Result struct {
	Port    int
	State   bool
	Service string
}

//Range of ports for a Scan
type Range struct {
	Start, End int
}

//ScanResult Results from all ports of hostname
type ScanResult struct {
	hostname string
	ip       []net.IP
	results  []Result
}

var portScanner Scanner = &PortScanner{}

//Common ports in range 1 to 1024
var common = map[int]string{
	7:    "echo",
	20:   "ftp",
	21:   "ftp",
	22:   "ssh",
	23:   "telnet",
	25:   "smtp",
	43:   "whois",
	53:   "dns",
	67:   "dhcp",
	68:   "dhcp",
	80:   "http",
	110:  "pop3",
	123:  "ntp",
	137:  "netbios",
	138:  "netbios",
	139:  "netbios",
	143:  "imap4",
	443:  "https",
	513:  "rlogin",
	540:  "uucp",
	554:  "rtsp",
	587:  "smtp",
	873:  "rsync",
	902:  "vmware",
	989:  "ftps",
	990:  "ftps",
	1194: "openvpn",
	3306: "mysql",
	5000: "unpn",
	8080: "https-proxy",
	8443: "https-alt",
}

//ScanPort Scans single port, returs a Result
func ScanPort(protocol, hostname, service string, port int, resultChannel chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	result := Result{Port: port, Service: service}
	address := hostname + ":" + strconv.Itoa(port)

	conn, err := portScanner.DialTimeout(protocol, address, 2*time.Second)
	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(1 * time.Second)
			ScanPort("tcp", hostname, service, port, resultChannel, wg)
		} else {
			//fmt.Println(port, "closed") //INDICATE CLOSED PORTS
			// fmt.Println("ERR", err)
			result.State = false
			resultChannel <- result
		}
		return
	}

	//fmt.Println(port, "open") //INDICATE OPEN PORTS

	defer conn.Close()
	result.State = true
	resultChannel <- result
	return
}

//ScanPorts Scans all ports of hostname in range the range given, returns a ScanResult
func ScanPorts(hostname string, ports Range, threads int) (ScanResult, error) {
	var results []Result
	var scanned ScanResult
	var wg sync.WaitGroup

	runtime.GOMAXPROCS(threads)

	resultChannel := make(chan Result, ports.End-ports.Start+1)

	addr, err := portScanner.LookupIP(hostname)
	if err != nil {
		return scanned, err
	}

	for i := ports.Start; i <= ports.End; i++ {
		service, _ := common[i]
		go ScanPort("tcp", hostname, service, i, resultChannel, &wg)
	}
	wg.Wait()

	close(resultChannel)

	for result := range resultChannel {

		results = append(results, result)
	}

	scanned = ScanResult{
		hostname: hostname,
		ip:       addr,
		results:  results,
	}
	return scanned, nil
}

//DisplayScanResult Displays the scan result
func DisplayScanResult(result ScanResult) {
	ip := result.ip[len(result.ip)-1]
	fmt.Printf("Open ports for %s (%s)\n", result.hostname, ip.String())
	for _, v := range result.results {
		if v.State {
			fmt.Printf("%d	%s\n", v.Port, v.Service)
		}
	}
}

//GetOpenPorts Calls ScanPorts and Displays the Results
func GetOpenPorts(hostname string, ports Range, threads int) {
	scanned, err := ScanPorts(hostname, ports, threads)
	if err != nil {
		fmt.Println(err)
	} else {
		DisplayScanResult(scanned)
	}
}
