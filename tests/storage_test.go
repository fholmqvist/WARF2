package tests

import (
	"projects/games/warf2/room"
	m "projects/games/warf2/worldmap"
	"testing"
)

func TestNearestStorage(t *testing.T) {
	mp := m.BoundariesMap()
	service := room.Service{}
	_, ok := service.FindNearestStorage(mp, 1, 1)
	if ok {
		t.Fatal("did not expect to be ok")
	}
	mp.DrawOutline(5, 5, 10, 10, m.WallSolid)
	mp.DrawOutline(20, 5, 25, 10, m.WallSolid)
	s1 := room.NewStorage(mp, 6, 6)
	s2 := room.NewStorage(mp, 21, 6)
	service.Storages = append(service.Storages, *s1)
	service.Storages = append(service.Storages, *s2)
	ns, ok := service.FindNearestStorage(mp, 1, 1)
	if !ok {
		t.Fatal("expected to be ok, wasn't")
	}
	if ns.Center != s1.Center {
		t.Fatalf("\nexpected ns to be s1\nns: %v\ns1: %v\ns2: %v", ns.Center, s1.Center, s2.Center)
	}
}
