package dwarf

import (
	"math/rand"
)

type Characteristics struct {
	Name         string
	DesireToRead uint16
}

func GenerateCharacteristics(name string) Characteristics {
	return Characteristics{
		Name:         name,
		DesireToRead: uint16(1 + rand.Intn(15)),
	}
}
