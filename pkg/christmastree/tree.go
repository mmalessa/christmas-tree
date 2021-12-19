package christmastree

import (
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

type ChristmasTree struct {
	gpioPin    int
	ledCount   int
	brightness int
	dev        *ws2811.WS2811
	patterns   map[string]ChristmasTreePattern
	matrix     map[int]map[int]int
}

func NewChristmasTree(gpioPin int, ledCount int, brightness int) *ChristmasTree {
	ch := &ChristmasTree{
		gpioPin:    gpioPin,
		ledCount:   ledCount,
		brightness: brightness,
		patterns:   make(map[string]ChristmasTreePattern),
		matrix:     make(map[int]map[int]int),
	}
	opt := ws2811.DefaultOptions
	opt.Channels[0].GpioPin = ch.gpioPin
	opt.Channels[0].Brightness = ch.brightness
	opt.Channels[0].LedCount = ch.ledCount

	var err error
	ch.dev, err = ws2811.MakeWS2811(&opt)
	if err != nil {
		panic(err)
	}
	err = ch.dev.Init()
	if err != nil {
		panic(err)
	}
	return ch
}

func (ch *ChristmasTree) Defer() {
	ch.TurnOff()
	ch.dev.Fini()
}

func (ch *ChristmasTree) TurnOff() {
	var err error
	for x := 0; x < ch.ledCount; x++ {
		color := uint32(0x000000)
		ch.dev.Leds(0)[x] = color
	}
	err = ch.dev.Render()
	if err != nil {
		panic(err)
	}
}
