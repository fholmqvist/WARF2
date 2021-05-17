package resource

type Resource int

const (
	None Resource = iota
	Rock
)

func (r Resource) String() string {
	return []string{"None", "Rock"}[r]
}
