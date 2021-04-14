package game

import "projects/games/warf2/room"

func (g *Game) AddLibrary(x1, x2, y1, y2 int) {
	l := room.NewLibrary(&g.WorldMap, x1, x2, y1, y2)
	g.Rooms.Libraries = append(g.Rooms.Libraries, l)
}
