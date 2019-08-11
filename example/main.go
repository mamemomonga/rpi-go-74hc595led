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
		for i:=0; i<2; i++ {
			leds.AllHigh()
			time.Sleep(time.Millisecond * 100)
			leds.AllLow()
			time.Sleep(time.Millisecond * 100)
		}
		for i := uint8(0); i < 8; i++ {
			leds.Set(0, i, 1)
			time.Sleep(time.Millisecond * 100)
		}
		for i := uint8(0); i < 8; i++ {
			leds.Set(0, 8-i, 0)
			time.Sleep(time.Millisecond * 100)
		}
		for i := uint8(0); i < 8; i++ {
			leds.Set(0, i, 1)
			time.Sleep(time.Millisecond * 50)
			leds.Set(0, i, 0)
		}
	}
}
