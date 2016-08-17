package misc

import (
	"testing"

	"reflect"
	"sort"
)

func TestFindChainDiagonal(t *testing.T) {

	var board = NewBoard(19, 19)

	// Left-to-right diagonal sequences

	board.SetCell(2, 2, X)
	board.SetCell(3, 3, X)
	board.SetCell(4, 4, X)
	board.SetCell(5, 5, X)

	board.SetCell(7, 3, X)
	board.SetCell(8, 4, X)
	board.SetCell(9, 5, X)
	board.SetCell(10, 6, X)

	board.SetCell(0, 15, X)
	board.SetCell(1, 16, X)
	board.SetCell(2, 17, X)
	board.SetCell(3, 18, X)

	generateWinningPatterns(4)
	result := FindPattern(board, getWinningPatterns(X).winNow)

	sort.Sort(result)

	assertEqual(t, result[0], Interval{LRDiagonal, CellPosition{2, 2}, CellPosition{5, 5}})
	assertEqual(t, result[1], Interval{LRDiagonal, CellPosition{7, 3}, CellPosition{10, 6}})
	assertEqual(t, result[2], Interval{LRDiagonal, CellPosition{0, 15}, CellPosition{3, 18}})

	// Right-to-left diagonal sequences

	board = NewBoard(19, 19)

	board.SetCell(13, 18, O)
	board.SetCell(14, 17, O)
	board.SetCell(15, 16, O)
	board.SetCell(16, 15, O)
	board.SetCell(17, 14, O)
	board.SetCell(18, 13, O)

	board.SetCell(0, 18, O)
	board.SetCell(1, 17, O)
	board.SetCell(2, 16, O)
	board.SetCell(3, 15, O)
	board.SetCell(4, 14, O)
	board.SetCell(5, 13, O)

	board.SetCell(2, 8, O)
	board.SetCell(3, 7, O)
	board.SetCell(4, 6, O)
	board.SetCell(5, 5, O)
	board.SetCell(6, 4, O)
	board.SetCell(7, 3, O)

	generateWinningPatterns(6)
	result = FindPattern(board, getWinningPatterns(O).winNow)

	sort.Sort(result)

	assertEqual(t, result[0], Interval{RLDiagonal, CellPosition{7, 3}, CellPosition{2, 8}})
	assertEqual(t, result[1], Interval{RLDiagonal, CellPosition{5, 13}, CellPosition{0, 18}})
	assertEqual(t, result[2], Interval{RLDiagonal, CellPosition{18, 13}, CellPosition{13, 18}})

}

func TestFindChainHorizVert(t *testing.T) {

	var board = NewBoard(19, 19)

	// horizontal

	board.SetCell(0, 0, X)
	board.SetCell(1, 0, X)
	board.SetCell(2, 0, X)
	board.SetCell(3, 0, X)

	board.SetCell(5, 0, X)
	board.SetCell(6, 0, X)
	board.SetCell(7, 0, X)
	board.SetCell(8, 0, X)

	board.SetCell(15, 18, X)
	board.SetCell(16, 18, X)
	board.SetCell(17, 18, X)
	board.SetCell(18, 18, X)

	// vertical

	board.SetCell(1, 0, X)
	board.SetCell(1, 1, X)
	board.SetCell(1, 2, X)
	board.SetCell(1, 3, X)

	board.SetCell(6, 0, X)
	board.SetCell(6, 1, X)
	board.SetCell(6, 2, X)
	board.SetCell(6, 3, X)

	board.SetCell(18, 15, X)
	board.SetCell(18, 16, X)
	board.SetCell(18, 17, X)
	board.SetCell(18, 18, X)

	generateWinningPatterns(4)
	result := FindPattern(board, getWinningPatterns(X).winNow)

	sort.Sort(result)

	assertEqual(t, result[0], Interval{horizontal, CellPosition{0, 0}, CellPosition{3, 0}})
	assertEqual(t, result[1], Interval{vertical, CellPosition{1, 0}, CellPosition{1, 3}})
	assertEqual(t, result[2], Interval{horizontal, CellPosition{5, 0}, CellPosition{8, 0}})
	assertEqual(t, result[3], Interval{vertical, CellPosition{6, 0}, CellPosition{6, 3}})
	assertEqual(t, result[4], Interval{vertical, CellPosition{18, 15}, CellPosition{18, 18}})
	assertEqual(t, result[5], Interval{horizontal, CellPosition{15, 18}, CellPosition{18, 18}})

}

