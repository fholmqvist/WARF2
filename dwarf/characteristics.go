package dwarf

import (
	"math/rand"
)

type Attributes struct {
	Name         string
	DesireToRead uint16
}

func GenerateAttributes(name string) Attributes {
	return Attributes{
		Name:         name,
		DesireToRead: uint16(1 + rand.Intn(15)),
	}
}
