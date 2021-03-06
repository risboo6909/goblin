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
	E = 0
)


// Cell structure defines possible cell states,
// it can be either X, O or EMPTY
type Cell byte

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
		math.Abs(float64(endRow-startRow)) + 2) * 0.5))
	return diagonalDistance
}

// NewBoard creates a new struct of type BoardDescription with allocated
// slice for a board contents
func NewBoard(cellsHoriz, cellsVert int) *BoardDescription {
	board := &BoardDescription{cellsHoriz, cellsVert, make([]Cell, cellsHoriz * cellsVert)}
	return board
}

// CloneBoard clones an existing board
func CloneBoard(p *BoardDescription) *BoardDescription {
	newBoard := NewBoard(p.CellsHoriz, p.CellsVert)
	copy(newBoard.Content, p.Content)
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

// FillBoardLinear fills up the board from a list of moves
func (p *BoardDescription) FillBoardLinear(moves []LinearMove) *BoardDescription {
	for _, move := range moves {
		p.SetCellLinear(move.position, move.player)
	}
	return p
}

// NumCells returns total number of cells
func (p *BoardDescription) NumCells() int {
	return p.CellsHoriz * p.CellsVert
}

// NumFreeCells returns number of free cells on a board
func (p *BoardDescription) NumFreeCells() int {
	freeCnt := 0
	for _, v := range p.Content {
		if v == E {
			freeCnt++
		}
	}
	return freeCnt
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

	var dd = diagonalDistance(startCol, startRow, endCol, endRow)

	if dd == 1 {
		return []Cell{p.GetCell(startCol, startRow)}
	}

	var idx, tmp = 0, make([]Cell, dd)

	if startCol > endCol {

		if startRow < endRow {
			for idx < dd {
				tmp[idx] = p.GetCell(startCol, startRow)
				idx++; startCol--; startRow++
			}

		} else if startRow > endRow {
			for idx < dd {
				tmp[idx] = p.GetCell(startCol, startRow)
				idx++; startCol--; startRow--
			}
		}

	} else if startCol < endCol {

		if startRow < endRow {
			for idx < dd {
				tmp[idx] = p.GetCell(startCol, startRow)
				idx++; startCol++; startRow++
			}

		} else if startRow > endRow {
			for idx < dd {
				tmp[idx] = p.GetCell(startCol, startRow)
				idx++; startCol++; startRow--
			}
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
		return row * p.CellsHoriz + col, nil
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
	idx, _ := p.ToLinear(col, row)
	p.SetCellLinear(idx, val)
}

// SetCellLinear setup cell value based on linear coord
func (p *BoardDescription) SetCellLinear(linearIdx int, val Cell) {
	if linearIdx < p.NumCells() {
		p.Content[linearIdx] = val
	} else {
		panic(errors.New("Index out of range"))
	}
}

// GetCell returns cell value for a given col and row
func (p *BoardDescription) GetCell(col, row int) Cell {
	idx, _ := p.ToLinear(col, row)
	return p.GetCellLinear(idx)
}

// GetCellLinear returns for a given linear index
func (p *BoardDescription) GetCellLinear(linearIdx int) Cell {
	if linearIdx < p.NumCells() {
		return p.Content[linearIdx]
	}
	panic(errors.New("Index out of range"))
}

// GetIndicesOfAKind returns all the cells which contain any item from a given
// kind list
func (p *BoardDescription) getIndicesOfAKind(kind... Cell) []int {

	var result []int = make([]int, 0, p.NumCells())

	for i, v := range p.Content {
		for _, b := range kind {
			if b == v {
				result = append(result, i)
				break
			}
		}
	}
	return result

}

// GetFreeIndices returns indices of free board cells
func (p *BoardDescription) GetFreeIndices() []int {
	return p.getIndicesOfAKind(E)
}

// GetOccupiedIndices returns indices of a board that occupied
// by X and O
func (p *BoardDescription) GetOccupiedIndices() []int {
	return p.getIndicesOfAKind(O, X)
}

// Represent board in human-readable format
func (p *BoardDescription) String() string {
	repr := "Board\n"
	for row := 0; row < p.CellsVert; row++ {
		for col := 0; col < p.CellsHoriz; col++ {
			if p.GetCell(col, row) == E {
				repr += " ."
			} else { repr += " " + string(p.GetCell(col, row)) }
		}
		repr += "\n"
	}
	return repr
}
