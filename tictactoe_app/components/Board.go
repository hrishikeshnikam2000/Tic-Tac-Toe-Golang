package components

type Board struct {
	Cells []*Cell
	Size  uint8
}

func NewBoard(size uint8) *Board {
	BoardSize := size * size
	var cellArray = make([]*Cell, BoardSize)
	newBoard := &Board{
		Cells: cellArray,
		Size:  size,
	}

	for i := range cellArray {
		cellArray[i] = NewCell()
	}
	return newBoard
}
