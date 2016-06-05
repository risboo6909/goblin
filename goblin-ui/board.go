package ui

import (
	"math"

	"github.com/nsf/termbox-go"
	"github.com/risboo6909/goblin/goblin-misc"
)

// Board visual attributes
type BoardAttrs struct {
	BoardColor, BoardBg   termbox.Attribute
	LabelsColor, LabelsBg termbox.Attribute
}

type DrawableBoard struct {

	*misc.BoardDescription

	// upper-left corner position of a board
	X, Y int

	BoardAttrs
}

func NewBoard(cellsHoriz, cellsVert, x, y int, boardColor, boardBg, labelsColor,
	labelsBg termbox.Attribute) *DrawableBoard {
	return &DrawableBoard{misc.NewBoard(cellsHoriz, cellsVert),
		              x, y,
		              BoardAttrs{boardColor, boardBg, labelsColor, labelsBg}}
}

func modN(n float64) func(int) float64 {
	return func(m int) float64 {
		return math.Mod(float64(m), n)
	}
}

var mod4 = modN(4)
var mod2 = modN(2)

func getScrX(board *DrawableBoard, col int) int {
	return board.X + 2 + col*4
}

func getScrY(board *DrawableBoard, row int) int {
	return board.Y + 1 + row*2
}

// drawHorizLine draws horizontal board lines
func drawHorizLine(color, bgcolor termbox.Attribute, x, y, width int) {

	for i := 0; i < width; i++ {
		if mod4(i) != 3 {
			termbox.SetCell(x+i, y, '─', color, bgcolor)
		} else {
			termbox.SetCell(x+i, y, '┼', color, bgcolor)
		}
	}

}

// drawVertLine draw vertical board lines
func drawVertLine(color, bgcolor termbox.Attribute, x, y, height int, s rune) {
	for i := 0; i < height; i++ {
		termbox.SetCell(x, y+i, s, color, bgcolor)
	}
}

// drawTopAndBottom draws upper or lower parts of a board
func drawTopAndBottom(x, y int, p *DrawableBoard, left, middle, right rune, labels bool) {

	if labels {

		var letter = 'A'

		printTb(x-2, y-1, p.LabelsColor, p.LabelsBg, "  ")
		for i := 0; i <= p.GetWidth(); i++ {

			if mod4(i) == 0 {
				termbox.SetCell(x+i, y-1, letter, p.LabelsColor, p.LabelsBg)
				letter++
			} else {
				termbox.SetCell(x+i, y-1, ' ', p.LabelsColor, p.LabelsBg)
			}

		}
	}

	termbox.SetCell(x, y, left, p.BoardColor, p.BoardBg)

	for i := 1; i < p.GetWidth(); i++ {

		xOffset := x + i

		if mod4(i) == 0 {
			termbox.SetCell(xOffset, y, middle, p.BoardColor, p.BoardBg)
			if labels {
				drawVertLine(p.BoardColor, p.BoardBg, xOffset, y+1, p.GetHeight(), '│')
			}

		} else {
			termbox.SetCell(xOffset, y, '─', p.BoardColor, p.BoardBg)
			if labels {
				drawVertLine(p.BoardColor, p.BoardBg, xOffset, y+1, p.GetHeight(), ' ')
			}
		}
	}

	termbox.SetCell(x+p.GetWidth(), y, right, p.BoardColor, p.BoardBg)
}

// drawLeftAndRight draws left and right parts of a board
func drawLeftAndRight(x, y int, p *DrawableBoard, middle rune, labels bool) {

	if labels {

		var index = 1

		for i := 0; i <= p.GetHeight(); i++ {

			if mod2(i) == 0 {
				printfTb(x-2, y+i, p.LabelsColor, p.LabelsBg, "%2d", index)
				index++
			} else {
				printTb(x-2, y+i, p.LabelsColor, p.LabelsBg, "  ")
			}
		}
	}

	for i := 1; i < p.GetHeight(); i++ {

		yOffset := y + i

		if mod2(i) == 0 {
			termbox.SetCell(x, yOffset, middle, p.BoardColor, p.BoardBg)
			if labels {
				drawHorizLine(p.BoardColor, p.BoardBg, x+1, yOffset, p.GetWidth())
			}
		} else {
			termbox.SetCell(x, yOffset, '│', p.BoardColor, p.BoardBg)
		}

	}

}

func fillBoard(board *DrawableBoard) {
	for i := 0; i < board.CellsHoriz; i++ {
		for j := 0; j < board.CellsVert; j++ {

			scrX := getScrX(board, i)
			scrY := getScrY(board, j)

			if board.GetCell(i, j) == misc.X {
				termbox.SetCell(scrX, scrY, misc.X, termbox.ColorYellow|termbox.AttrBold, board.BoardBg)

			} else if board.GetCell(i, j) == misc.O {
				termbox.SetCell(scrX, scrY, misc.O, termbox.ColorGreen|termbox.AttrBold, board.BoardBg)
			}

		}
	}
}

func drawCursor(board *DrawableBoard, cursor *Cursor) {
	val := board.GetCell(cursor.Col, cursor.Row)
	termbox.SetCell(getScrX(board, cursor.Col), getScrY(board, cursor.Row),
		rune(val), cursor.FgColor, cursor.BgColor)
}

// DrawBoard draws ASCII game board
func DrawBoard(board *DrawableBoard, cursor *Cursor) {

	x := board.X + 2
	y := board.Y + 1

	// draw top and bottom parts
	drawTopAndBottom(x, y, board, '┌', '┬', '┐', true)
	drawTopAndBottom(x, y+board.GetHeight(), board, '└', '┴', '┘', false)

	// draw left and right parts
	drawLeftAndRight(x, y, board, '├', true)
	drawLeftAndRight(x+board.GetWidth(), y, board, '┤', false)

	fillBoard(board)
	drawCursor(board, cursor)
}
