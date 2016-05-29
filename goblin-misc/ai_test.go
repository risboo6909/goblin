package misc

import (
	"testing"

	"github.com/nsf/termbox-go"
	"fmt"
)


func TestFindChainDiagonal(t *testing.T) {

	var board = NewBoard(19, 19, 10, 10, termbox.ColorBlack, termbox.ColorBlue,
		termbox.ColorRed, termbox.ColorBlack)

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

	if result[0] != (Interval{2, 2, 5, 5}) {
		t.Fail()
	}
	if result[1] != (Interval{7, 3, 10, 6}) {
		t.Fail()
	}
	if result[2] != (Interval{0, 15, 3, 18}) {
		t.Fail()
	}

	// Right-to-left diagonal sequences

	board = NewBoard(19, 19, 10, 10, termbox.ColorBlack, termbox.ColorBlue,
		termbox.ColorRed, termbox.ColorBlack)

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

	fmt.Println(result)
}


func TestFindChainHorizVert(t *testing.T) {

	var board = NewBoard(19, 19, 10, 10, termbox.ColorBlack, termbox.ColorBlue,
		termbox.ColorRed, termbox.ColorBlack)

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

//	fmt.Printf("%v", result)

	if result != nil && len(result) == 6 {

		if result[0] != (Interval{0, 0, 3, 0}) {
			t.Fail()
		}
		if result[1] != (Interval{5, 0, 8, 0}) {
			t.Fail()
		}
		if result[2] != (Interval{15, 18, 18, 18}) {
			t.Fail()
		}

		if result[3] != (Interval{0, 0, 0, 3}) {
			t.Fail()
		}
		if result[4] != (Interval{5, 0, 5, 3}) {
			t.Fail()
		}
		if result[5] != (Interval{18, 15, 18, 18}) {
			t.Fail()
		}

	} else {
		t.Fail()
	}

}
