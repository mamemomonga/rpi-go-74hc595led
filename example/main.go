package main

import (
	leds "github.com/mamemomonga/rpi-go-74hc595led/simple"
	"log"
	"time"
)

func main() {
	leds.Start(1)
	defer leds.Finalize()
	log.Printf("Start")

	for {
		for i := uint8(0); i < 8; i++ {
			leds.Set(0, i, 1)
			time.Sleep(time.Millisecond * 500)
			leds.Set(0, i, 0)
		}
	}
}
