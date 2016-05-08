package main

import (
//    "github.com/risboo6909/goblin/ai"
    "github.com/nsf/termbox-go"
    "fmt"
)

var gameState = StateGameplay

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

            boardParams := BoardParams{x:10, y:10, cellsHoriz:19, cellsVert:19,
            boardColor: termbox.ColorCyan, boardBg: termbox.ColorBlack,
            labelsColor: termbox.ColorRed, labelsBg: termbox.ColorBlack}

            drawBoard(boardParams)

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

            ev := termbox.PollEvent();

            switch  ev.Type {
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
