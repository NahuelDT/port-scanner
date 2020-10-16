package port

import (
	"net"
	"strconv"
	"time"
)

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

func ScanPort(hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
