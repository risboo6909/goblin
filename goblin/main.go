package main

import (
	"os"

	"github.com/nsf/termbox-go"
	"github.com/risboo6909/goblin/misc"
	"github.com/risboo6909/goblin/ui"
)

var (
	gameState = StateGameplay
	moveBoard = false

	gameSession = misc.CreateNewSession(15, misc.X)

	board = ui.CloneExistingBoard(gameSession.Board, 0, 0, termbox.ColorBlack, termbox.ColorBlue,
		termbox.ColorRed, termbox.ColorBlack)

	cursor = ui.Cursor{Board: board, Col: 2, Row: 1,
		FgColor: termbox.ColorGreen, BgColor: termbox.ColorWhite}
)

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

			if ev.Key == termbox.KeyF2 {
				moveBoard = !moveBoard
			}

			// cursor control

			if ev.Key == termbox.KeyArrowRight {
				if moveBoard {
					board.X++
				} else {
					cursor.MoveRight()
				}
			}

			if ev.Key == termbox.KeyArrowLeft {
				if moveBoard {
					board.X--
				} else {
					cursor.MoveLeft()
				}
			}

			if ev.Key == termbox.KeyArrowUp {
				if moveBoard {
					board.Y--
				} else {
					cursor.MoveUp()
				}
			}

			if ev.Key == termbox.KeyArrowDown {
				if moveBoard {
					board.Y++
				} else {
					cursor.MoveDown()
				}
			}

			if (ev.Key == termbox.KeySpace || ev.Key == termbox.KeyEnter) && !moveBoard {
				// can't modify occupied cell
				if board.GetCell(cursor.Col, cursor.Row) == misc.E {
					board.SetCell(cursor.Col, cursor.Row, misc.X)
				}
				gameSession.MakeMove()
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
