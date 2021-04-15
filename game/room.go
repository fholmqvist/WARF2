package game

import "projects/games/warf2/room"

func (g *Game) AddLibrary(x, y int) {
	l := room.NewLibrary(&g.WorldMap, x, y)
	g.Rooms.Libraries = append(g.Rooms.Libraries, l)
}
