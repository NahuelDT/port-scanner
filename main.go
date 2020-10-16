package main

import (
	"fmt"

	"github.com/NahuelDT/port-scanner.git/port"
)

func main() {
	fmt.Println("Hello WildLife!")

	result := port.ScanPort("stackoverflow.com", 30)
	fmt.Println(result)

}
