package christmastree

import (
	"time"
)

/*
config:
	tick: 100						// sleep time (ms)
*/
func (ch *ChristmasTree) PlayTemplateRainbow(config map[string]interface{}) error {
	tick := config["tick"].(int)
	for i := 0; i < ch.ledCount; i++ {
		color := ch.GetRainbowColor(i, ch.ledCount-1, 0.666666)
		ch.dev.Leds(0)[i] = color
		if err := ch.dev.Render(); err != nil {
			return err
		}
		time.Sleep(time.Duration(tick) * time.Millisecond)
	}
	return nil
}
