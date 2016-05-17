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

	if result1 != nil && len(result1) == 19 {
		if result1[0] != X || result1[1] != X || result1[5] != X || result1[18] != X {
			t.Fail()
		}
		if result1[2] != EMPTY || result1[3] != EMPTY || result1[4] != EMPTY || result1[17] != EMPTY {
			t.Fail()
		}
	}

	board.SetCell(0, 14, O)
	board.SetCell(1, 15, O)
	board.SetCell(2, 16, O)
	board.SetCell(3, 17, O)
	board.SetCell(4, 18, X)

	result2 := board.GetDiagonalSliceXY(0, 14, 5, 19)

	if result2[0] != O || result2[1] != O || result2[2] != O || result2[3] != O || result2[4] != X {
		t.Fail()
	}

	board.SetCell(16, 0, X)
	board.SetCell(17, 1, O)
	board.SetCell(18, 2, X)

	result3 := board.GetDiagonalSliceXY(16, 0, 19, 3)

	// Test Left->Right slide

	newResult1 := board.GetRightDiagonal(10, 10)
	if !cmpSlices(result1, newResult1) {
		t.Fail()
	}

	newResult2 := board.GetRightDiagonal(2, 16)
	if !cmpSlices(result2, newResult2) {
		t.Fail()
	}

	newResult3 := board.GetRightDiagonal(16, 0)
	if !cmpSlices(result3, newResult3) {
		t.Fail()
	}

	fmt.Println(result3)

}
