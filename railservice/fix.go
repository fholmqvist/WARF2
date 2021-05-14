package rail

import (
	"math"
	m "projects/games/warf2/worldmap"
)

func (r *RailService) FixRails(min, max int) {
	for idx := min; idx <= max; idx++ {
		t := &r.Map.Rails[idx]
		if t.Sprite == 0 {
			continue
		}
		up := r.Map.OneRailUp(idx).Sprite
		right := r.Map.OneRailRight(idx).Sprite
		down := r.Map.OneRailDown(idx).Sprite
		left := r.Map.OneRailLeft(idx).Sprite
		// CROSS
		if m.IsRail(up) && m.IsRail(right) &&
			m.IsRail(down) && m.IsRail(left) {
			t.Sprite = m.Cross
			continue
		}
		// HORIZONTAL
		if m.IsRail(left) && m.IsRail(right) {
			t.Rotation = math.Pi * 1.5
			continue
		}
		// UP RIGHT
		if m.IsRail(up) && m.IsRail(right) {
			t.Sprite = m.Curve
			continue
		}
		// UP LEFT
		if m.IsRail(up) && m.IsRail(left) {
			t.Sprite = m.Curve
			t.Rotation = math.Pi * 1.5
			continue
		}
		// DOWN RIGHT
		if m.IsRail(down) && m.IsRail(right) {
			t.Sprite = m.Curve
			t.Rotation = -math.Pi * 1.5
			continue
		}
		// DOWN LEFT
		if m.IsRail(down) && m.IsRail(left) {
			t.Sprite = m.Curve
			t.Rotation = math.Pi * 3.0
			continue
		}
		// LONE UP
		if !m.IsRail(up) && m.IsRail(down) {
			t.Sprite = m.Stop
			continue
		}
		// LONE RIGHT
		if !m.IsRail(right) && m.IsRail(left) {
			t.Sprite = m.Stop
			t.Rotation = -math.Pi * 1.5
			continue
		}
		// LONE DOWN
		if !m.IsRail(down) && m.IsRail(up) {
			t.Sprite = m.Stop
			t.Rotation = math.Pi * 3.0
			continue
		}
		// LONE LEFT
		if !m.IsRail(left) && m.IsRail(right) {
			t.Sprite = m.Stop
			t.Rotation = math.Pi * 1.5
			continue
		}
	}
}
