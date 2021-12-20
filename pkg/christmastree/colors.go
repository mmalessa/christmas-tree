package christmastree

import (
	gocolor "github.com/gerow/go-color"
)

func (ch *ChristmasTree) GetRainbowColor(step int, steps int) uint32 {
	// 240deg 0.666666 (blue)
	// 300deg 0.833333 (violet)
	limit := 0.666666

	angle := (float64(step) / float64(steps)) * limit
	return ch.GetColorFromHSL(angle, 1, 0.5)
}

func (ch *ChristmasTree) GetColorFromHSL(h float64, s float64, l float64) uint32 {
	rgb := gocolor.HSL{h, s, l}.ToRGB()
	r := int32(rgb.R * 0xFF)
	g := int32(rgb.G * 0xFF)
	b := int32(rgb.B * 0xFF)
	color := r<<16 + g<<8 + b
	return uint32(color)
}
