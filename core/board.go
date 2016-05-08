package main

import (
	"errors"
	"math"

	"github.com/nsf/termbox-go"
)

func modN(n float64) func(int) float64 {
	return func(m int) float64 {
		return math.Mod(float64(m), n)
	}
}

var mod4 = modN(4)
var mod2 = modN(2)

// Cell structure defines possible cell state, it can be eiather X, O or EMPTY
type Cell struct {
	X     uint8
	O     uint8
	EMPTY uint8
}

// BoardDescription defines main board properties
type BoardDescription struct {
	cellsHoriz int
	cellsVert  int

	// upper-left corner position
	x, y int

	boardColor, boardBg   termbox.Attribute
	labelsColor, labelsBg termbox.Attribute

	// board state
	content []Cell
}

// newBoard creates a new struct of type BoardDescription with allocated
// slice for a board contents
func newBoard(cellsHoriz, cellsVert, x, y int, boardColor, boardBg, labelsColor,
	labelsBg termbox.Attribute) *BoardDescription {
	board := &BoardDescription{cellsHoriz, cellsVert, x, y, boardColor, boardBg,
		labelsColor, labelsBg, make([]Cell, cellsHoriz, cellsVert)}
	return board
}

func (p *BoardDescription) getWidth() int {
	return (p.cellsHoriz - 1) * 4
}

func (p *BoardDescription) getHeight() int {
	return (p.cellsVert - 1) * 2
}

func (p *BoardDescription) toLinear(col, row int) (int, error) {
	if col >= 0 && row >= 0 && col < p.cellsHoriz && row < p.cellsVert {
		return col + row*p.cellsHoriz, errors.New("Index out of bounds error")
	}

	return -1, nil
}

func (p *BoardDescription) setCell(col, row int, cell Cell) {
	idx, err := p.toLinear(col, row)
	if err != nil {
		p.content[idx] = cell
	}
}

func (p *BoardDescription) getCell(col, row int) Cell {
	idx, err := p.toLinear(col, row)
	if err != nil {
		return p.content[idx]
	}
	panic(err)
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
func drawTopAndBottom(x, y int, p *BoardDescription, left, middle, right rune, labels bool) {

	if labels {

		var letter = 'A'

		printTb(x-2, y-1, p.labelsColor, p.labelsBg, "  ")
		for i := 0; i <= p.getWidth(); i++ {

			if mod4(i) == 0 {
				termbox.SetCell(x+i, y-1, letter, p.labelsColor, p.labelsBg)
				letter++
			} else {
				termbox.SetCell(x+i, y-1, ' ', p.labelsColor, p.labelsBg)
			}

		}
	}

	termbox.SetCell(x, y, left, p.boardColor, p.boardBg)

	for i := 1; i < p.getWidth(); i++ {

		xOffset := x + i

		if mod4(i) == 0 {
			termbox.SetCell(xOffset, y, middle, p.boardColor, p.boardBg)
			if labels {
				drawVertLine(p.boardColor, p.boardBg, xOffset, y+1, p.getHeight(), '│')
			}

		} else {
			termbox.SetCell(xOffset, y, '─', p.boardColor, p.boardBg)
			if labels {
				drawVertLine(p.boardColor, p.boardBg, xOffset, y+1, p.getHeight(), ' ')
			}
		}
	}

	termbox.SetCell(x+p.getWidth(), y, right, p.boardColor, p.boardBg)
}

// drawLeftAndRight draws left and right parts of a board
func drawLeftAndRight(x, y int, p *BoardDescription, middle rune, labels bool) {

	if labels {

		var index = 1

		for i := 0; i <= p.getHeight(); i++ {

			if mod2(i) == 0 {
				printfTb(x-2, y+i, p.labelsColor, p.labelsBg, "%2d", index)
				index++
			} else {
				printTb(x-2, y+i, p.labelsColor, p.labelsBg, "  ")
			}
		}
	}

	for i := 1; i < p.getHeight(); i++ {

		yOffset := y + i

		if mod2(i) == 0 {
			termbox.SetCell(x, yOffset, middle, p.boardColor, p.boardBg)
			if labels {
				drawHorizLine(p.boardColor, p.boardBg, x+1, yOffset, p.getWidth())
			}
		} else {
			termbox.SetCell(x, yOffset, '│', p.boardColor, p.boardBg)
		}

	}

}

// drawBoard draws ASCII game board
func drawBoard(boardParams *BoardDescription) {

	x := boardParams.x + 2
	y := boardParams.y + 2

	// draw top and bottom parts
	drawTopAndBottom(x, y, boardParams, '┌', '┬', '┐', true)
	drawTopAndBottom(x, y+boardParams.getHeight(), boardParams, '└', '┴', '┘', false)

	// draw left and right parts
	drawLeftAndRight(x, y, boardParams, '├', true)
	drawLeftAndRight(x+boardParams.getWidth(), y, boardParams, '┤', false)
}
