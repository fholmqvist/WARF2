package game

type GameState int

const (
	MainMenu GameState = iota
	HelpMenu
	Gameplay
)
