package christmastree

/*
config:
	background: 0xRRGGBB	// background color
	foreground: 0xRRGGBB	// foreground color
	direction: 0			// 0 or 1
	tick: 100				// sleep time (ms)
*/
func (ch *ChristmasTree) PlayTemplateVLine(config map[string]interface{}) error {

	// background := config["background"].(int)
	// foreground := config["foreground"].(int)
	// matrixh := len(ch.matrix)
	// matrixw := len(ch.matrix[0])
	// direction := config["direction"].(int)
	// width := 1
	// tick := config["tick"].(int)

	// if direction == 0 {
	// 	for w := 0; w < (matrixw + width); w++ {
	// 		for h := 0; h < matrixh; h++ {
	// 			if w < matrixw {
	// 				ch.dev.Leds(0)[ch.matrix[h][w]] = uint32(foreground)
	// 			}
	// 			if w > 0 {
	// 				ch.dev.Leds(0)[ch.matrix[h][w-1]] = uint32(background)
	// 			}
	// 		}
	// 		if err := ch.dev.Render(); err != nil {
	// 			return err
	// 		}
	// 		time.Sleep(time.Duration(tick) * time.Millisecond)
	// 	}
	// } else {
	// 	for w := matrixw - 1; w >= -width; w-- {
	// 		for h := 0; h < matrixh; h++ {
	// 			if w >= 0 {
	// 				ch.dev.Leds(0)[ch.matrix[h][w]] = uint32(foreground)
	// 			}
	// 			if w < matrixw {
	// 				ch.dev.Leds(0)[ch.matrix[h][w+1]] = uint32(background)
	// 			}
	// 		}
	// 		if err := ch.dev.Render(); err != nil {
	// 			return err
	// 		}
	// 		time.Sleep(time.Duration(tick) * time.Millisecond)
	// 	}
	// }

	return nil
}
