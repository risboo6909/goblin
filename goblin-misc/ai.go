package misc

// Interval represents indexes of start and end of an n-length chain`
type Interval struct {
	startCol, startRow int
	endCol, endRow     int
}

type ScanDirection uint8

const (
	horizontal = iota
	vertical
	diagonal
)

func scanLine(line []Cell, cellNo, chainLen int, player Cell, direction ScanDirection) []Interval {
	var (
		result  []Interval
		counter = 0
	)

	for idx, cell := range line {
		if player == cell {
			if counter++; counter == chainLen {

				if direction == horizontal {
					result = append(result, Interval{startCol: idx - chainLen + 1, startRow: cellNo,
						endCol: idx, endRow: cellNo})

				} else if direction == vertical {
					result = append(result, Interval{startRow: idx - chainLen + 1, startCol: cellNo,
						endRow: idx, endCol: cellNo})
				}

			}
		} else {
			counter = 0
		}
	}

	return result
}

// FindChain finds vertical, horizontal or diagonal chains of
// successive cells with the same content as a slice of Intervals
func FindChain(board *BoardDescription, chainLen int, player Cell) []Interval {

	var (
		matchHoriz []Interval
		matchVert  []Interval
	)

	// scan horizontal first

	for i := 0; i < board.CellsVert; i++ {
		// get slice of each row
		row := board.GetHorizSlice(i, 0, board.CellsHoriz-1)
		// and scan for a chain
		tmp := scanLine(row, i, chainLen, player, horizontal)

		// if there was a positive result copy it to the global result
		if tmp != nil {
			matchHoriz = append(matchHoriz, tmp...)
		}
	}

	// then vertical

	for i := 0; i < board.CellsHoriz; i++ {
		// get slice of each column
		column := board.GetVertSlice(i, 0, board.CellsVert-1)
		// and scan for a chain
		tmp := scanLine(column, i, chainLen, player, vertical)

		// if there was a positive result copy it to the global result
		if tmp != nil {
			matchVert = append(matchVert, tmp...)
		}
	}

	return append(matchHoriz, matchVert...)
}

// FindAllChains finds all the chains of length in interval [chainLenMin..chainLenMax]
func FindAllChains(board *BoardDescription, chainLenMin, chainLenMax int, player rune) {

}
