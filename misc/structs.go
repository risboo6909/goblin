package misc


type ScanDirection uint8

type Move struct {
	col, row int
}

// Interval represents indexes of start and end of an n-length chain`
type Interval struct {

	Direction  ScanDirection

	Col1, Row1 int
	Col2, Row2 int

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

