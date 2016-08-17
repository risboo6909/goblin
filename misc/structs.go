package misc


// Mimic python set
type Set map[interface{}]bool

type IntFloatPair struct {
	Fst int
	Snd float64
}


type IntFloatPairs []IntFloatPair

func (slice IntFloatPairs) Len() int {
	return len(slice)
}

func (slice IntFloatPairs) Less(i, j int) bool {
	// descending order by default
	return slice[i].Snd > slice[j].Snd
}

func (slice IntFloatPairs) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}


type LinearMove struct {
	position int
	player Cell
}

type ScanDirection uint8

type CellPosition struct {
	Col, Row int
}

// Interval represents indexes of start and end of an n-length chain

type Interval struct {

	Direction  ScanDirection

	From CellPosition
	To CellPosition

}

// List of intervals is sortable by coordinates

type IntervalList []Interval

func (slice IntervalList) Len() int {
	return len(slice)
}

func (slice IntervalList) Less(i, j int) bool {
	if slice[i].From.Row < slice[j].From.Row {
		return true
	} else if slice[i].From.Row == slice[j].From.Row {
		if slice[i].From.Col < slice[j].From.Col{
			return true
		}
	}
	return false
}

func (slice IntervalList) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
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
			result[idx] = CellPosition{interval.From.Col, interval.From.Row + idx}
		}
		break
	}

	return result

}

type AIOptions struct {

	AIPlayer Cell
	winSequenceLength int
	maxDepth int
	useAlphaBeta bool
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
	WON = 10
	LOST = -WON
	NOTHING = 0
)

type PatternType struct {
	winNow []Cell
	winInAMove []Cell
}

var winningPatternsX PatternType
var winningPatternsO PatternType
