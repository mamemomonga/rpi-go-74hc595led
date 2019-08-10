package pi595led

import(
	"github.com/janne/bcm2835"
)

func Init() {
	bcm2835.Init()
	bcm2835.SpiBegin() // SPI0
	bcm2835.SpiSetBitOrder(BCM2835_SPI_BIT_ORDER_MSBFIRST)
	bcm2835.SpiSetDataMode(BCM2835_SPI_MODE2)
	bcm2835.SpiSetClockDivider(BCM2835_SPI_CLOCK_DIVIDER_128)
	bcm2835.SpiChipSelect(BCM2835_SPI_CS0)
	bcm2835.SpiSetChipSelectPolarity(BCM2835_SPI_CS0, LOW)
}

