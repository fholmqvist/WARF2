package game

const maxTurns = 60

// Clock to manage internal game time.
type Clock struct {
	time int
}

// Increment adds to the current time
// until maxTurns has been reached,
// in which the clock starts over.
func (c *Clock) Increment() {
	c.time++

	if c.time >= maxTurns {
		c.time = 0
	}
}

// FreshTurn returns if the
// current time is a new cycle.
func (c *Clock) FreshTurn() bool {
	return c.time == 0
}

// Half that of FreshTurn.
func (c *Clock) HalfTurn() bool {
	return c.time%(maxTurns/2) == 0
}
