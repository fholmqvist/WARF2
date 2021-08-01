package mouse

// Mode enum for managing mouse action state.
type Mode int

// Mode enum.
const (
	Normal Mode = iota
	Storage
	SleepHall
	Farm
	Brewery
	Library
	Delete
)

var ModeFromString = map[string]Mode{
	"Normal":    Normal,
	"Storage":   Storage,
	"SleepHall": SleepHall,
	"Farm":      Farm,
	"Brewery":   Brewery,
	"Library":   Library,
	"Delete":    Delete,
}

func (m Mode) String() string {
	return []string{
		"Normal",
		"Storage",
		"SleepHall",
		"Farm",
		"Brewery",
		"Library",
		"Delete",
	}[m]
}
