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

	result := FindChain(board, 4, X)

	if result[0] != (Interval{LRDiagonal, 2, 2, 5, 5}) {
		t.Fail()
	}
	if result[1] != (Interval{LRDiagonal, 7, 3, 10, 6}) {
		t.Fail()
	}
	if result[2] != (Interval{LRDiagonal, 0, 15, 3, 18}) {
		t.Fail()
	}

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

	result = FindChain(board, 6, O)

	if result[0] != (Interval{RLDiagonal, 2, 8, 7, 3}) {
		t.Fail()
	}
	if result[1] != (Interval{RLDiagonal, 0, 18, 5, 13}) {
		t.Fail()
	}
	if result[2] != (Interval{RLDiagonal, 13, 18, 18, 13}) {
		t.Fail()
	}

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

	result := FindChain(board, 4, X)

	if result[0] != (Interval{horizontal, 0, 0, 3, 0}) {
		t.Fail()
	}
	if result[1] != (Interval{horizontal, 5, 0, 8, 0}) {
		t.Fail()
	}
	if result[2] != (Interval{horizontal, 15, 18, 18, 18}) {
		t.Fail()
	}

	if result[3] != (Interval{vertical, 0, 0, 0, 3}) {
		t.Fail()
	}
	if result[4] != (Interval{vertical, 5, 0, 5, 3}) {
		t.Fail()
	}
	if result[5] != (Interval{vertical, 18, 15, 18, 18}) {
		t.Fail()
	}
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

	board.SetCell(16, 0, X)
	board.SetCell(17, 0, X)
	board.SetCell(18, 0, X)

	result := FindAllChains(board, 2, 5, X)

	if !reflect.DeepEqual(result, map[int][]Interval {
		3:[]Interval{Interval{horizontal, 16,0,18,0}, Interval{horizontal, 0,18,2,18}},
		4:[]Interval{Interval{horizontal, 5,0,8,0}}, 5:[]Interval{Interval{RLDiagonal, 2,14,6,10}}}) {
		t.Fail()
	}

	board = NewBoard(19, 19)

	result = FindAllChains(board, 3, 4, O)

	// return empty map if nothing has been found
	if len(result) > 0 {
		t.Fail()
	}
}

func TestShuffleIntSlice(t *testing.T) {

	inp := []int{1,2,3,4,5,6,7,8}

	ShuffleIntSlice(inp)

	if !cmpSlices(inp, []int{8,6,2,1,7,5,4,3}) {
		t.Fail()
	}

	inp = []int{1,2,3,4,5,6,7,8}

	ShuffleIntSlice(inp)

	if !cmpSlices(inp, []int{7,1,4,2,3,5,8,6}) {
		t.Fail()
	}

	inp = []int{}

	ShuffleIntSlice(inp)

	if !cmpSlices(inp, []int{}) {
		t.Fail()
	}

}

func TestMakeSearchPatterns(t *testing.T) {

	result := MakeSearchPatterns(5, X)

	if !cmpSlices(result[0], []Cell{88, 88, 88, 88, 88}) {
		t.Fail()
	}

	if !cmpSlices(result[1], []Cell{32, 88, 88, 88, 88}) {
		t.Fail()
	}

	if !cmpSlices(result[2], []Cell{88, 88, 88, 88, 32}) {
		t.Fail()
	}

	if !cmpSlices(result[3], []Cell{32, 32, 88, 88, 88, 32, 32}) {
		t.Fail()
	}

}

//func TestMonteCarloEval(t *testing.T) {
//
//	var board = NewBoard(5, 5)
//
//	board.SetCell(1, 1, X)
//	board.SetCell(2, 2, X)
//	board.SetCell(3, 3, X)
//	//board.SetCell(3, 3, O)
//
//	fmt.Println(MonteCarloEval(board, AIOptions{AIPlayer: X, winSequenceLength: 5, maxDepth: 10}, 10, 19*19, O))
//}
