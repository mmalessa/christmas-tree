package christmastree

import (
	"time"
)

/*
config:
	colors: [0xRRGGBB, 0xRRGGBB]	// wipe color
	tick: 100						// sleep time (ms)
*/
func (ch *ChristmasTree) PlayTemplateWipe(config map[string]interface{}) error {
	colors := config["colors"].([]interface{})
	tick := config["tick"].(int)
	for _, color := range colors {
		for i := 0; i < ch.ledCount; i++ {
			ch.dev.Leds(0)[i] = uint32(color.(int))
			if err := ch.dev.Render(); err != nil {
				return err
			}
			time.Sleep(time.Duration(tick) * time.Millisecond)
		}
	}
	return nil
}
