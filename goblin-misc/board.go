package misc

import (
	"errors"
	"fmt"

	"github.com/nsf/termbox-go"
)

const (
	// X value
	X = 'X'
	// O value
	O = 'O'
	// EMPTY cell value
	EMPTY = ' '
)

// Cursor represents board cursor position
type Cursor struct {
	Col, Row         int
	FgColor, BgColor termbox.Attribute
}

// Cell structure defines possible cell state, it can be eiather X, O or EMPTY
type Cell rune

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
	// fill default EMPTY cells
	for i := range board.Content {
		board.Content[i] = EMPTY
	}
	return board
}

// GetWidth returns the actual width of a board
func (p *BoardDescription) GetWidth() int {
	return (p.CellsHoriz - 1) * 4
}

// GetHeight returns the actual height of a board
func (p *BoardDescription) GetHeight() int {
	return (p.CellsVert - 1) * 2
}

// GetHorizSlice returns a slice of any row of a board from start to end inclusive
func (p *BoardDescription) GetHorizSlice(row, start, end int) []Cell {
	startIdx, _ := p.ToLinear(start, row)
	endIdx, _ := p.ToLinear(end, row)
	return p.Content[startIdx : endIdx+1]
}

func (p *BoardDescription) GetVertSlice(col, start, end int) []Cell {
	var tmp = make([]Cell, end-start)

	for i := start; i <= end; i++ {
		fmt.Printf("%d ", i)
		idx, _ := p.ToLinear(col, i)
		tmp = append(tmp, p.Content[idx])
	}
	return tmp
}

// ToLinear converts col and row into linear address
func (p *BoardDescription) ToLinear(col, row int) (int, error) {
	if col >= 0 && row >= 0 && col < p.CellsHoriz && row < p.CellsVert {
		return row*p.CellsHoriz + col, nil
	}
	return -1, errors.New("Index out of bounds error")
}

// SetCell setup cell value for a give col and row
func (p *BoardDescription) SetCell(col, row int, val Cell) {
	idx, err := p.ToLinear(col, row)
	if err == nil {
		p.Content[idx] = val
	}
}

// GetCell returns cell value for a given col and row
func (p *BoardDescription) GetCell(col, row int) Cell {
	idx, err := p.ToLinear(col, row)
	if err == nil {
		return p.Content[idx]
	}
	panic(err)
}
