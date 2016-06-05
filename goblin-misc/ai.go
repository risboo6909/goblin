package misc

import (
	"math/rand"
)


// Interval represents indexes of start and end of an n-length chain`
type Interval struct {
	col1, row1 int
	col2, row2 int
}

type ScanDirection uint8

const (
	horizontal = iota
	vertical
	LRDiagonal
	RLDiagonal
)

func reverseSlice(slice []Cell) []Cell {
	for i, j := 0, len(slice) - 1; i < j; i, j = i + 1, j - 1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}


// ShuffleSlice shuffles a slice of cells and shuffles it in-place
func ShuffleSlice(slice []Cell)  {
	for i := 0; i < len(slice) - 1; i++ {
		idx := i + 1 + rand.Intn(len(slice) - i - 1)
		slice[i], slice[idx] = slice[idx], slice[i]
	}
}

// scanLine accept a slice (horizontal, vertical or diagonal), col and row are coordinates of a sequence start
// length of desired sequence, type of cell (X, O or E) and a scan direction
func scanLine(line []Cell, col, row, chainLen int, player Cell, direction ScanDirection) []Interval {


	var (
		result  []Interval
		counter = 0
	)

	// convenience functions
	pred_match := func (idx int, cell, player Cell) bool {
		return player == cell && idx < len(line) - 1
	}

	pred_nomatch := func (idx int, cell, player Cell) bool {
		return player == cell && idx == len(line) - 1
	}

	for idx, cell := range line {

		if pred_match(idx, cell, player) {

			counter++

		} else {

			if pred_nomatch(idx, cell, player) {
				counter++
			} else {
				idx--
			}

			if counter == chainLen {
				if direction == horizontal {
					result = append(result, Interval{idx - chainLen + 1, row, idx, row})

				} else if direction == vertical {
					result = append(result, Interval{col, idx - chainLen + 1, col, idx})

				} else if direction == LRDiagonal {
					result = append(result, Interval{col + idx - chainLen + 1, row + idx - chainLen + 1,
						col + idx, row + idx})

				} else if direction == RLDiagonal {
					result = append(result, Interval{col + idx - chainLen + 1, row - idx + chainLen - 1,
						col + idx, row - idx})
				}
			}
			counter = 0
		}
	}

	return result
}

// FindChain finds vertical, horizontal or diagonal chains of
// successive cells with the same content as a slice of Intervals
func FindChain(board *BoardDescription, chainLen int, player Cell) []Interval {

	var (
		matchHoriz []Interval
		matchVert  []Interval
		matchDiag  []Interval
	)

	// scan horizontal first

	for i := 0; i < board.CellsVert; i++ {
		// get slice of each row
		row := board.GetHorizSlice(i, 0, board.CellsHoriz-1)
		// and scan for a chain
		tmp := scanLine(row, 0, i, chainLen, player, horizontal)

		// if there was a positive result copy it to the global result
		if tmp != nil {
			matchHoriz = append(matchHoriz, tmp...)
		}
	}

	// then vertical

	for i := 0; i < board.CellsHoriz; i++ {
		// get slice of each column
		column := board.GetVertSlice(i, 0, board.CellsVert-1)
		// and scan for a chain
		tmp := scanLine(column, i, 0, chainLen, player, vertical)

		// if there was a positive result copy it to the global result
		if tmp != nil {
			matchVert = append(matchVert, tmp...)
		}
	}

	// and finally diagonal

	for i := 0; i < board.CellsVert; i++ {
		tmp := scanLine(reverseSlice(board.GetRLDiagonal(0, i)), 0, i, chainLen, player, RLDiagonal)
		if tmp != nil {	matchDiag = append(matchDiag, tmp...) }
	}

	for i := 1; i < board.CellsHoriz; i++ {
		tmp := scanLine(reverseSlice(board.GetRLDiagonal(i, board.CellsVert - 1)), i,
			board.CellsVert - 1, chainLen, player, RLDiagonal)
		if tmp != nil {	matchDiag = append(matchDiag, tmp...) }
	}

	for i := 0; i < board.CellsHoriz; i++ {
		tmp := scanLine(board.GetLRDiagonal(i, 0), i, 0, chainLen, player, LRDiagonal)
		if tmp != nil { matchDiag = append(matchDiag, tmp...) }
	}

	for i := 1; i < board.CellsVert; i++ {
		tmp := scanLine(board.GetLRDiagonal(0, i), 0, i, chainLen, player, LRDiagonal)
		if tmp != nil {	matchDiag = append(matchDiag, tmp...) }
	}

	return append(matchHoriz, append(matchVert, matchDiag...)...)
}

// FindAllChains finds all the chains of length in interval [chainLenMin..chainLenMax]
func FindAllChains(board *BoardDescription, chainLenMin, chainLenMax int, player Cell) map[int][]Interval {
	result := make(map[int][]Interval)
	for chainLen := chainLenMin; chainLen <= chainLenMax; chainLen++ {
		result[chainLen] = FindChain(board, chainLen, player)
	}
	return result
}


// MonteCarloEval uses Monte-Carlo method to asess current position, intended
// to be used as static evaluator for leaf nodes
func MonteCarloEval(board *BoardDescription, maxMoves int) {

	for i := 0; i < maxMoves; i++ {

	}
}