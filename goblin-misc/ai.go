package misc

import (
	"math/rand"
	"fmt"
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

// KMPPrefixTable is a helper function for KMPSearch that generates
// prefix table for Knuth-Morris-Pratt algorithm
func KMPPrefixTable(pattern []Cell) []int {

	result := make([]int, len(pattern))
	i, j := 0, 1

	for ;i < len(pattern) && j < len(pattern); {

		if pattern[i] == pattern[j] {

			result[j] = i + 1
			i++; j++

		} else {

			if i > 0 {
				i = result[i - 1]
			}

			if pattern[i] == pattern[j] {
				result[j] = result[i] + 1
			}

			if i == 0 || (i != 0 && pattern[i] == pattern[j]) {
				j++
			}

		}
	}

	return result
}

// KMPSearch uses Knuth-Morris-Pratt algorithm to find needles in haystacks
// in O(n) instead of naive O(n^2) =)
func KMPSearch(needle, haystack []Cell) (bool, int) {
	table := KMPPrefixTable(needle)
	i, j := 0, 0

	if len(needle) == 0 {
		return true, 0
	}

	for ;j < len(haystack); {
		if needle[i] == haystack[j] {
			if i == len(needle) - 1 {
				return true, j - i
			}
			i++; j++
		} else {

			if i != 0 {
				j += i - table[i - 1] - 1
				i = table[i - 1]
			} else {
				j++
			}
		}
	}

	return false, -1

}


// findAllSubslices returns a list of indices of all position of subslice xs in
// slice ys, returns [] if xs is not in ys
func findAllSubslices(xs, ys []Cell) []int {

	var indices []int
	var offset int

	for {
		found, idx := KMPSearch(xs, ys)
		if found {
			delta := idx + len(xs) - 1
			indices = append(indices, offset+idx)
			ys = ys[delta:]
			offset += delta
		} else {
			break
		}
	}

	return indices
}

// scanLine accept a slice (horizontal, vertical or diagonal), col and row of slice star in board coordinates,
// patterns list to find and returns all intervals which match given patterns
func scanLine(line []Cell, col, row int, patterns [][]Cell, player Cell, direction ScanDirection) []Interval {

	result := []Interval{}

	for _, pattern := range patterns {

		for _, position := range findAllSubslices(pattern, line) {

			seqLen := len(pattern)

			if direction == horizontal {
				result = append(result, Interval{horizontal, position, row, position + seqLen - 1, row})

			} else if direction == vertical {
				result = append(result, Interval{vertical, col, position, col, position + seqLen - 1})

			} else if direction == LRDiagonal {
				result = append(result, Interval{LRDiagonal, col + position,
					row + position, col + position + seqLen - 1, row + position + seqLen - 1})

			} else if direction == RLDiagonal {
				result = append(result, Interval{RLDiagonal, col - position,
					row + position, col - position - seqLen + 1, row + position + seqLen - 1})
			}

		}
	}

	return result
}

// MakePatterns generates winning patterns of specified length to search on a board
func MakePatterns(targetLen int, p Cell) [][]Cell {

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

// FindPattern finds vertical, horizontal or diagonal patterns generated using MakePatterns
func FindPattern(board *BoardDescription, seqLen int, player Cell) []Interval {

	var (
		matchHoriz []Interval
		matchVert  []Interval
		matchDiag  []Interval
	)

	patterns := MakePatterns(seqLen, player)

	// scan horizontal first

	for i := 0; i < board.CellsVert; i++ {
		// get slice of each row
		row := board.GetHorizSlice(i, 0, board.CellsHoriz-1)
		// and scan for a chain
		tmp := scanLine(row, 0, i, patterns, player, horizontal)

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
		tmp := scanLine(column, i, 0, patterns, player, vertical)

		// if there was a positive result copy it to the global result
		if tmp != nil {
			matchVert = append(matchVert, tmp...)
		}
	}

	// and finally diagonal

	for i := 0; i < board.CellsVert; i++ {
		tmp := scanLine(board.GetRLDiagonal(0, i), i, 0, patterns, player, RLDiagonal)
		if tmp != nil {	matchDiag = append(matchDiag, tmp...) }
	}

	for i := 1; i < board.CellsHoriz; i++ {
		tmp := scanLine(board.GetRLDiagonal(i, board.CellsVert - 1), board.CellsVert - 1, i,
			patterns, player, RLDiagonal)
		if tmp != nil {	matchDiag = append(matchDiag, tmp...) }
	}

	for i := 0; i < board.CellsHoriz; i++ {
		tmp := scanLine(board.GetLRDiagonal(i, 0), i, 0, patterns, player, LRDiagonal)
		if tmp != nil { matchDiag = append(matchDiag, tmp...) }
	}

	for i := 1; i < board.CellsVert; i++ {
		tmp := scanLine(board.GetLRDiagonal(0, i), 0, i, patterns, player, LRDiagonal)
		if tmp != nil {	matchDiag = append(matchDiag, tmp...) }
	}

	return append(matchHoriz, append(matchVert, matchDiag...)...)
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
		ai_wins, player_wins int
	)

	for trial := 0; trial < trials; trial++ {

		clonedBoard := CloneBoard(board)
		free := ShuffleIntSlice(board.GetFreeIndices())

		maxMoves = minIntPair(minIntPair(board.NumCells(), maxMoves), len(free))

		for i := 0; i < maxMoves; i++ {

			// check whether we have guaranteed winning position
			results := FindPattern(clonedBoard, options.winSequenceLength, whoMoves)
			fmt.Println(clonedBoard)
			fmt.Println(results)

			if len(results) != 0 {
				if whoMoves == options.AIPlayer {
					// computer will win in next move(s)
					ai_wins++
					fmt.Println(i)
					break
				} else {
					// opponent will win in next move(s)
					player_wins++
					fmt.Println("Oh no!")
					break
				}
			}

			col, row, _:= clonedBoard.FromLinear(free[0])
			clonedBoard.SetCell(col, row, whoMoves)

			free = free[1:]
			whoMoves = SwitchPlayer(whoMoves)
		}

	}

	wining_prob := float64(ai_wins) / (float64(ai_wins) + float64(player_wins))

	return wining_prob
}

// Function to choose the best move from a given position
func MakeMove(board *BoardDescription, options AIOptions) {

}
