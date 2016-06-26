package misc

type ScanDirection uint8

type CellPosition struct {
	Col, Row int
}

// Interval represents indexes of start and end of an n-length chain`
type Interval struct {

	Direction  ScanDirection

	From CellPosition
	To CellPosition

}

// Unfolds interval encoded by Interval data structure
func (interval Interval) Unfold() []CellPosition {

	result := []CellPosition{}


	switch interval.Direction {
	case LRDiagonal:
		result = make([]CellPosition, interval.To.Col - interval.From.Col + 1)
		for idx := 0; idx < interval.To.Col - interval.From.Col + 1; idx++ {
			result[idx] = CellPosition{interval.From.Col + idx, interval.From.Row + idx}
		}
		break
	case RLDiagonal:
		result = make([]CellPosition, interval.From.Col - interval.To.Col + 1)
		for idx := 0; idx < interval.From.Col - interval.To.Col + 1; idx++ {
			result[idx] = CellPosition{interval.From.Col - idx, interval.From.Row + idx}
		}
		break
	case horizontal:
		result = make([]CellPosition, interval.To.Col - interval.From.Col + 1)
		for idx := 0; idx < interval.To.Col - interval.From.Col + 1; idx++ {
			result[idx] = CellPosition{interval.From.Col + idx, interval.From.Row}
		}
		break
	case vertical:
		result = make([]CellPosition, interval.To.Row - interval.From.Row + 1)
		for idx := 0; idx < interval.To.Row - interval.From.Row + 1; idx++ {
			result[idx] = CellPosition{interval.From.Col, interval.To.Row + idx}
		}
		break
	}

	return result

}

type AIOptions struct {

	AIPlayer Cell
	winSequenceLength int
	maxDepth int

}

const (
	horizontal ScanDirection = iota
	vertical

	// LR means that we are scanning from the left side of the board
	// towards its right side
	LRDiagonal

	// RL means that we are scanning from the right side of the board
	// towards its left side
	RLDiagonal
)

const (
	AI_WINS = 10
	AI_LOSES = -10
	DRAW = 0
)

