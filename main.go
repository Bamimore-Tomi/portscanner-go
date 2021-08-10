package main

import (
	"fmt"

	"github.com/Bamimore-Tomi/portscanner-go/port"
)

func main() {
	fmt.Println("Port Scanning")
	results := port.InitialScan("localhost")
	fmt.Println(results)

	wisescanresult := port.WideScan("localhost")
	fmt.Println(wisescanresult)

}
