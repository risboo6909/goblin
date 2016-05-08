package main

import (
	//    "github.com/risboo6909/goblin/ai"
	"fmt"

	"github.com/nsf/termbox-go"
)

var gameState = StateGameplay
var board = newBoard(15, 15, 10, 10, termbox.ColorCyan, termbox.ColorBlack,
	termbox.ColorRed, termbox.ColorBlack)

func update() {

	// switch gameState {
	//
	//     case StateGameplay:
	//         fmt.Println("Gameplay update")
	//
	// }

}

func paint() {

	termbox.Clear(termbox.ColorBlue, termbox.ColorBlue)

	switch gameState {

	case StateGameplay:

		drawBoard(board)

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

mainloop:

	for {
		ev := termbox.PollEvent()
		switch ev.Type {
		// handle ESC as exit
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break mainloop
			}
		}
		update()
		paint()
	}

}
