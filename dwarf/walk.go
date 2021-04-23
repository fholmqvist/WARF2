package dwarf

import (
	e "projects/games/warf2/entity"
	m "projects/games/warf2/worldmap"

	"github.com/beefsack/go-astar"
)

// Walker defines functionality for walking
type Walker struct {
	Path []int
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

// InitiateWalk attempts to set
// a new path based on the given
// destinations and proceeds to start
// walking it if it was successful.
// Return value determins whether
// path was found.
func (w *Walker) InitiateWalk(from, to *m.Tile) bool {
	path, _, ok := astar.Path(from, to)
	if !ok {
		return false
	}

	var pathIdxs []int
	for _, t := range m.Reverse(path) {
		tile := t.(*m.Tile)
		pathIdxs = append(pathIdxs, tile.Idx)
	}

	w.Path = pathIdxs
	return true
}
