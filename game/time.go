package game

import (
	"math/rand"
	"projects/games/warf2/globals"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Time holds all functionality relating to time and timing
type Time struct {
	Frame int
}

// Tick decriments time by one
func (t *Time) Tick() {
	t.Frame--

	if t.Frame <= -1 {
		t.Frame = globals.CycleLength
	}
}

// TimeToMove returns whether its time for characters to move
func (t *Time) TimeToMove() bool {
	return t.Frame%3 == 0
}

// NewCycle returns a bool informing
// whether a complete cycle has been finished
func (t *Time) NewCycle() bool {
	return t.Frame == 0
}
