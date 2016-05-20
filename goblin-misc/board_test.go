package misc

import (
	"testing"

	"github.com/nsf/termbox-go"
	"fmt"
)

func cmpSlices(a, b []Cell) bool {
	for i, v := range a { if v != b[i] { return false } }
	return true
}

func TestDiagonalSlices(t *testing.T) {

	var board = NewBoard(19, 19, 10, 10, termbox.ColorBlack, termbox.ColorBlue,
		termbox.ColorRed, termbox.ColorBlack)

	board.SetCell(0, 0, X)
	board.SetCell(1, 1, X)
	board.SetCell(5, 5, X)
	board.SetCell(18, 18, X)

	// Test generic diagonal slice generator

	result1 := board.GetDiagonalSliceXY(0, 0, 19, 19)

	if !cmpSlices(result1, []Cell{X, X, E, E, E, X, E, E, E, E, E, E, E, E, E, E, E, E, X}) {
		t.Fail()
	}

	board.SetCell(0, 14, O)
	board.SetCell(1, 15, O)
	board.SetCell(2, 16, O)
	board.SetCell(3, 17, O)
	board.SetCell(4, 18, X)

	result2 := board.GetDiagonalSliceXY(0, 14, 5, 19)

	if !cmpSlices(result2, []Cell{O, O, O, O, X}) {
		t.Fail()
	}

	board.SetCell(16, 0, X)
	board.SetCell(17, 1, O)
	board.SetCell(18, 2, X)

	result3 := board.GetDiagonalSliceXY(16, 0, 19, 3)

	if !cmpSlices(result3, []Cell{X, O, X}) {
		t.Fail()
	}

	board.SetCell(14, 18, O)
	board.SetCell(18, 14, O)

	result4 := board.GetDiagonalSliceXY(18, 14, 15, 19)

	if !cmpSlices(result4, []Cell{O, E, E, E, O}) {
		t.Fail()
	}


	// Test Left->Right diagonal slicer

	newResult1 := board.GetLRDiagonal(10, 10)
	if !cmpSlices(result1, newResult1) {
		t.Fail()
	}

	newResult2 := board.GetLRDiagonal(2, 16)
	if !cmpSlices(result2, newResult2) {
		t.Fail()
	}

	newResult3 := board.GetLRDiagonal(16, 0)
	if !cmpSlices(result3, newResult3) {
		t.Fail()
	}

	// Test Right->Left diagonal slicer


	fmt.Println(result1)

}
