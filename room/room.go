package room

import m "github.com/Holmqvist1990/WARF2/worldmap"

type Room interface {
	GetID() int // Needs to check for nil due to race conditions!
	String() string
	Update(*m.Map)
	Tiles() []int
}
