package game

import (
	"math/rand"
	"time"

	"github.com/Holmqvist1990/WARF2/globals"
)

const (
	NORMAL = 2
	FAST   = 1
	SUPER  = 0
)

var FramesToMove = NORMAL

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Time struct {
	Frame        int
	stop         bool
	framesToMove int
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
	t.framesToMove--
	if t.framesToMove <= -1 {
		t.framesToMove = FramesToMove
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

// Stops time from incrementing.
func (t *Time) Stop() {
	t.stop = !t.stop
}
