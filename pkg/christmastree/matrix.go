package christmastree

func (ch *ChristmasTree) SetMatrix(matrix map[int]map[int]int) error {
	ch.matrix = matrix
	return nil
}
