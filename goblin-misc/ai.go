package misc

import (
	"math/rand"
)

// Interval represents indexes of start and end of an n-length chain`
type Interval struct {

	direction int

	col1, row1 int
	col2, row2 int
}

type ScanDirection uint8

type AIOptions struct {

	AIPlayer Cell

	winSequenceLength int

	maxDepth int

}

const (
	horizontal = iota
	vertical

	// LR means that we are scanning from the left side of the board
	// towards its right side
	LRDiagonal

	// RL means that we are scanning from the right side of the board
	// towards its left side
	RLDiagonal
)

const (
	AI_WINS = 10
	AI_LOSES = -10
	DRAW = 0
)


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
					result = append(result, Interval{horizontal, idx - chainLen + 1, row, idx, row})

				} else if direction == vertical {
					result = append(result, Interval{vertical, col, idx - chainLen + 1, col, idx})

				} else if direction == LRDiagonal {
					result = append(result, Interval{LRDiagonal, col + idx - chainLen + 1,
						row + idx - chainLen + 1, col + idx, row + idx})

				} else if direction == RLDiagonal {
					result = append(result, Interval{RLDiagonal, col + idx - chainLen + 1,
						row - idx + chainLen - 1, col + idx, row - idx})
				}
			}
			counter = 0
		}
	}

	return result
}

// MakeSearchPatterns generates winning patterns of
// specified length to search on a board
func MkaeWinningPatterns(targetLen int, p Cell) [][]Cell {

	winningPatterns := [][]Cell{}

	// test all in a row (for instance: X, X, X, X, X is a winner)
	winningPatterns = append(winningPatterns, make([]Cell, targetLen))

	// test all minus 1 in a row
	winningPatterns = append(winningPatterns, make([]Cell, targetLen - 1 + 1))
	winningPatterns = append(winningPatterns, make([]Cell, targetLen - 1 + 1))

	// test all minus 2 in a row
	winningPatterns = append(winningPatterns, make([]Cell, targetLen - 2 + 4))

	// fill all patterns patterns with player cells
	for i := 0; i < 4; i++ {
		for j := 0; j < targetLen + 2; j++ {
			if len(winningPatterns[i]) > j {
				winningPatterns[i][j] = p
			}
		}
	}

	// add empty cells

	winningPatterns[1][0] = E
	winningPatterns[2][targetLen - 1] = E

	l := len(winningPatterns[3])
	winningPatterns[3][0] = E
	winningPatterns[3][1] = E
	winningPatterns[3][l - 1] = E
	winningPatterns[3][l - 2] = E

	return winningPatterns
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
		tmp := scanLine(ReverseCellsSlice(board.GetRLDiagonal(0, i)), 0, i, chainLen, player, RLDiagonal)
		if tmp != nil {	matchDiag = append(matchDiag, tmp...) }
	}

	for i := 1; i < board.CellsHoriz; i++ {
		tmp := scanLine(ReverseCellsSlice(board.GetRLDiagonal(i, board.CellsVert - 1)), i,
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
		tmp := FindChain(board, chainLen, player)
		if len(tmp) != 0 {
			result[chainLen] = tmp
		}
	}
	return result
}

// ShuffleIntSlice shuffles a slice of ints in-place
func ShuffleIntSlice(slice []int) []int {
	for i := 0; i < len(slice) - 1; i++ {
		idx := i + 1 + rand.Intn(len(slice) - i - 1)
		slice[i], slice[idx] = slice[idx], slice[i]
	}
	return slice
}

func SwitchPlayer(player Cell) Cell {
	if player == X {
		return O
	}
	return X
}

// TestWin determines whether requested player guaranteed to win the game from
// the current board position (very useful for quick position evaluation
// without performing thorough analysis)
//func TestWin(board *BoardDescription, options AIOptions, player Cell) (bool, Interval) {
//
//	MAYBE_WIN := options.winSequenceLength - 2
//	WILL_WIN := options.winSequenceLength - 1
//
//	for k, v := range FindAllChains(board, MAYBE_WIN, WILL_WIN, player) {
//
//		if k == MAYBE_WIN {
//			// we have to ensure that the given position will lead to win 100% of
//			// all the cases
//			for _, interval := range v {
//				stCol, stRow := interval.col1, interval.row1
//				endCol, endRow := interval.col2, interval.row2
//			}
//
//		} else if k == WILL_WIN {
//			// given player will 100% win next move from this position
//			return true, v
//		}
//	}
//
//	return false, Interval{}
//
//}

// MonteCarloEval uses Monte-Carlo method to assess current position, intended
// to be used as static evaluator for leaf nodes
func MonteCarloEval(board *BoardDescription, options AIOptions, trials, maxMoves int, whoMoves Cell) float64 {

	var (
		cellIdx int
		ai_wins, player_wins int
	)

	maxMoves = maxIntPair(board.NumCells(), maxMoves)

	for trial := 0; trial < trials; trial++ {

		clonedBoard := CloneBoard(board)
		free := ShuffleIntSlice(board.GetFreeIndices())

		for i := 0; i < maxMoves; i++ {

			cellIdx = free[0]
			col, row, err := clonedBoard.FromLinear(cellIdx)

			if err == nil {

				clonedBoard.SetCell(col, row, whoMoves)

				// check whether we have guaranteed winning position
				results := FindAllChains(clonedBoard, options.winSequenceLength - 1,
					options.winSequenceLength, whoMoves)

				if len(results) != 0 {
					if whoMoves == options.AIPlayer {
						// computer will win in next move(s)
						ai_wins++
						break
					} else {
						// opponent will win in next move(s)
						player_wins++
						break
					}
				}

				free = free[1:]
				whoMoves = SwitchPlayer(whoMoves)

			} else {
				panic("Index out of bounds in method MonteCarlEval")
			}
		}

	}

	wining_prob := float64(ai_wins) / (float64(ai_wins) + float64(player_wins))

	return wining_prob
}

// Function to choose the best move from a given position
func MakeMove(board *BoardDescription, options AIOptions) {

}