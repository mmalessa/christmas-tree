package christmastree

import (
	"time"
)

/*
config:
	direction: 0		// 0 or 1 (1==reverse)
	tick: 100			// sleep time (ms)
	repeat: 10			// repeat n times
*/
func (ch *ChristmasTree) PlayTemplateRainbowUnicorn(config map[string]interface{}) error {

	tick := config["tick"].(int)
	repeat := config["repeat"].(int)
	direction := config["direction"].(int)
	rows := ch.treeconfig["rows"].([]interface{})
	rowcount := len(rows)

	colorTree := func(offset int) error {
		for step, row := range rows {
			rowrange := row.(map[interface{}]interface{})
			min := rowrange["min"].(int)
			max := rowrange["max"].(int)
			color := ch.GetRainbowColor((step+offset)%rowcount, rowcount, 1)
			for i := min; i <= max; i++ {
				ch.dev.Leds(0)[i] = color
			}
		}
		if err := ch.dev.Render(); err != nil {
			return err
		}
		time.Sleep(time.Duration(tick) * time.Millisecond)
		return nil
	}

	for r := 0; r < repeat; r++ {
		if direction == 1 {
			for offset := 0; offset < rowcount; offset++ {
				if err := colorTree(offset); err != nil {
					return err
				}
			}
		} else {
			for offset := rowcount - 1; offset >= 0; offset-- {
				if err := colorTree(offset); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
