package helpers

const maxTurns = 60

type Clock struct {
	time int
}

func (c *Clock) Increment() {
	c.time++

	if c.time >= maxTurns {
		c.time = 0
	}
}

func (c *Clock) FreshTurn() bool {
	return c.time == 0
}
