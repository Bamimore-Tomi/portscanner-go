package main

import (
	"fmt"

	"github.com/Bamimore-Tomi/portscanner-go/port"
)

func main() {
	fmt.Println("Port Scanning")
	results := port.InitialScan("localhost")
	fmt.Println(results)
}
