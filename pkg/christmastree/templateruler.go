package christmastree

import (
	"time"
)

/*
config:
	normal: 0xRRGGBB	// color of other led
	fifth: 0xRRGGBB		// color of every fifth led
	tenth: 0xRRGGBB		// color of every tenth led
*/
func (ch *ChristmasTree) PlayTemplateRuler(config map[string]interface{}) error {

	normal := config["normal"].(int)
	tenth := config["tenth"].(int)
	fifth := config["fifth"].(int)

	for i := 0; i < ch.ledCount; i++ {
		var color int
		if i%10 == 0 {
			color = tenth
		} else if i%5 == 0 {
			color = fifth
		} else {
			color = normal
		}
		ch.dev.Leds(0)[i] = uint32(color)
	}
	if err := ch.dev.Render(); err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	return nil
}
