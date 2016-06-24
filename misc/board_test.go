package misc

import (
	"testing"

	"math/rand"
)

func TestDiagonalSlices(t *testing.T) {

	var board = NewBoard(19, 19)

	board.SetCell(0, 0, X)
	board.SetCell(1, 1, X)
	board.SetCell(5, 5, X)
	board.SetCell(18, 18, X)

	// Test generic diagonal slice generator

	result1 := board.GetDiagonalSliceXY(0, 0, 18, 18)

	assertEqual(t, result1, []Cell{X, X, E, E, E, X, E, E, E, E, E, E, E, E, E, E, E, E, X})

	board.SetCell(0, 14, O)
	board.SetCell(1, 15, O)
	board.SetCell(2, 16, O)
	board.SetCell(3, 17, O)
	board.SetCell(4, 18, X)

	result2 := board.GetDiagonalSliceXY(0, 14, 4, 18)

	assertEqual(t, result2, []Cell{O, O, O, O, X})

	board.SetCell(16, 0, X)
	board.SetCell(17, 1, O)
	board.SetCell(18, 2, X)

	result3 := board.GetDiagonalSliceXY(16, 0, 18, 2)

	assertEqual(t, result3, []Cell{X, O, X})

	board.SetCell(14, 18, O)
	board.SetCell(18, 14, O)

	result4 := board.GetDiagonalSliceXY(18, 14, 14, 18)

	assertEqual(t, result4, []Cell{O, E, E, E, O})

	board.SetCell(0, 3, X)

	result5 := board.GetDiagonalSliceXY(3, 0, 0, 3)

	assertEqual(t, result5, []Cell{E, E, E, X})


	// Test Left->Right diagonal slicer

	newResult1 := board.GetLRDiagonal(10, 10)
	assertEqual(t, result1, newResult1)

	newResult2 := board.GetLRDiagonal(0, 14)
	assertEqual(t, result2, newResult2)

	newResult3 := board.GetLRDiagonal(18, 2)
	assertEqual(t, result3, newResult3)


	// Test Right->Left diagonal slicer

	newResult4 := board.GetRLDiagonal(16, 16)
	assertEqual(t, result4, newResult4)

	newResult5 := board.GetRLDiagonal(3, 0)
	assertEqual(t, result5, newResult5)


	// Some randomized tests

	for i := 0; i < 10000; i++ {

		board := GetRandomizedBoard(19, 19, 60.0)
		startCol, startRow := rand.Intn(19), rand.Intn(19)

		stCol1, stRow1, endCol1, endRow1 := board.GetBounds(startCol, startRow, LeftToRight)
		stCol2, stRow2, endCol2, endRow2 := board.GetBounds(startCol, startRow, RightToLeft)

		lrDiagonal1 := board.GetDiagonalSliceXY(stCol1, stRow1, endCol1, endRow1)
		rlDiagonal1 := board.GetDiagonalSliceXY(stCol2, stRow2, endCol2, endRow2)

		lrDiagonal2 := board.GetLRDiagonal(startCol, startRow)
		rlDiagonal2 := board.GetRLDiagonal(startCol, startRow)

		assertEqual(t, lrDiagonal1, lrDiagonal2)
		assertEqual(t, rlDiagonal1, rlDiagonal2)

	}

}
