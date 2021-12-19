package main

import (
	"fmt"
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

const (
	brightness = 128
	width      = 50
	height     = 0
	ledCounts  = width * height
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Elo, elo")
	opt := ws2811.DefaultOptions
	opt.Channels[0].Brightness = brightness
	opt.Channels[0].LedCount = ledCounts

	dev, err := ws2811.MakeWS2811(&opt)
	checkError(err)

	checkError(dev.Init())
	defer dev.Fini()

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			color := uint32(0xff0000)
			if x > 2 && x < 5 && y > 0 && y < 7 {
				color = 0xffffff
			}
			if x > 0 && x < 7 && y > 2 && y < 5 {
				color = 0xffffff
			}
			dev.Leds(0)[x*height+y] = color
		}
	}
	checkError(dev.Render())
}