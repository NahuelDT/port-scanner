package main

import (
	"fmt"

	"github.com/NahuelDT/port-scanner.git/port"
)

func main() {
	fmt.Println("Hello WildLife!")
	port.GetOpenPorts("scanme.nmap.org", port.Range{Start: 20, End: 100})

}
