package misc

import (
	"errors"
	"math"

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

// GetVertSlice returns a slice of any column of a board from start to end inclusive
func (p *BoardDescription) GetVertSlice(col, start, end int) []Cell {
	var tmp = make([]Cell, end-start+1)
	for i := start; i <= end; i++ {
		tmp[i-start] = p.GetCell(col, i)
	}
	return tmp
}

// GetDiagonalSliceXY returns a slice of a diagonal starts at startCol, startRow to
// the endCol, endRow inclusive
func (p *BoardDescription) GetDiagonalSliceXY(startCol, startRow, endCol, endRow int) []Cell {
	idx := 0
	diagonalDistance := int(((math.Abs(float64(endCol-startCol)) +
		math.Abs(float64(endRow-startRow)) + 2) / 2))

	var tmp = make([]Cell, diagonalDistance)

	for startCol < endCol && startRow < endRow {
		tmp[idx] = p.GetCell(startCol, startRow)
		idx++
		// We could handle this with just one variable increment
		// but I leave both of them for clarity
		startCol++
		startRow++
	}

	return tmp
}

type Direction uint8

const (
	Left  = iota // Right to Left diagonal
	Right        // Left to Right diagonal
)

func getBounds(col, row int, direction Direction) (int, int) {
	if direction == Left {

	} else if direction == Right {

	}
}

// GetRightDiagonal returns diagonal starting at col, row till
// the end of the board (from Left to Right)
func (p *BoardDescription) GetRightDiagonal(col, row int) []Cell {

	return p.GetDiagonalSliceXY(col, row, endCol, endRow)

}

// GetLeftDiagonal returns diagonal starting at col, row till
// the end of the board (from Right to Left)
func (p *BoardDescription) GetLeftDiagonal(col, row int) []Cell {

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
