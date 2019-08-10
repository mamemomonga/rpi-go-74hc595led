package simple

import (
	"github.com/janne/bcm2835"
	pi595led "github.com/mamemomonga/rpi-go-74hc595led"
)

var (
	chip_num uint8
	leds []uint8
)

func Start(chip_n uint8) {
	pi595led.Init()
	chip_num = chip_n
	leds = make([]uint8,chip_num)
}

func Finalize() {
	bcm2835.SpiEnd()
	bcm2835.Close()
}

func Set(chip uint8, pin uint8, state uint8) {
	switch(state) {
	case 0:
		leds[chip] &=^ ( 1 << pin )
	case 1:
		leds[chip] |=  ( 1 << pin )
	}
	bcm2835.SpiTransfern(leds)
}

