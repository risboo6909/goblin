package misc

import (
	"time"
	"math/rand"
)


// Sessions manager

type Session struct {
	Board      *BoardDescription
	AI         AIOptions
	SessionID  string
	Winner     Cell
	Intervals  []Interval
}


const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// generate random ID n symbols length, taken from
// http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func generateSessionId(n int) string {

	var src = rand.NewSource(time.Now().UnixNano())

	// session id will be just a time it was created
	b := make([]byte, n)

	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func CreateNewSession(boardSide int, player Cell) Session {

	rand.Seed(time.Now().UTC().UnixNano())

	return Session{
		NewBoard(boardSide, boardSide),

		AIOptions{switchPlayer(player),
			  5,
		          4},

		generateSessionId(10),
		E,
		[]Interval{},
	}
}

func RemoveSession() {

}


func (s *Session) MakeMove() {
	winner, intervals := MakeMove(s.Board, s.AI)
	s.Winner = winner
	s.Intervals = intervals
}