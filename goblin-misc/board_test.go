package misc

import (
	"testing"

	"github.com/nsf/termbox-go"
	"fmt"
)

func TestDiagonalSlices(t *testing.T) {

	var board = NewBoard(19, 19, 10, 10, termbox.ColorBlack, termbox.ColorBlue,
		termbox.ColorRed, termbox.ColorBlack)

	board.SetCell(0, 0, X)
	board.SetCell(1, 1, X)
	board.SetCell(5, 5, X)
	board.SetCell(18, 18, X)

	// Test generic diagonal slice generator

	result := board.GetDiagonalSliceXY(0, 0, 19, 19)

	if result != nil && len(result) == 19 {
		if result[0] != X || result[1] != X || result[5] != X || result[18] != X {
			t.Fail()
		}
		if result[2] != EMPTY || result[3] != EMPTY || result[4] != EMPTY || result[17] != EMPTY {
			t.Fail()
		}
	}

	// Test convenience slicers

	result = board.GetRightDiagonal(10, 10)

	if result != nil && len(result) == 19 {
		if result[0] != X || result[1] != X || result[5] != X || result[18] != X {
			t.Fail()
		}
		if result[2] != EMPTY || result[3] != EMPTY || result[4] != EMPTY || result[17] != EMPTY {
			t.Fail()
		}
	}

	result = board.GetRightDiagonal(2, 16)
	fmt.Println(result)

}
