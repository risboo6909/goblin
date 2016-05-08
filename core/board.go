package main

import (
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

// BoardParams defines main board properties
type BoardParams struct {
    cellsHoriz int
    cellsVert int
    x, y int
    boardColor, boardBg termbox.Attribute
    labelsColor, labelsBg termbox.Attribute
}

func (p *BoardParams) getWidth() int {
    return (p.cellsHoriz - 1) * 4
}

func (p *BoardParams) getHeight() int {
    return (p.cellsVert - 1) * 2
}

// drawHorizLine draws horizontal board lines
func drawHorizLine(color, bgcolor termbox.Attribute, x, y, width int, s rune) {
    for i := 0; i < width; i++ {
        if mod4(i) != 3 {
            termbox.SetCell(x + i, y, s, color, bgcolor)
        } else {
            termbox.SetCell(x + i, y, '┼', color, bgcolor)
        }
    }
}

// drawVertLine draw vertical board lines
func drawVertLine(color, bgcolor termbox.Attribute, x, y, height int, s rune) {
    for i := 0; i < height; i++ {
        termbox.SetCell(x, y + i, s, color, bgcolor)
    }
}

// drawTopAndBottom draws upper or lower parts of a board
func drawTopAndBottom(x, y int, p BoardParams, left, middle, right rune, labels bool) {
    if labels {

        var letter = 'A'

        printTb(x - 2, y - 1, p.labelsColor, p.labelsBg, "  ")
        for i := 0; i <= p.getWidth(); i++ {
            if mod4(i) == 0 {
                termbox.SetCell(x + i, y - 1, letter, p.labelsColor, p.labelsBg)
                letter++
            } else {
                termbox.SetCell(x + i, y - 1, ' ', p.labelsColor, p.labelsBg)
            }
        }
    }

    termbox.SetCell(x, y, left, p.boardColor, p.boardBg)

    for i := 1; i < p.getWidth(); i++ {
        xOffset := x + i
        if mod4(i) == 0 {
            termbox.SetCell(xOffset, y, middle, p.boardColor, p.boardBg)
            if labels {
                drawVertLine(p.boardColor, p.boardBg, xOffset, y + 1, p.getHeight(), '│')
            }
        } else {
            termbox.SetCell(xOffset, y, '─', p.boardColor, p.boardBg)
            if labels {
                drawVertLine(p.boardColor, p.boardBg, xOffset, y + 1, p.getHeight(), ' ')
            }
        }
    }

    termbox.SetCell(x + p.getWidth(), y, right, p.boardColor, p.boardBg)
}

// drawLeftAndRight draws left and right parts of a board
func drawLeftAndRight(x, y int, p BoardParams, middle rune, labels bool) {

    if labels {
        var index = 1
        for i := 0; i <= p.getHeight(); i++ {
            if mod2(i) == 0 {
                printfTb(x - 2, y + i, p.labelsColor, p.labelsBg, "%2d", index)
                index++
            } else {
                printTb(x - 2, y + i, p.labelsColor, p.labelsBg, "  ")
            }
        }
    }

    for i := 1; i < p.getHeight(); i++ {
        yOffset := y + i
        if mod2(i) == 0 {
            termbox.SetCell(x, yOffset, middle, p.boardColor, p.boardBg)
            if labels {
                drawHorizLine(p.boardColor, p.boardBg, x + 1, yOffset, p.getWidth(), '─')
            }
        } else {
            termbox.SetCell(x, yOffset, '│', p.boardColor, p.boardBg)
        }
    }

}

// drawBoard draws ASCII game board
func drawBoard(boardParams BoardParams) {

    x := boardParams.x + 2; y := boardParams.y + 2;

    // draw top and bottom parts
    drawTopAndBottom(x, y, boardParams, '┌', '┬', '┐', true)
    drawTopAndBottom(x, y + boardParams.getHeight(), boardParams, '└', '┴', '┘', false)

    // draw left and right parts
    drawLeftAndRight(x, y, boardParams, '├', true)
    drawLeftAndRight(x + boardParams.getWidth(), y, boardParams, '┤', false)
}
