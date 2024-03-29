package game

import (
	"math/rand"
	"time"

	gl "github.com/Holmqvist1990/WARF2/globals"
)

const (
	NORMAL = 2
	FAST   = 1
	SUPER  = 0
)

var GAME_SPEED = NORMAL

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Time struct {
	Frame        int
	framesToMove int
}

// Decriments until one cycle
// has been consumed, then resets.
func (t *Time) Tick() bool {
	if gl.GAME_PAUSED {
		return false
	}
	t.Frame--
	if t.Frame <= -1 {
		t.Frame = gl.CycleLength
	}
	t.framesToMove--
	if t.framesToMove <= -1 {
		t.framesToMove = GAME_SPEED
	}
	return true
}

// Time to update all entities.
func (t *Time) TimeToMove() bool {
	return t.framesToMove == 0
}

// One game cycle has passed.
func (t *Time) NewCycle() bool {
	return t.Frame == 0
}

// Half a game cycle has passed.
func (t *Time) HalfCycle() bool {
	return t.NewCycle() || t.Frame == gl.CycleLength/2
}

// A quarter of a game cycle has passed.
func (t *Time) QuarterCycle() bool {
	return t.HalfCycle() || t.Frame == gl.CycleLength/4 ||
		t.Frame == gl.CycleLength/4*2 || t.Frame == gl.CycleLength/4*3
}
