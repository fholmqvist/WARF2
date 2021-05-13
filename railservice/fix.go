package rail

import (
	"math"
)

const (
	up    = 0
	right = 1
	down  = 2
	left  = 3
)

func (r *RailService) FixRails(min, max int) {
	for idx := min; idx <= max; idx++ {
		if r.Map.Rails[idx].Sprite == 0 {
			continue
		}
		rails := []int{
			r.Map.OneRailUp(idx).Idx,
			r.Map.OneRailRight(idx).Idx,
			r.Map.OneRailDown(idx).Idx,
			r.Map.OneRailLeft(idx).Idx,
		}
		t := &r.Map.Rails[idx]
		// CROSS
		if IsRail(r.Map, rails[up]) && IsRail(r.Map, rails[right]) &&
			IsRail(r.Map, rails[down]) && IsRail(r.Map, rails[left]) {
			t.Sprite = Cross
			continue
		}
		// HORIZONTAL
		if IsRail(r.Map, rails[left]) && IsRail(r.Map, rails[right]) {
			t.Rotation = math.Pi * 1.5
			continue
		}
		// UP RIGHT
		if IsRail(r.Map, rails[up]) && IsRail(r.Map, rails[right]) {
			t.Sprite = Curve
			continue
		}
		// UP LEFT
		if IsRail(r.Map, rails[up]) && IsRail(r.Map, rails[left]) {
			t.Sprite = Curve
			t.Rotation = math.Pi * 1.5
			continue
		}
		// DOWN RIGHT
		if IsRail(r.Map, rails[down]) && IsRail(r.Map, rails[right]) {
			t.Sprite = Curve
			t.Rotation = -math.Pi * 1.5
			continue
		}
		// DOWN LEFT
		if IsRail(r.Map, rails[down]) && IsRail(r.Map, rails[left]) {
			t.Sprite = Curve
			t.Rotation = math.Pi * 3.0
			continue
		}
		// LONE UP
		if !IsRail(r.Map, rails[up]) && IsRail(r.Map, rails[down]) {
			t.Sprite = Stop
			continue
		}
		// LONE RIGHT
		if !IsRail(r.Map, rails[right]) && IsRail(r.Map, rails[left]) {
			t.Sprite = Stop
			t.Rotation = -math.Pi * 1.5
			continue
		}
		// LONE DOWN
		if !IsRail(r.Map, rails[down]) && IsRail(r.Map, rails[up]) {
			t.Sprite = Stop
			t.Rotation = math.Pi * 3.0
			continue
		}
		// LONE LEFT
		if !IsRail(r.Map, rails[left]) && IsRail(r.Map, rails[right]) {
			t.Sprite = Stop
			t.Rotation = math.Pi * 1.5
			continue
		}
	}
}
