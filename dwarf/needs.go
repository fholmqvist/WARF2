package dwarf

const MAX_NEED = uint16(100)

type Needs struct {
	ToRead uint16
}

func (n *Needs) Update(c Characteristics) {
	n.DesireToRead(c)
}

func (n *Needs) DesireToRead(c Characteristics) {
	if n.ToRead > MAX_NEED {
		n.ToRead = MAX_NEED
		return
	}
	n.ToRead += c.DesireToRead
}
