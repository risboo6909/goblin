package ui

import "github.com/nsf/termbox-go"

// Cursor represents board cursor position
type Cursor struct {
	Board *DrawableBoard

	Col, Row         int
	FgColor, BgColor termbox.Attribute
}


func (c *Cursor) MoveRight() {
	if c.Col < c.Board.CellsHoriz - 1 {
		c.Col++
	}
}

func (c *Cursor) MoveLeft() {
	if c.Col > 0 {
		c.Col--
	}
}

func (c *Cursor) MoveUp() {
	if c.Row > 0 {
		c.Row--
	}
}

func (c *Cursor) MoveDown() {
	if c.Row < c.Board.CellsVert - 1 {
		c.Row++
	}
}
