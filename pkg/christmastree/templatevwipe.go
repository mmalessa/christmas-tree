package christmastree

import (
	"time"
)

/*
config:
	colors: [0xRRGGBB, 0xRRGGBB]	// wipe color
	tick: 100						// sleep time (ms)
*/
func (ch *ChristmasTree) PlayTemplateVWipe(config map[string]interface{}) error {
	colors := config["colors"].([]interface{})
	tick := config["tick"].(int)
	rows := ch.treeconfig["rows"].([]interface{})

	for _, color := range colors {
		for _, row := range rows {
			rowrange := row.(map[interface{}]interface{})
			min := rowrange["min"].(int)
			max := rowrange["max"].(int)
			for i := min; i <= max; i++ {
				ch.dev.Leds(0)[i] = uint32(color.(int))
			}
			if err := ch.dev.Render(); err != nil {
				return err
			}
			time.Sleep(time.Duration(tick) * time.Millisecond)
		}
	}
	return nil
}
