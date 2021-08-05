package dwarf

const (
	MAX       = uint16(100)
	SLEEP_INC = uint16(25)
	DRINK_INC = uint16(25)
)

type Needs struct {
	Sleep uint16
	Drink uint16
	Read  uint16
}

func (n *Needs) Update(c Characteristics) {
	n.Sleep = inc(n.Sleep, SLEEP_INC)
	n.Drink = inc(n.Drink, DRINK_INC)
	n.Read = inc(n.Read, c.DesireToRead)
}

func inc(val, amount uint16) uint16 {
	if val+amount > MAX {
		return MAX
	}
	return val + amount
}
