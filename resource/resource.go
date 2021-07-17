package resource

import gl "github.com/Holmqvist1990/WARF2/globals"

type Resource int

const (
	None Resource = iota
	Rock
	Wheat
)

func (r Resource) String() string {
	return []string{"None", "Rock", "Wheat"}[r]
}

func SpriteToResource(sprite int) Resource {
	if sprite == gl.Wheat {
		return Wheat
	}
	if gl.IsCrumbledWall(sprite) {
		return Rock
	}
	return None
}
