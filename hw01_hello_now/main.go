package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	localTime, err := ntp.Time("ntp1.stratum2.ru")
	if err != nil {
		log.Fatal(err)
	}
	currentTime := time.Now()
	fmt.Printf("current time: %s\n", currentTime)
	fmt.Printf("exact time: %s\n", localTime)
}
