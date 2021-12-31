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
func (ch *ChristmasTree) PlayTemplateVRainbow(config map[string]interface{}) error {

	tick := config["tick"].(int)
	repeat := config["repeat"].(int)
	direction := config["direction"].(int)
	rows := ch.treeconfig["rows"].([]interface{})
	rowcount := len(rows)

	coloRow := func(rownumber int) error {
		rowrange := rows[rownumber].(map[interface{}]interface{})
		min := rowrange["min"].(int)
		max := rowrange["max"].(int)
		color := ch.GetRainbowColor(rownumber, rowcount, 1)
		for i := min; i <= max; i++ {
			ch.dev.Leds(0)[i] = color
		}
		if err := ch.dev.Render(); err != nil {
			return err
		}
		time.Sleep(time.Duration(tick) * time.Millisecond)
		return nil
	}

	for r := 0; r < repeat; r++ {
		if direction == 0 {
			for rownumber := 0; rownumber < rowcount; rownumber++ {
				if err := coloRow(rownumber); err != nil {
					return err
				}
			}
		} else {
			for rownumber := rowcount - 1; rownumber >= 0; rownumber-- {
				if err := coloRow(rownumber); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
