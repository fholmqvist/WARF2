package room

import (
	"projects/games/warf2/resource"
	m "projects/games/warf2/worldmap"
	"testing"
)

func TestNearestStorage(t *testing.T) {
	mp := m.BoundariesMap()
	service := Service{}
	_, _, ok := service.FindNearestStorage(mp, 1, 1)
	if ok {
		t.Fatal("did not expect to be ok")
	}
	mp.DrawOutline(5, 5, 10, 10, m.WallSolid)
	mp.DrawOutline(20, 5, 25, 10, m.WallSolid)
	s1 := NewStorage(mp, 6, 6)
	s2 := NewStorage(mp, 21, 6)
	service.Storages = append(service.Storages, *s1)
	service.Storages = append(service.Storages, *s2)
	ns, _, ok := service.FindNearestStorage(mp, 1, 1)
	if !ok {
		t.Fatal("expected to be ok, wasn't")
	}
	if ns.Center != s1.Center {
		t.Fatalf("\nexpected ns to be s1\nns: %v\ns1: %v\ns2: %v", ns.Center, s1.Center, s2.Center)
	}
}

func TestStorageTileAdd(t *testing.T) {
	st := StorageTile{
		Idx:    0,
		Tpe:    resource.Rock,
		Amount: 0,
	}
	r := st.Add(resource.Rock, 5)
	if st.Amount != 5 || r != 0 {
		t.Fatalf("wanted [%v, %v] got [%v %v]", 5, 0, st.Amount, r)
	}
	r = st.Add(resource.Rock, 5)
	if st.Amount != 8 || r != 2 {
		t.Fatalf("wanted [%v, %v] got [%v %v]", 8, 2, st.Amount, r)
	}
}

func TestStorageTileTake(t *testing.T) {
	st := StorageTile{
		Idx:    0,
		Tpe:    resource.Rock,
		Amount: 10,
	}
	r := st.Take(5)
	if st.Amount != 5 || r != 5 {
		t.Fatalf("wanted [%v, %v] got [%v %v]", 5, 5, st.Amount, r)
	}
	r = st.Take(10)
	if st.Amount != 0 || r != 5 {
		t.Fatalf("wanted [%v, %v] got [%v %v]", 0, 5, st.Amount, r)
	}
}
