package misc

import (
	"time"
	"math/rand"
)

// Sessions manager

type Session struct {
	Board *BoardDescription
	AI AIOptions
}


func CreateNewSession(boardSide int, player Cell) Session {

	rand.Seed(time.Now().UTC().UnixNano())

	return Session{
		NewBoard(boardSide, boardSide),

		AIOptions{switchPlayer(player),
			  5,
		          4},
	}
}

func RemoveSession() {

}


func (s Session) MakeMove() {
	MakeMove(s.Board, s.AI)
}