package misc

import (
	"errors"
	"math"
	"math/rand"
)

const (
	// X value
	X = 'X'
	// O value
	O = 'O'
	// EMPTY cell value
	E = ' '
)


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

	// number of horizontal and vertical cells
	CellsHoriz int
	CellsVert  int

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

// ReverseCellsSlice reverses a slice of board cells in-place
func ReverseCellsSlice(slice []Cell) []Cell {
	for i, j := 0, len(slice) - 1; i < j; i, j = i + 1, j - 1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// NewBoard creates a new struct of type BoardDescription with allocated
// slice for a board contents
func NewBoard(cellsHoriz, cellsVert int) *BoardDescription {

	board := &BoardDescription{cellsHoriz, cellsVert, make([]Cell, cellsHoriz*cellsVert)}

	// fill default EMPTY cells
	for i := range board.Content {
		board.Content[i] = E
	}

	return board
}

// CloneBoard clones an existing board
func CloneBoard(p *BoardDescription) *BoardDescription {
	newBoard := NewBoard(p.CellsHoriz, p.CellsVert)
	for i, v := range(p.Content) {
		newBoard.Content[i] = v
	}
	return newBoard
}

// GetRandomizedBoard returns a board randomly filled with Xs and Os
// and with the given percent of empty cells
func GetRandomizedBoard(cellsHoriz, cellsVert int, emptyPercent float64) *BoardDescription {

	board := NewBoard(cellsHoriz, cellsVert)

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
		dd = diagonalDistance(startCol, startRow, endCol, endRow)
		idx, tmp = 0, make([]Cell, dd)
	)

	if dd == 1 {
		return []Cell{p.GetCell(startCol, startRow)}
	}

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
	} else if startCol > endCol && startRow > endRow {
		for idx < dd {
			tmp[idx] = p.GetCell(startCol, startRow)
			idx++; startCol--; startRow--
		}
	} else if startCol < endCol && startRow > endRow {
		for idx < dd {
			tmp[idx] = p.GetCell(startCol, startRow)
			idx++; startCol++; startRow--
		}
	}

	return tmp
}

// GetBounds returns start and end coordinates of a diagonal specified by one of its cells
// and direction
func (p *BoardDescription) GetBounds(col, row int, direction Direction) (int, int, int, int) {

	if direction == RightToLeft {
		maxDeltaUp := minIntPair(p.CellsHoriz-col-1, row)
		maxDeltaDown := minIntPair(col, p.CellsVert-row-1)

		return col + maxDeltaUp, row - maxDeltaUp,
			col - maxDeltaDown, row + maxDeltaDown
	}

	maxDeltaUp := minIntPair(col, row)
	maxDeltaDown := minIntPair(p.CellsHoriz-col-1, p.CellsVert-row-1)

	return col - maxDeltaUp, row - maxDeltaUp,
		col + maxDeltaDown, row + maxDeltaDown

}

// GetRightDiagonal returns diagonal starting at col, row till
// the end of the board (from Left to Right)
func (p *BoardDescription) GetLRDiagonal(col, row int) []Cell {
	startCol, startRow, endCol, endRow := p.GetBounds(col, row, LeftToRight)
	return p.GetDiagonalSliceXY(startCol, startRow, endCol, endRow)
}

// GetLeftDiagonal returns diagonal starting at col, row till
// the end of the board (from Right to Left)
func (p *BoardDescription) GetRLDiagonal(col, row int) []Cell {
	startCol, startRow, endCol, endRow := p.GetBounds(col, row, RightToLeft)
	return p.GetDiagonalSliceXY(startCol, startRow, endCol, endRow)
}

// ToLinear converts col and row into linear address
func (p *BoardDescription) ToLinear(col, row int) (int, error) {
	if col >= 0 && row >= 0 && col < p.CellsHoriz && row < p.CellsVert {
		return row*p.CellsHoriz + col, nil
	}
	return -1, errors.New("Index out of bounds error")
}

// FromLinear converts linear index into pair (col, row)
func (p *BoardDescription) FromLinear(idx int) (int, int, error) {
	if idx >= p.CellsVert * p.CellsHoriz || idx < 0 {
		return 0, 0, errors.New("Index out of bounds error")
	}
	row := idx / p.CellsHoriz
	col := idx % p.CellsVert
	return col, row, nil
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

// GetFreeIndices returns indices of free board cells
func (p *BoardDescription) GetFreeIndices() []int {

	var result []int

	for i, v := range p.Content {
		if v == E {
			result = append(result, i)
		}
	}
	return result
}
