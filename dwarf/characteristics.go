package dwarf

import "math/rand"

type Characteristics struct {
	DesireToRead uint16
}

func GenerateCharacteristics() Characteristics {
	return Characteristics{
		DesireToRead: uint16(1 + rand.Intn(15)),
	}
}
