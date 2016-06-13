package misc

import (
	"testing"

	"reflect"
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

	result := FindPattern(board, 4, X)

	assertEqual(t, result[0], Interval{LRDiagonal, 2, 2, 5, 5})
	assertEqual(t, result[1], Interval{LRDiagonal, 1, 1, 4, 4})
	assertEqual(t, result[2], Interval{LRDiagonal, 3, 3, 6, 6})
	assertEqual(t, result[3], Interval{LRDiagonal, 7, 3, 10, 6})
	assertEqual(t, result[4], Interval{LRDiagonal, 6, 2, 9, 5})
	assertEqual(t, result[5], Interval{LRDiagonal, 8, 4, 11, 7})
	assertEqual(t, result[6], Interval{LRDiagonal, 0, 15, 3, 18})

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

	result = FindPattern(board, 6, O)

	assertEqual(t, result[0], Interval{RLDiagonal, 7, 3, 2, 8})
	assertEqual(t, result[1], Interval{RLDiagonal, 8, 2, 3, 7})
	assertEqual(t, result[2], Interval{RLDiagonal, 6, 4, 1, 9})
	assertEqual(t, result[3], Interval{RLDiagonal, 5, 13, 0, 18})
	assertEqual(t, result[4], Interval{RLDiagonal, 6, 12, 1, 17})
	assertEqual(t, result[5], Interval{RLDiagonal, 18, 13, 13, 18})

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

	board.SetCell(0, 0, X)
	board.SetCell(0, 1, X)
	board.SetCell(0, 2, X)
	board.SetCell(0, 3, X)

	board.SetCell(5, 0, X)
	board.SetCell(5, 1, X)
	board.SetCell(5, 2, X)
	board.SetCell(5, 3, X)

	board.SetCell(18, 15, X)
	board.SetCell(18, 16, X)
	board.SetCell(18, 17, X)
	board.SetCell(18, 18, X)

	result := FindPattern(board, 4, X)

	assertEqual(t, result[0], Interval{horizontal, 0, 0, 3, 0})
	assertEqual(t, result[1], Interval{horizontal, 5, 0, 8, 0})
	assertEqual(t, result[2], Interval{horizontal, 4, 0, 7, 0})
	assertEqual(t, result[3], Interval{horizontal, 1, 0, 4, 0})
	assertEqual(t, result[4], Interval{horizontal, 6, 0, 9, 0})
	assertEqual(t, result[5], Interval{horizontal, 15, 18, 18, 18})
	assertEqual(t, result[6], Interval{horizontal, 14, 18, 17, 18})

	assertEqual(t, result[7], Interval{vertical, 0, 0, 0, 3})
	assertEqual(t, result[8], Interval{vertical, 0, 1, 0, 4})
	assertEqual(t, result[9], Interval{vertical, 5, 0, 5, 3})
	assertEqual(t, result[10], Interval{vertical, 5, 1, 5, 4})
	assertEqual(t, result[11], Interval{vertical, 18, 15, 18, 18})
	assertEqual(t, result[12], Interval{vertical, 18, 14, 18, 17})

}

func TestFindAllChains(t *testing.T) {

	var board = NewBoard(19, 19)

	board.SetCell(5, 0, X)
	board.SetCell(6, 0, X)
	board.SetCell(7, 0, X)
	board.SetCell(8, 0, X)

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

	result := FindPattern(board, 5, X)

	if !reflect.DeepEqual(result, []Interval {
		Interval{horizontal, 4,0,8,0}, Interval{horizontal, 5,0,9,0},
		Interval{horizontal, 12,0,18,0}, Interval{RLDiagonal, 6,10,2,14},
		Interval{RLDiagonal, 7,9,3,13}, Interval{RLDiagonal, 5,11,1,15}}) {
		t.Fatalf("Error in FindAllChains")
	}

	board = NewBoard(19, 19)

	result = FindPattern(board, 3, O)

	// return empty map if nothing has been found
	if len(result) > 0 {
		t.Fatalf("Error in FindAllChains")
	}

	board = NewBoard(6, 6)

	board.SetCell(2, 2, X)
	board.SetCell(3, 3, X)
	board.SetCell(4, 4, X)
	board.SetCell(5, 5, X)

	result = FindPattern(board, 5, X)

	assertEqual(t, result[0], Interval{LRDiagonal, 1, 1, 5, 5})

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

	result := MakePatterns(5, X)

	assertEqual(t, result[0], []Cell{X, X, X, X, X})
	assertEqual(t, result[1], []Cell{E, X, X, X, X})
	assertEqual(t, result[2], []Cell{X, X, X, X, E})
	assertEqual(t, result[3], []Cell{E, E, X, X, X, E, E})

	result = MakePatterns(3, O)

	assertEqual(t, result[0], []Cell{O, O, O})
	assertEqual(t, result[1], []Cell{E, O, O})
	assertEqual(t, result[2], []Cell{O, O, E})
	assertEqual(t, result[3], []Cell{E, E, O, E, E})

}

func TestMonteCarloEval(t *testing.T) {

	var board = NewBoard(6, 6)

	board.SetCell(2, 2, X)
	board.SetCell(3, 3, X)
	board.SetCell(4, 4, X)
	board.SetCell(5, 5, X)

	// Winning position for X
	assertEqual(t, MonteCarloEval(board, AIOptions{AIPlayer: X, winSequenceLength: 5}, 100, 6*6, O), 1.0)

	board = NewBoard(6, 6)

	board.SetCell(2, 2, O)
	board.SetCell(3, 3, O)
	board.SetCell(4, 4, O)
	board.SetCell(5, 5, O)
	board.SetCell(6, 6, O)

	// O wins here, no chances for X
	assertEqual(t, MonteCarloEval(board, AIOptions{AIPlayer: X, winSequenceLength: 5}, 100, 6*6, X), 0.0)

	board = NewBoard(6, 6)

	board.SetCell(0, 0, O)
	board.SetCell(1, 1, O)
	board.SetCell(2, 2, O)

	// X still have a chance to win
	assertEqual(t, roundFloat(MonteCarloEval(board, AIOptions{AIPlayer: X, winSequenceLength: 5}, 100, 6*6, O), 2), 0.26)

	board = NewBoard(6, 6)

	// AI (X) moves first so it has a small advantage
	assertEqual(t, roundFloat(MonteCarloEval(board, AIOptions{AIPlayer: X, winSequenceLength: 5}, 1000, 6*6, X), 2), 0.54)

}
