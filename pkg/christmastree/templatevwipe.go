package christmastree

import (
	"time"
)

/*
config:
	colors: [0xRRGGBB, 0xRRGGBB]	// wipe color
	tick: 100						// sleep time (ms)
	direction: 0					// 0 or 1 (1==reverse)
*/
func (ch *ChristmasTree) PlayTemplateVWipe(config map[string]interface{}) error {
	colors := config["colors"].([]interface{})
	tick := config["tick"].(int)
	direction := config["direction"].(int)
	rows := ch.treeconfig["rows"].([]interface{})
	countrows := len(rows)

	colorRow := func(color int, rowrange map[interface{}]interface{}) error {
		min := rowrange["min"].(int)
		max := rowrange["max"].(int)
		for i := min; i <= max; i++ {
			ch.dev.Leds(0)[i] = uint32(color)
		}
		if err := ch.dev.Render(); err != nil {
			return err
		}
		time.Sleep(time.Duration(tick) * time.Millisecond)
		return nil
	}

	for _, color := range colors {
		if direction == 0 {
			for i := 0; i < countrows; i++ {
				row := rows[i]
				colorRow(color.(int), row.(map[interface{}]interface{}))
			}
		} else {
			for i := countrows - 1; i >= 0; i-- {
				row := rows[i]
				colorRow(color.(int), row.(map[interface{}]interface{}))
			}
		}
	}
	return nil
}