func TestFindAllChains(t *testing.T) {

	var board = NewBoard(19, 19)

	board.SetCell(5, 0, X)
	board.SetCell(6, 0, X)
	board.SetCell(7, 0, X)
	board.SetCell(8, 0, X)
	board.SetCell(9, 0, X)

	board.SetCell(6, 10, X)
	board.SetCell(5, 11, X)
	board.SetCell(4, 12, X)
	board.SetCell(3, 13, X)
	board.SetCell(2, 14, X)

	board.SetCell(0, 18, X)
	board.SetCell(1, 18, X)
	board.SetCell(2, 18, X)

	board.SetCell(14, 0, X)
	board.SetCell(15, 0, X)
	board.SetCell(16, 0, X)

	generateWinningPatterns(5)
	result := FindPattern(board, getWinningPatterns(X).winNow)

	sort.Sort(result)

	if !reflect.DeepEqual(result, IntervalList {
		Interval{horizontal, CellPosition{5,0}, CellPosition{9,0}},
		Interval{RLDiagonal, CellPosition{6,10},CellPosition{2,14}}}) {
		t.Fatalf("Error in FindAllChains")
	}

	board = NewBoard(19, 19)

	generateWinningPatterns(3)
	result = FindPattern(board, getWinningPatterns(O).winNow)

	// return empty map if nothing has been found
	if len(result) > 0 {
		t.Fatalf("Error in FindAllChains")
	}

	board = NewBoard(6, 6)

	board.SetCell(1, 1, X)
	board.SetCell(2, 2, X)
	board.SetCell(3, 3, X)
	board.SetCell(4, 4, X)
	board.SetCell(5, 5, X)

	generateWinningPatterns(5)
	result = FindPattern(board, getWinningPatterns(X).winNow)

	assertEqual(t, result[0], Interval{LRDiagonal, CellPosition{1, 1}, CellPosition{5, 5}})
}

func TestShuffleIntSlice(t *testing.T) {

	inp := []int{1,2,3,4,5,6,7,8}
	ShuffleIntSlice(inp)
	assertEqual(t, inp, []int{8,6,2,1,7,5,4,3})

	inp = []int{1,2,3,4,5,6,7,8}
	ShuffleIntSlice(inp)
	assertEqual(t, inp, []int{7,1,4,2,3,5,8,6})

	inp = []int{}
	ShuffleIntSlice(inp)
	assertEqual(t, inp, []int{})

}

func TestMakeSearchPatterns(t *testing.T) {

	generateWinningPatterns(5)
	result := getWinningPatterns(X).winNow

	assertEqual(t, result, []Cell{X, X, X, X, X})

	generateWinningPatterns(3)
	result = getWinningPatterns(O).winNow

	assertEqual(t, result, []Cell{O, O, O})

}

func TestMonteCarloBestMove(t *testing.T) {

	var board = NewBoard(6, 6)

	board.SetCell(1, 1, X)
	board.SetCell(2, 2, X)
	board.SetCell(3, 3, X)
	board.SetCell(4, 4, X)

	generateWinningPatterns(5)
	result, _ := MonteCarloBestMove(board, AIOptions{AIPlayer: X, winSequenceLength: 5}, 6*6, 100, O)

	// Yeahh, X will certainly w in!
	assertEqual(t, result, CellPosition{0, 0})

	board = NewBoard(6, 6)

	board.SetCell(1, 1, O)
	board.SetCell(2, 2, O)
	board.SetCell(3, 3, O)
	board.SetCell(4, 4, O)

	generateWinningPatterns(5)
	result, _ = MonteCarloBestMove(board, AIOptions{AIPlayer: X, winSequenceLength: 5}, 6*6, 100, O)

	// Better luck next time, O was really good in this game
	assertEqual(t, result, CellPosition{0, 0})

	board = NewBoard(6, 6)

	board.SetCell(0, 1, O)
	board.SetCell(1, 1, O)
	board.SetCell(2, 1, O)
	board.SetCell(3, 1, O)

	generateWinningPatterns(5)
	result, _ = MonteCarloBestMove(board, AIOptions{AIPlayer: X, winSequenceLength: 5}, 6*6, 100, X)

	// The only move for X is 4, 1
	assertEqual(t, result, CellPosition{4, 1})

}

