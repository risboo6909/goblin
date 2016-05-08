package misc

import (
	"errors"

	"github.com/nsf/termbox-go"
)

const (
	X     = 'X'
	O     = 'O'
	EMPTY = ' '
)

// Cursor represents board cursor position
type Cursor struct {
	Col, Row         int
	FgColor, BgColor termbox.Attribute
}

// Cell structure defines possible cell state, it can be eiather X, O or EMPTY
type Cell struct {
	Val rune
}

// BoardDescription defines main board properties
type BoardDescription struct {
	CellsHoriz int
	CellsVert  int

	// upper-left corner position
	X, Y int

	BoardColor, BoardBg   termbox.Attribute
	LabelsColor, LabelsBg termbox.Attribute

	// board state
	Content []Cell
}

// NewBoard creates a new struct of type BoardDescription with allocated
// slice for a board contents
func NewBoard(cellsHoriz, cellsVert, x, y int, boardColor, boardBg, labelsColor,
	labelsBg termbox.Attribute) *BoardDescription {
	board := &BoardDescription{cellsHoriz, cellsVert, x, y, boardColor, boardBg,
		labelsColor, labelsBg, make([]Cell, cellsHoriz*cellsVert)}
	for i := range board.Content {
		board.Content[i].Val = EMPTY
	}
	return board
}

func (p *BoardDescription) GetWidth() int {
	return (p.CellsHoriz - 1) * 4
}

func (p *BoardDescription) GetHeight() int {
	return (p.CellsVert - 1) * 2
}

func (p *BoardDescription) ToLinear(col, row int) (int, error) {
	if col >= 0 && row >= 0 && col < p.CellsHoriz && row < p.CellsVert {
		return col + row*p.CellsHoriz, errors.New("Index out of bounds error")
	}

	return -1, nil
}

func (p *BoardDescription) SetCell(col, row int, val rune) {
	idx, err := p.ToLinear(col, row)
	if err != nil {
		p.Content[idx] = Cell{val}
	}
}

func (p *BoardDescription) GetCell(col, row int) rune {
	idx, err := p.ToLinear(col, row)
	if err != nil {
		return p.Content[idx].Val
	}
	panic(err)
}