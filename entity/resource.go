package entity

type Resource int

const (
	ResourceNone Resource = iota
	ResourceRock
	ResourceWheat
	ResourceBeer
)

func (r Resource) String() string {
	return []string{"None", "Rock", "Wheat", "Beer"}[r]
}

func SpriteToResource(sprite int) Resource {
	if sprite == Wheat {
		return ResourceWheat
	}
	if IsCrumbledWall(sprite) {
		return ResourceRock
	}
	if sprite == FilledBarrel {
		return ResourceBeer
	}
	return ResourceNone
}
