package game

import (
	"math/rand"
	"projects/games/warf2/globals"
	"time"
)

var FramesToMove = 3

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Time struct {
	Frame int
	stop  bool
}

// Decriments until one cycle
// has been consumed, then resets.
func (t *Time) Tick() bool {
	if t.stop {
		return false
	}
	t.Frame--
	if t.Frame <= -1 {
		t.Frame = globals.CycleLength
	}
	return true
}

// Time to update all entities.
func (t *Time) TimeToMove() bool {
	return t.Frame%FramesToMove == 0
}

// One game cycle has passed.
func (t *Time) NewCycle() bool {
	return t.Frame == 0
}

// Stops time from incrementing.
func (t *Time) Stop() {
	t.stop = !t.stop
}
