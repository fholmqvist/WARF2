package mouse

// System for handling
// all functionality by mouse.
type System struct {
	Mode Mode
}

// Mode enum for managing mouse action state.
type Mode int

// Mode enum.
const (
	Normal Mode = iota
	FloorTiles
	ResetFloor
)
