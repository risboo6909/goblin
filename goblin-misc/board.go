package misc

import (
	"errors"
	"math"
	"math/rand"
	"github.com/nsf/termbox-go"
)

const (
	// X value
	X = 'X'
	// O value
	O = 'O'
	// EMPTY cell value
	E = ' '
)

// Cursor represents board cursor position
type Cursor struct {
	Col, Row         int
	FgColor, BgColor termbox.Attribute
}

// Mimic python set
type Set map[interface{}]bool

// Cell structure defines possible cell state, it can be eiather X, O or EMPTY
type Cell rune

func randomCell(cellValues ...Cell) Cell {
	idx := rand.Intn(minIntPair(len(cellValues), 3))
	return cellValues[idx]
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

type Direction uint8

const (
	RightToLeft = iota
	LeftToRight
)

func diagonalDistance(startCol, startRow, endCol, endRow int) int {
	diagonalDistance := int(((math.Abs(float64(endCol-startCol)) +
		math.Abs(float64(endRow-startRow)) + 2) / 2))
	return diagonalDistance
}

func minIntPair(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

// NewBoard creates a new struct of type BoardDescription with allocated
// slice for a board contents
func NewBoard(cellsHoriz, cellsVert, x, y int, boardColor, boardBg, labelsColor,
	labelsBg termbox.Attribute) *BoardDescription {

	board := &BoardDescription{cellsHoriz, cellsVert, x, y, boardColor, boardBg,
		labelsColor, labelsBg, make([]Cell, cellsHoriz*cellsVert)}

	// fill default EMPTY cells
	for i := range board.Content {
		board.Content[i] = E
	}

	return board
}

// GetRandomizedBoard returns a board randomly filled with Xs and Os
// and with the given percent of empty cells
func GetRandomizedBoard(cellsHoriz, cellsVert int, emptyPercent float64) *BoardDescription {

	board := NewBoard(cellsHoriz, cellsVert, 0, 0, termbox.ColorBlack, termbox.ColorBlue,
		termbox.ColorRed, termbox.ColorBlack)

	// reserve empty cells
	emptyCount := int(emptyPercent * float64(board.NumCells()) / 100)
	emptyCells := make(Set)

	for {
		if len(emptyCells) > emptyCount {break}

		randIdx := rand.Intn(board.NumCells())
		if _, found := emptyCells[randIdx]; !found {
			emptyCells[randIdx] = true
		}
	}

	for idx := 0; idx < board.NumCells(); idx++ {
		if _, found := emptyCells[idx]; !found {
			board.Content[idx] = randomCell(X, O)
		}
	}

	return board
}

// NumCells returns total number of cells
func (p *BoardDescription) NumCells() int {
	return p.CellsHoriz * p.CellsVert
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

	var (
		dd = diagonalDistance(startCol, startRow, endCol-1, endRow-1)
		tmp = make([]Cell, dd)
		idx = 0
	)

	if startCol < endCol && startRow < endRow {
		for idx < dd {
			tmp[idx] = p.GetCell(startCol, startRow)
			idx++; startCol++; startRow++
		}

	} else if startCol > endCol && startRow < endRow {
		for idx < dd {
			tmp[idx] = p.GetCell(startCol, startRow)
			idx++; startCol--; startRow++
		}
	}

	return tmp
}

// getBounds returns start and end coordinates of a diagonal specified by one of its cells
// and direction
func (p *BoardDescription) getBounds(col, row int, direction Direction) (int, int, int, int) {

	if direction == RightToLeft {
		maxDeltaUp := minIntPair(p.CellsHoriz-col, row) - 1
		maxDeltaDown := minIntPair(col, p.CellsVert-row) - 1

		return col + maxDeltaUp, row - maxDeltaUp,
			col - maxDeltaDown, row + maxDeltaDown
	}

	maxDeltaUp := minIntPair(col, row)
	maxDeltaDown := minIntPair(p.CellsHoriz-col, p.CellsVert-row)

	return col - maxDeltaUp, row - maxDeltaUp,
		col + maxDeltaDown, row + maxDeltaDown

}

// GetRightDiagonal returns diagonal starting at col, row till
// the end of the board (from Left to Right)
func (p *BoardDescription) GetLRDiagonal(col, row int) []Cell {
	startCol, startRow, endCol, endRow := p.getBounds(col, row, LeftToRight)
	return p.GetDiagonalSliceXY(startCol, startRow, endCol, endRow)
}

// GetLeftDiagonal returns diagonal starting at col, row till
// the end of the board (from Right to Left)
func (p *BoardDescription) GetRLDiagonal(col, row int) []Cell {
	startCol, startRow, endCol, endRow := p.getBounds(col, row, RightToLeft)
	return p.GetDiagonalSliceXY(startCol, startRow, endCol, endRow)
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
