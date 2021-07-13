package mouse

// Mode enum for managing mouse action state.
type Mode int

// Mode enum.
const (
	Normal Mode = iota

	Library
	Storage

	Delete
)

var ModeFromString = map[string]Mode{"Normal": Normal, "Library": Library, "Storage": Storage, "Delete": Delete}

func (m Mode) String() string {
	return []string{"Normal", "Library", "Storage", "Delete"}[m]
}
