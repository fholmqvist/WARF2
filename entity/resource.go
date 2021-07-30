package entity

type Resource int

const (
	ResourceNone Resource = iota
	ResourceRock
	ResourceWheat
)

func (r Resource) String() string {
	return []string{"None", "Rock", "Wheat"}[r]
}

func SpriteToResource(sprite int) Resource {
	if sprite == Wheat {
		return ResourceWheat
	}
	if IsCrumbledWall(sprite) {
		return ResourceRock
	}
	return ResourceNone
}
