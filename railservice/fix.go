package rail

import (
	"math"
)

func (r *RailService) FixRails(min, max int) {
	for idx := min; idx <= max; idx++ {
		t := &r.Map.Rails[idx]
		if t.Sprite == 0 {
			continue
		}
		up := r.Map.OneRailUp(idx).Idx
		right := r.Map.OneRailRight(idx).Idx
		down := r.Map.OneRailDown(idx).Idx
		left := r.Map.OneRailLeft(idx).Idx
		// CROSS
		if IsRail(r.Map, up) && IsRail(r.Map, right) &&
			IsRail(r.Map, down) && IsRail(r.Map, left) {
			t.Sprite = Cross
			continue
		}
		// HORIZONTAL
		if IsRail(r.Map, left) && IsRail(r.Map, right) {
			t.Rotation = math.Pi * 1.5
			continue
		}
		// UP RIGHT
		if IsRail(r.Map, up) && IsRail(r.Map, right) {
			t.Sprite = Curve
			continue
		}
		// UP LEFT
		if IsRail(r.Map, up) && IsRail(r.Map, left) {
			t.Sprite = Curve
			t.Rotation = math.Pi * 1.5
			continue
		}
		// DOWN RIGHT
		if IsRail(r.Map, down) && IsRail(r.Map, right) {
			t.Sprite = Curve
			t.Rotation = -math.Pi * 1.5
			continue
		}
		// DOWN LEFT
		if IsRail(r.Map, down) && IsRail(r.Map, left) {
			t.Sprite = Curve
			t.Rotation = math.Pi * 3.0
			continue
		}
		// LONE UP
		if !IsRail(r.Map, up) && IsRail(r.Map, down) {
			t.Sprite = Stop
			continue
		}
		// LONE RIGHT
		if !IsRail(r.Map, right) && IsRail(r.Map, left) {
			t.Sprite = Stop
			t.Rotation = -math.Pi * 1.5
			continue
		}
		// LONE DOWN
		if !IsRail(r.Map, down) && IsRail(r.Map, up) {
			t.Sprite = Stop
			t.Rotation = math.Pi * 3.0
			continue
		}
		// LONE LEFT
		if !IsRail(r.Map, left) && IsRail(r.Map, right) {
			t.Sprite = Stop
			t.Rotation = math.Pi * 1.5
			continue
		}
	}
}
