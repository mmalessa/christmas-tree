package christmastree

import (
	"time"
)

/*
config:
	time: 100		// sleep time (ms)
*/
func (ch *ChristmasTree) PlayTemplateSleep(config map[string]interface{}) error {
	timems := config["time"].(int)
	time.Sleep(time.Duration(timems) * time.Millisecond)
	return nil
}
