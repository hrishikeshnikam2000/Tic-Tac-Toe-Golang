package components

import (
	"errors"
)

const (
	Xmark  = "X"
	Omark  = "O"
	Nomark = "-"
)

type Cell struct {
	Mark string
}

func NewCell() *Cell {
	newCell := &Cell{Mark: Nomark}
	return newCell
}

func (c *Cell) SetMark(mark string) error {
	if c.Mark == Nomark {
		c.Mark = mark
		return nil
	}
	return errors.New("The cell is already marked !")
}

func (c *Cell) GetMark() string {
	return c.Mark
}
