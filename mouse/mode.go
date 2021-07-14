package mouse

// Mode enum for managing mouse action state.
type Mode int

// Mode enum.
const (
	Normal Mode = iota
	Storage
	Farm
	Library
	Delete
)

var ModeFromString = map[string]Mode{
	"Normal":  Normal,
	"Storage": Storage,
	"Farm":    Farm,
	"Library": Library,
	"Delete":  Delete,
}

func (m Mode) String() string {
	return []string{
		"Normal",
		"Storage",
		"Farm",
		"Library",
		"Delete",
	}[m]
}
