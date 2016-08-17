package misc

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sort"
)

// KMPPrefixTable is a helper function for KMPSearch that generates
// prefix table for Knuth-Morris-Pratt algorithm
func KMPPrefixTable(pattern []Cell) []int {

	result := make([]int, len(pattern))
	i, j := 0, 1

	for ;i < len(pattern) && j < len(pattern); {

		if pattern[i] == pattern[j] {

			result[j] = i + 1
			i++
			j++

		} else {

			if i > 0 {
				i = result[i-1]
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

	indices := []int{}
	offset := 0

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
func scanLine(line []Cell, col, row int, pattern []Cell, direction ScanDirection) IntervalList {

	result := IntervalList{}

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

	return result
}

func generateWinningPatterns(winLength int) {

	fillPattern := func(player Cell, length int) []Cell {
		tmp := make([]Cell, length)
		for j := 0; j < length; j++ {
			tmp[j] = player
		}
		return tmp
	}

	// required number in a row
	winningPatternsX.winNow = fillPattern(X, winLength)

	// all - 1 in a row + empty cells on sides always
	// leads to victory for player
	winningPatternsX.winInAMove = fillPattern(X, winLength+1)
	winningPatternsX.winInAMove[0] = E
	winningPatternsX.winInAMove[winLength] = E

	// same for O player
	winningPatternsO.winNow = fillPattern(O, winLength)

	winningPatternsO.winInAMove = fillPattern(O, winLength+1)
	winningPatternsO.winInAMove[0] = E
	winningPatternsO.winInAMove[winLength] = E

}

func getWinningPatterns(player Cell) PatternType {
	if player == X {
		return winningPatternsX
	}
	return winningPatternsO
}

// FindPattern finds vertical, horizontal or diagonal patterns generated using MakePatterns,
// returns list of pattern matched intervals or empty list if nothing was found
func FindPattern(board *BoardDescription, pattern []Cell) IntervalList {

	var (
		matchHoriz IntervalList
		matchVert  IntervalList
		matchDiag  IntervalList
	)

	// scan horizontal first

	for i := 0; i < board.CellsVert; i++ {
		tmp := scanLine(board.GetHorizSlice(i, 0, board.CellsHoriz-1), 0, i, pattern, horizontal)
		if tmp != nil {
			matchHoriz = append(matchHoriz, tmp...)
		}
	}

	// then vertical

	for i := 0; i < board.CellsHoriz; i++ {
		tmp := scanLine(board.GetVertSlice(i, 0, board.CellsVert-1), i, 0, pattern, vertical)
		if tmp != nil {
			matchVert = append(matchVert, tmp...)
		}
	}

	// and finally diagonal

	for i := 0; i < board.CellsVert; i++ {
		tmp := scanLine(board.GetRLDiagonal(0, i), i, 0, pattern, RLDiagonal)
		if tmp != nil {
			matchDiag = append(matchDiag, tmp...)
		}
	}

	for i := 1; i < board.CellsHoriz; i++ {
		tmp := scanLine(board.GetRLDiagonal(i, board.CellsVert-1), board.CellsVert-1, i,
			pattern, RLDiagonal)
		if tmp != nil {
			matchDiag = append(matchDiag, tmp...)
		}
	}

	for i := 0; i < board.CellsHoriz; i++ {
		tmp := scanLine(board.GetLRDiagonal(i, 0), i, 0, pattern, LRDiagonal)
		if tmp != nil {
			matchDiag = append(matchDiag, tmp...)
		}
	}

	for i := 1; i < board.CellsVert; i++ {
		tmp := scanLine(board.GetLRDiagonal(0, i), 0, i, pattern, LRDiagonal)
		if tmp != nil {
			matchDiag = append(matchDiag, tmp...)
		}
	}

	return append(matchHoriz, append(matchVert, matchDiag...)...)
}

// ShuffleIntSlice shuffles a slice of ints in-place
func ShuffleIntSlice(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		idx := i + 1 + rand.Intn(len(slice)-i-1)
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

// checkWin determines whether there is N in-a-row Xs or Os on a board
// which would mean that there is a winner and the game is over
func checkWin(board *BoardDescription, player Cell) (bool, IntervalList) {

	pattern := getWinningPatterns(player).winNow

	intervals := FindPattern(board, pattern)

	if len(intervals) != 0 {
		return true, intervals
	}

	return false, intervals
}

// updateScores updates scores array according to Monte-Carlo outcomes
func updateScores(board *BoardDescription, opponent, winner Cell, scores []int) {

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

func normalizeScores(scores []int) []float64 {

	sum, scoresNorm := 0, make([]float64, len(scores))

	for _, v := range scores {
		sum += v
	}

	for idx, v := range scores {
		scoresNorm[idx] = float64(v) / float64(sum)
	}

	return scoresNorm
}

// MonteCarloEval uses Monte-Carlo method to assess current position, intended to be used
// as a heuristic to reduce search space
func MonteCarloEval(board *BoardDescription, options AIOptions, maxDepth, trials int, movesFirst Cell) []float64 {

	// for implementation simplicity we only search for a full length winning sequence here
	// it allows to make such a simple method without handling special cases, more thorough
	// analysis will be performed in other methods

	opponent := switchPlayer(options.AIPlayer)

	type trialType struct {
		board    *BoardDescription
		whoMoves Cell
	}

	sem := make(chan int, runtime.GOMAXPROCS(0))
	out := make(chan trialType, trials)

	// freeIndices are always the same
	freeIndices := board.GetFreeIndices()

	for trial := 0; trial < trials; trial++ {

		numFreeCells := board.NumFreeCells()

		go func() {

			sem <- 1

			// clone existing board
			clonedBoard := CloneBoard(board)

			// shuffle free cells
			tmp := make([]int, board.NumCells())
			copy(tmp, freeIndices)

			freeShuffled := ShuffleIntSlice(tmp)

			// compute number of iterations for each trial
			iterations := minIntPair(minIntPair(numFreeCells, maxDepth), len(freeShuffled))

			whoMoves := movesFirst

			for i := 1;; i++ {

				clonedBoard.SetCellLinear(freeShuffled[i-1], whoMoves)

				// if there is a winner on current move
				winner, _ := checkWin(clonedBoard, whoMoves)

				if winner {
					out <- trialType{clonedBoard, whoMoves}
					break
				}

				if i < iterations {
					whoMoves = switchPlayer(whoMoves)
				} else {
					out <- trialType{}
					break
				}
			}

			<-sem
		}()

	}

	scores := make([]int, board.NumCells())

	for resultsReceived := 0; resultsReceived < trials; resultsReceived++ {
		data := <-out
		if data.board != nil {
			updateScores(data.board, opponent, data.whoMoves, scores)
		}
	}

	return normalizeScores(scores)
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

// ArrangeMonteCarloResults sorts result of monte carlo evaluation in descending order
func ArrangeMonteCarloResults(board *BoardDescription, options AIOptions, maxDepth, trials int, whoMoves Cell) IntFloatPairs {

	scores := MonteCarloEval(board, options, maxDepth, trials, whoMoves)
	tmp := make(IntFloatPairs, len(scores))

	freeIndices := make(Set)

	for _, v := range board.GetFreeIndices() {
		freeIndices[v] = true
	}

	for idx, v := range scores {
		// accept free cells only
		if _, found := freeIndices[idx]; found {
			tmp[idx].Fst = idx
			tmp[idx].Snd = v
		}
	}

	sort.Sort(tmp)

	return tmp
}

// filterArrangedResults leaves moves valued greater than threshold and returns
// their indices on a board
func filterArrangedResults(moves IntFloatPairs, threshold float64) []int {
	result := make([]int, 0, len(moves))
	for _, v := range moves {
		if v.Snd >= threshold {
			result = append(result, v.Fst)
		}
	}
	return result
}

// Reduce search space for minmax analysis by figuring out the most important move using
// Monte-Carlo heuristic
func ReduceSearchSpaceMonteCarlo(board *BoardDescription, options AIOptions, maxDepth int, threshold float64) []int {
	arrangedMoves := ArrangeMonteCarloResults(board, options, board.NumFreeCells(), maxDepth, options.AIPlayer)
	cellsToCheck := filterArrangedResults(arrangedMoves, threshold)
	return cellsToCheck
}

// Statically analyze board position by search some simple winning patterns
// Main principles are:
// 1. n in-a-row or n-1 in a row have the biggest grade
// 2. longer chains - bigger grade
func StaticPositionAnalyzer(board *BoardDescription, options AIOptions, whoMoves Cell) int {

	winningPatterns := getWinningPatterns(whoMoves)

	// winning/losing in a move position
	intervals := FindPattern(board, winningPatterns.winInAMove)

	if len(intervals) != 0 {
		if whoMoves == options.AIPlayer {
			return WON - 1
		} else {
			return LOST + 1
		}
	}

	// win/lose now
	intervals2 := FindPattern(board, winningPatterns.winNow)

	if len(intervals2) != 0 {
		if whoMoves == options.AIPlayer {
			return WON
		} else {
			return LOST
		}
	}

	// scan for longest chains

	return NOTHING

}

// MinMax evaluation with optional alpha-beta pruning
func MinMaxEval(board *BoardDescription, options AIOptions, cellsToCheck []int,
	lastMove LinearMove, depth int) (int, int) {

	whoMoves := lastMove.player
	whoMoved := switchPlayer(whoMoves)
	selectedMove := lastMove.position

	positionScore := StaticPositionAnalyzer(board, options, whoMoved)

	if board.NumFreeCells() != 0 {

		cellsGen := false
		if cellsToCheck == nil {
			cellsToCheck = intRange(board.NumCells())
			cellsGen = true
		}

		if depth > 0 && positionScore != WON && positionScore != LOST {

			positionScore = math.MaxInt64
			if whoMoves == options.AIPlayer {
				positionScore = -positionScore
			}

			boardCopy := CloneBoard(board)

			for idx, cellIdx := range cellsToCheck {

				if cellsGen {
					// swap index in value in case of intRange generator
					if boardCopy.GetCellLinear(idx) != E { continue }
					cellIdx = idx
				}

				board = CloneBoard(boardCopy)

				board.SetCellLinear(cellIdx, whoMoves)

				//if options.useGoRoutines && depth == options.maxDepth

				_, curVal := MinMaxEval(board, options, nil,
					LinearMove{cellIdx, whoMoved}, depth-1)

				if whoMoves == options.AIPlayer {

					// try to maximize score
					if curVal >= positionScore {
						selectedMove = cellIdx
						positionScore = curVal
					}

				} else {

					// opponent tries to minimize score
					if curVal <= positionScore {
						selectedMove = cellIdx
						positionScore = curVal
					}

				}

			}
		}
	}

	return selectedMove, positionScore

}

// Function to choose the best move from a given position
func MakeMove(board *BoardDescription, options AIOptions) (Cell, []Interval) {

	generateWinningPatterns(options.winSequenceLength)

	// Use Monte-Carlo for static evaluation

	opponent := switchPlayer(options.AIPlayer)

	playerWon, intervals := checkWin(board, opponent)

	if playerWon {
		return opponent, intervals
	}

	cellsToCheck := ReduceSearchSpaceMonteCarlo(board, options, 500, 0.1)

	if len(cellsToCheck) == 0 {
		cellsToCheck = nil
	}

	fmt.Println(cellsToCheck)

	bestLinear, bestVal := MinMaxEval(board, options, cellsToCheck,
		LinearMove{0, options.AIPlayer}, options.maxDepth)

	col, row, err := board.FromLinear(bestLinear)

	if err == nil && board.GetCell(col, row) == E {
		bestMove := CellPosition{col, row}
		fmt.Println(bestLinear, bestMove, bestVal)
		board.SetCell(bestMove.Col, bestMove.Row, options.AIPlayer)

	} else {
		// TODO: Add proper logging
		fmt.Println("ERROR")
	}

	AIWon, intervals := checkWin(board, options.AIPlayer)

	if AIWon {
		return options.AIPlayer, intervals
	}

	return E, []Interval{}
}
