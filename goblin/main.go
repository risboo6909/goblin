package main

import (
	"os"

	"github.com/nsf/termbox-go"
	"github.com/risboo6909/goblin/goblin-misc"
	"github.com/risboo6909/goblin/goblin-ui"
)

var gameState = StateGameplay

var board = misc.NewBoard(10, 10, 10, 10, termbox.ColorBlack, termbox.ColorBlue,
	termbox.ColorRed, termbox.ColorBlack)

var cursor = &misc.Cursor{Col: 2, Row: 1,
	FgColor: termbox.ColorGreen, BgColor: termbox.ColorWhite}

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
				cursor.Col++
			}
			if ev.Key == termbox.KeyArrowLeft {
				cursor.Col--
			}
			if ev.Key == termbox.KeyArrowUp {
				cursor.Row--
			}
			if ev.Key == termbox.KeyArrowDown {
				cursor.Row++
			}
			if ev.Key == termbox.KeySpace || ev.Key == termbox.KeyEnter {
				board.SetCell(cursor.Col, cursor.Row, misc.X)
			}
		}
	}
}

func paint() {

	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)

	switch gameState {

	case StateGameplay:
		ui.DrawBoard(board, cursor)
	}

	termbox.Flush()
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
