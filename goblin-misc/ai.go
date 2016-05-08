package misc

// Interval represents indexes of start and end of an n-length chain`
type Interval struct {
	startCol, startRow int
	endCol, endRow     int
}

func scanLine(line []Cell, rowNo, chainLen int, player rune) []Interval {
	var (
		result  []Interval
		counter = 0
	)

	for idx, cell := range line {
		if player == cell.Val {
			if counter++; counter == chainLen {
				result = append(result, Interval{startCol: idx - chainLen + 1, startRow: rowNo,
					endCol: idx, endRow: rowNo})
			}
		} else {
			counter = 0
		}
	}
	return result
}

// FindChain finds vertical, horizontal or diagonal chains of
// successive cells with the same content as a slice of Intervals
func FindChain(board *BoardDescription, chainLen int, player rune) []Interval {

	// scan horizontal first
	var matchHoriz []Interval

	for i := 0; i < board.CellsVert; i++ {
		// get slice of each row
		row := board.GetHorizSlice(i, 0, board.CellsHoriz-1)
		// and scan for a chain
		tmp := scanLine(row, i, chainLen, player)

		// if there was a positive result copy it to the global result
		if tmp != nil {
			matchHoriz = append(matchHoriz, tmp...)
		}
	}
	return matchHoriz
}

// FindAllChains finds all the chains of length in interval [chainLenMin..chainLenMax]
func FindAllChains(board *BoardDescription, chainLenMin, chainLenMax int, player rune) {

}
