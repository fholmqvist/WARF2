package item

var blockingItems = append([]int{Table}, bookShelves...)

// IsBlocking returns true
// if the sprite is of a
// type that blocks.
func IsBlocking(sprite int) bool {
	for _, i := range blockingItems {
		if i == sprite {
			return true
		}
	}
	return false
}
