package christmastree

import (
	"time"
)

func (ch *ChristmasTree) PlayTemplateVRainbow(config map[string]interface{}) error {

	matrixh := len(ch.matrix)
	matrixw := len(ch.matrix[0])
	direction := config["direction"].(int)
	tick := config["tick"].(int)

	if direction == 0 {
		for w := 0; w < matrixw; w++ {
			color := ch.GetRainbowColor(w, matrixw-1)
			for h := 0; h < matrixh; h++ {
				if w < matrixw {
					ch.dev.Leds(0)[ch.matrix[h][w]] = color
				}
			}
			if err := ch.dev.Render(); err != nil {
				return err
			}
			time.Sleep(time.Duration(tick) * time.Millisecond)
		}
	} else {
		for w := matrixw - 1; w >= 0; w-- {
			color := ch.GetRainbowColor(w, matrixw-1)
			for h := 0; h < matrixh; h++ {
				if w >= 0 {
					ch.dev.Leds(0)[ch.matrix[h][w]] = color
				}
			}
			if err := ch.dev.Render(); err != nil {
				return err
			}
			time.Sleep(time.Duration(tick) * time.Millisecond)
		}
	}

	return nil
}
