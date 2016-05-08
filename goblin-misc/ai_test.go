package misc

import (
	"testing"

	"github.com/nsf/termbox-go"
)

func TestFindChain(t *testing.T) {
	var board = NewBoard(19, 19, 10, 10, termbox.ColorBlack, termbox.ColorBlue,
		termbox.ColorRed, termbox.ColorBlack)

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

	result := FindChain(board, 4, X)

	if result != nil && len(result) == 3 {
		//fmt.Printf("%v", result)
		if result[0] != (Interval{0, 0, 3, 0}) {
			t.Fail()
		}
		if result[1] != (Interval{5, 0, 8, 0}) {
			t.Fail()
		}
		if result[2] != (Interval{15, 18, 18, 18}) {
			t.Fail()
		}

	} else {
		t.Fail()
	}

}