func TestMinMaxEval(t *testing.T) {

	// alpha-beta pruning disabled

	// test 1 - four in a row - WIN

	board := NewBoard(5, 5)

	board.SetCell(0, 0, X)
	board.SetCell(1, 1, X)
	board.SetCell(2, 2, X)

	options := AIOptions{ AIPlayer: X,
			winSequenceLength: 4,
			maxDepth: 3,
			useAlphaBeta: false }

	generateWinningPatterns(4)

	move, score := MinMaxEval(board, options, nil, LinearMove{0, options.AIPlayer}, options.maxDepth)
	col, row, _ := board.FromLinear(move)

	assertEqual(t, col, 3)
	assertEqual(t, row, 3)
	assertEqual(t, score, WON)

	// test 2 - five in a row - WIN

	board = NewBoard(5, 5)

	board.SetCell(0, 0, O)
	board.SetCell(0, 1, O)
	board.SetCell(0, 2, O)
	board.SetCell(0, 3, O)

	options = AIOptions{ AIPlayer: O,
			winSequenceLength: 5,
			maxDepth: 3,
			useAlphaBeta: false }

	generateWinningPatterns(5)

	move, score = MinMaxEval(board, options, nil, LinearMove{0, options.AIPlayer}, options.maxDepth)
	col, row, _ = board.FromLinear(move)

	assertEqual(t, col, 0)
	assertEqual(t, row, 4)
	assertEqual(t, score, WON)

	// test 3 - five in a row - LOST

	board = NewBoard(6, 6)

	board.SetCell(3, 3, O)
	board.SetCell(3, 4, O)
	board.SetCell(4, 3, O)

	options = AIOptions{ AIPlayer: X,
			winSequenceLength: 4,
			maxDepth: 4,
			useAlphaBeta: false }

	generateWinningPatterns(4)

	move, score = MinMaxEval(board, options, nil, LinearMove{0, options.AIPlayer}, options.maxDepth)

	assertEqual(t, score, LOST)

	// not enough analysis depth to detect lost position

	options = AIOptions{ AIPlayer: X,
			winSequenceLength: 4,
			maxDepth: 3,
			useAlphaBeta: false }

	generateWinningPatterns(4)

	move, score = MinMaxEval(board, options, nil, LinearMove{0, options.AIPlayer}, options.maxDepth)

	assertEqual(t, score, NOTHING)

	// alpha-beta pruning enabled

	// add some tests here!
}


// Some benchmarks

// Monte-Carlo benchmarks start >>>

func BenchmarkMonteCarloBestMove10(b *testing.B) {
	generateWinningPatterns(5)
	board := NewBoard(19, 19)
	for n := 0; n < b.N; n++ {
		MonteCarloBestMove(board, AIOptions{AIPlayer: X, winSequenceLength: 5}, 10 * 10, 10, X)
	}
}

func BenchmarkMonteCarloBestMove50(b *testing.B) {
	generateWinningPatterns(5)
	board := NewBoard(19, 19)
	for n := 0; n < b.N; n++ {
		MonteCarloBestMove(board, AIOptions{AIPlayer: X, winSequenceLength: 5}, 10 * 10, 50, X)
	}
}

func BenchmarkMonteCarloBestMove100(b *testing.B) {
	generateWinningPatterns(5)
	board := NewBoard(19, 19)
	for n := 0; n < b.N; n++ {
		MonteCarloBestMove(board, AIOptions{AIPlayer: X, winSequenceLength: 5}, 10 * 10, 100, X)
	}
}

// Monte-Carlo benchmarks end <<<

// FindPattern benchmarks start >>>

func BenchmarkFindPattern(b *testing.B) {
	board := NewBoard(10, 10)
	for n := 0; n < b.N; n++ {
		FindPattern(board, []Cell{X, X, X, X, X})
	}
}

// FindPattern benchmarks end <<<

// MinMax benchmarks start >>>

func BenchmarkMinMaxEval6x6_1(b *testing.B) {

	board := NewBoard(6, 6)

	options := AIOptions{ AIPlayer: X,
		              winSequenceLength: 3,
		              maxDepth: 2,
		              useAlphaBeta: false }

	for n := 0; n < b.N; n++ {
		MinMaxEval(board, options, nil, LinearMove{0, options.AIPlayer}, options.maxDepth)
	}
}

// MinMax benchmarks end <<<