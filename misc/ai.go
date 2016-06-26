package misc

import (
	"math/rand"
	"math"
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

// scanLine accept a slice (horizontal, vertical or diagonal), col and row of slice start in board coordinates,
// patterns list to find and returns all intervals which match given patterns
func scanLine(line []Cell, col, row int, patterns [][]Cell, player Cell, direction ScanDirection) []Interval {

	result := []Interval{}

	for _, pattern := range patterns {

		for _, position := range findAllSubslices(pattern, line) {

			seqLen := len(pattern)

			if direction == horizontal {
				result = append(result, Interval{horizontal, CellPosition{position, row},
					CellPosition{position + seqLen - 1, row}})

			} else if direction == vertical {
				result = append(result, Interval{vertical, CellPosition{col, position},
					CellPosition{col, position + seqLen - 1}})

			} else if direction == LRDiagonal {
				result = append(result, Interval{LRDiagonal, CellPosition{col + position,
					row + position}, CellPosition{col + position + seqLen - 1, row + position + seqLen - 1}})

			} else if direction == RLDiagonal {
				result = append(result, Interval{RLDiagonal, CellPosition{col - position,
					row + position}, CellPosition{col - position - seqLen + 1, row + position + seqLen - 1}})
			}

		}
	}

	return result
}

// patternBuilder is a helper function which returns another
// function to effectively generate sequences to scan on a board with caching
func patternBuilder() func(int, Cell) [][]Cell {

	winningSequences := make(map[struct{ int; Cell}][][]Cell)

	return func (targetLen int, p Cell) [][]Cell {

		key := struct{int; Cell}{targetLen, p}

		sequences, ok := winningSequences[key]

		if !ok {

			winningSequences[key] = [][]Cell{}

			// test all in a row (for instance: X, X, X, X, X is a winner)
			winningSequences[key] = append(winningSequences[key], make([]Cell, targetLen))

			// test all minus 1 in a row
			winningSequences[key] = append(winningSequences[key], make([]Cell, targetLen + 1))

			// fill all patterns patterns with player cells
			for i := 0; i < 2; i++ {
				for j := 0; j < targetLen + 2; j++ {
					if len(winningSequences[key][i]) > j {
						winningSequences[key][i][j] = p
					}
				}
			}

			// add empty cells
			winningSequences[key][1][0] = E
			winningSequences[key][1][len(winningSequences[key][1]) - 1] = E

			sequences = winningSequences[key]

		}

		return sequences

	}

}

// MakePatterns generates winning patterns of specified length to search on a board
var MakePatterns = patternBuilder()


// FindPattern finds vertical, horizontal or diagonal patterns generated using MakePatterns
func FindPattern(board *BoardDescription, player Cell, patterns [][]Cell) []Interval {

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

func switchPlayer(player Cell) Cell {
	if player == X {
		return O
	}
	return X
}

func checkWin(board *BoardDescription, opt AIOptions) Cell {

	// check whether we have guaranteed winning position

	opponent := switchPlayer(opt.AIPlayer)

	patterns := MakePatterns(opt.winSequenceLength, opt.AIPlayer)
	if len(FindPattern(board, opt.AIPlayer, patterns)) != 0 {
		return opt.AIPlayer
	}

	patterns = MakePatterns(opt.winSequenceLength, opponent)
	if len(FindPattern(board, opponent, patterns)) != 0 {
		return opponent
	}

	return E
}

func isGameOver(board *BoardDescription, opt AIOptions, player Cell) (bool, Interval) {

	pattern := [][]Cell{MakePatterns(opt.winSequenceLength, player)[0]}

	intervals := FindPattern(board, player, pattern)

	if len(intervals) != 0 {
		return true, intervals[0]
	}

	return false, Interval{}
}

// updateScores updates scores array according to Monte-Carlo outcomes
func updateScores(board *BoardDescription, opponent, winner Cell, scores []float64) {

	for idx := 0; idx < board.NumCells(); idx++ {

		if board.Content[idx] == E {
			continue
		}

		if board.Content[idx] == opponent {
			if opponent == winner {
				scores[idx]++
			} else {
				scores[idx]--
			}
		} else {
			if opponent == winner {
				scores[idx]--
			} else {
				scores[idx]++
			}
		}
	}
}

// MonteCarloEval uses Monte-Carlo method to assess current position, intended to be used
// as a static evaluator for leaf nodes
func MonteCarloEval(board *BoardDescription, options AIOptions, maxDepth, trials int, movesFirst Cell) []float64 {

	opponent := switchPlayer(options.AIPlayer)
	scores := make([]float64, board.NumCells())

	for trial := 0; trial < trials; trial++ {

		// clone existing board
		clonedBoard := CloneBoard(board)

		// shuffle free cells
		free := ShuffleIntSlice(clonedBoard.GetFreeIndices())

		// compute number of iterations for each trial
		iterations := minIntPair(minIntPair(board.NumCells(), maxDepth), len(free))

		whoMoves := movesFirst

		for i := 0;; i++ {

			// if there is a 100% winner on this trial
			winner := checkWin(clonedBoard, options)

			if winner != E {

				updateScores(clonedBoard, opponent, winner, scores)
				break

			}

			if i < iterations {

				clonedBoard.SetCellLinear(free[0], whoMoves)

				whoMoves = switchPlayer(whoMoves)
				free = free[1:]

			} else if i >= iterations { break }
		}

	}

	return scores
}

// MonteCarloBestMove search for a best move by using Monte-Carlo evaluation, accepts board description, ai options
// struct, maxDepth - is a maximum depth to scan, trials - number of trial and whoMoves which states who is going
// to make next move
func MonteCarloBestMove(board *BoardDescription, options AIOptions, maxDepth, trials int, whoMoves Cell) (CellPosition, float64) {

	scores := MonteCarloEval(board, options, maxDepth, trials, whoMoves)

	bestValue, bestMove := -math.MaxFloat64, 0

	for _, idx := range board.GetFreeIndices() {
		if bestValue < scores[idx] {
			bestValue = scores[idx]
			bestMove = idx
		}
	}

	col, row, _ := board.FromLinear(bestMove)

	return CellPosition{col, row}, bestValue
}


// Function to choose the best move from a given position
func MakeMove(board *BoardDescription, options AIOptions) (Cell, Interval) {
	// Use Monte-Carlo for static evaluation

	opponent := switchPlayer(options.AIPlayer)

	playerWon, interval := isGameOver(board, options, opponent)

	if playerWon {
		return opponent, interval
	}

	bestMove, _ := MonteCarloBestMove(board, options, board.NumFreeCells(),
		300, options.AIPlayer)

	board.SetCell(bestMove.col, bestMove.row, options.AIPlayer)

	AIWon, interval := isGameOver(board, options, options.AIPlayer)

	if AIWon {
		return options.AIPlayer, interval
	}

	return E, Interval{}
}
