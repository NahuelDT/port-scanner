package port

import (
	"fmt"
	"net"
	"strconv"
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

//Common ports in range 1 to 1024
var common = map[int]string{
	7:   "echo",
	20:  "ftp",
	21:  "ftp",
	22:  "ssh",
	23:  "telnet",
	25:  "smtp",
	43:  "whois",
	53:  "dns",
	67:  "dhcp",
	68:  "dhcp",
	80:  "http",
	110: "pop3",
	123: "ntp",
	137: "netbios",
	138: "netbios",
	139: "netbios",
	143: "imap4",
	443: "https",
	513: "rlogin",
	540: "uucp",
	554: "rtsp",
	587: "smtp",
	873: "rsync",
	902: "vmware",
	989: "ftps",
	990: "ftps",
}

//ScanPort Scans single port, returs a Result
func ScanPort(protocol, hostname string, port int) Result {
	result := Result{Port: port}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 1*time.Second)
	if err != nil {
		result.State = false
		return result
	}
	defer conn.Close()

	result.State = true
	return result
}

//ScanPorts Scans all ports of hostname in range the range given, returns a ScanResult
func ScanPorts(hostname string, ports Range) (ScanResult, bool) {
	var results []Result
	var scanned ScanResult
	addr, err := net.LookupIP(hostname)
	if err != nil {
		return scanned, false
	}
	for i := ports.Start; i <= ports.End; i++ {
		result := ScanPort("tcp", hostname, i)
		if v, ok := common[i]; ok {
			result.Service = v
		}
		results = append(results, result)
	}
	scanned = ScanResult{
		hostname: hostname,
		ip:       addr,
		results:  results,
	}
	return scanned, true
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
func GetOpenPorts(hostname string, ports Range) {
	scanned, ok := ScanPorts(hostname, ports)
	if ok {
		DisplayScanResult(scanned)
	} else {
		fmt.Println("Error: Invalid IP address")
	}
}
