package christmastree

func (ch *ChristmasTree) SafeSetLed(ledid int, color uint32) error {
	if ledid >= 0 && ledid < ch.ledCount {
		ch.dev.Leds(0)[ledid] = color
	}
	return nil
}
