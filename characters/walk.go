package characters

import (
	e "projects/games/warf2/entity"
	m "projects/games/warf2/worldmap"

	"github.com/beefsack/go-astar"
)

// Walker defines functionality for walking
type Walker struct {
	path []int
}

// Move attempts to move an entity given a direction
func (w *Walker) Move(mp *m.Map, e *e.Entity, d m.Direction) bool {
	switch d {
	case m.Up:
		return w.moveUp(mp, e)
	case m.Right:
		return w.moveRight(mp, e)
	case m.Down:
		return w.moveDown(mp, e)
	case m.Left:
		return w.moveLeft(mp, e)
	}
	return false
}

func (w *Walker) moveUp(mp *m.Map, e *e.Entity) bool {
	if e.Idx > m.TilesW && m.NotColliding(mp, e.Idx, m.Up) {
		e.Idx = m.OneTileUp(e.Idx)
		return true
	}
	return false
}

func (w *Walker) moveRight(mp *m.Map, e *e.Entity) bool {
	if e.Idx%m.TilesW-(m.TilesW-1) != 0 && m.NotColliding(mp, e.Idx, m.Right) {
		e.Idx = m.OneTileRight(e.Idx)
		return true
	}
	return false
}

func (w *Walker) moveDown(mp *m.Map, e *e.Entity) bool {
	if e.Idx < m.TilesT-m.TilesW && m.NotColliding(mp, e.Idx, m.Down) {
		e.Idx = m.OneTileDown(e.Idx)
		return true
	}
	return false
}

func (w *Walker) moveLeft(mp *m.Map, e *e.Entity) bool {
	if e.Idx%m.TilesW != 0 && m.NotColliding(mp, e.Idx, m.Left) {
		e.Idx = m.OneTileLeft(e.Idx)
		return true
	}
	return false
}

// InitiateWalk sets a new path
// and proceeds to start walking it.
func (w *Walker) InitiateWalk(path []astar.Pather) {
	var pathIdxs []int
	for _, t := range m.Reverse(path) {
		tile := t.(*m.Tile)
		pathIdxs = append(pathIdxs, tile.Idx)
	}

	w.path = pathIdxs
}
