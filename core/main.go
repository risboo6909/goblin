package main

import (
	//"github.com/risboo6909/goblin/ai"
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

var gameState = StateGameplay

var board = newBoard(10, 10, 10, 10, termbox.ColorBlack, termbox.ColorBlue,
	termbox.ColorRed, termbox.ColorBlack)
var cursor = &Cursor{col: 2, row: 1, bgColor: termbox.ColorGreen,
	fgColor: termbox.ColorWhite}

func update() {

	switch gameState {

	case StateGameplay:
		ev := termbox.PollEvent()
		switch ev.Type {

		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				termbox.Close()
				os.Exit(0)
			}

			// cursor control
			if ev.Key == termbox.KeyArrowRight {
				cursor.col++
			}
			if ev.Key == termbox.KeyArrowLeft {
				cursor.col--
			}
			if ev.Key == termbox.KeyArrowUp {
				cursor.row--
			}
			if ev.Key == termbox.KeyArrowDown {
				cursor.row++
			}
			if ev.Key == termbox.KeySpace || ev.Key == termbox.KeyEnter {
				board.setCell(cursor.col, cursor.row, X)
			}
		}
	}
}

func paint() {

	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

	switch gameState {

	case StateGameplay:
		drawBoard(board, cursor)
	}

	termbox.Flush()
}

func printTb(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func printfTb(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	printTb(x, y, fg, bg, s)
}

func main() {

	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	for {
		paint()
		update()
	}

}
