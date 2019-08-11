package simple

import (
	"github.com/janne/bcm2835"
	pi595led "github.com/mamemomonga/rpi-go-74hc595led"
	"sync"
)

var (
	chip_num uint8
	leds []uint8
	mleds sync.Mutex
)

func Start(chip_n uint8) {
	pi595led.Init()
	chip_num = chip_n
	leds = make([]uint8,chip_num)
	AllLow()
}

func Finalize() {
	bcm2835.SpiEnd()
	bcm2835.Close()
}

func AllHigh() {
	mleds.Lock()
	defer mleds.Unlock()
	for i:=uint8(0); i<chip_num; i++ {
		leds[i]=255
	}
	bcm2835.SpiTransfern(append([]uint8{},leds...))
}

func AllLow() {
	mleds.Lock()
	defer mleds.Unlock()
	for i:=uint8(0); i<chip_num; i++ {
		leds[i]=0
	}
	bcm2835.SpiTransfern(append([]uint8{},leds...))
}

func Set(chip uint8, pin uint8, state uint8) {
	mleds.Lock()
	defer mleds.Unlock()
	switch(state) {
	case 0:
		leds[chip] &=^ ( 1 << pin )
	case 1:
		leds[chip] |=  ( 1 << pin )
	}
	bcm2835.SpiTransfern(append([]uint8{},leds...))
}

func reverse8Bit(b byte) (r byte) {
	b = ((b & 0x55) << 1) | ((b & 0xAA) >> 1)
	b = ((b & 0x33) << 2) | ((b & 0xCC) >> 2)
	return (b << 4) | (b >> 4)
}
