package room

import (
	"testing"

	"github.com/Holmqvist1990/WARF2/entity"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

func TestNearestStorage(t *testing.T) {
	mp := m.BoundariesMap()
	service := Service{}
	_, _, ok := service.FindNearestStorage(mp, 1, 1, entity.ResourceNone)
	if ok {
		t.Fatal("did not expect to be ok")
	}
	mp.DrawOutline(5, 5, 10, 10, m.WallSolid)
	mp.DrawOutline(20, 5, 25, 10, m.WallSolid)
	if s, ok := NewStorage(mp, 6, 6); ok {
		service.Rooms = append(service.Rooms, s)
	}
	if s, ok := NewStorage(mp, 21, 6); ok {
		service.Rooms = append(service.Rooms, s)
	}
	ns, _, ok := service.FindNearestStorage(mp, 1, 1, entity.ResourceNone)
	if !ok {
		t.Fatal("expected to be ok, wasn't")
	}
	if ns.Center != service.Rooms[0].(*Storage).Center {
		t.Fatalf("\nexpected ns to be s1\nns: %v\ns1: %v\ns2: %v",
			ns.Center, service.Rooms[0].(*Storage).Center, service.Rooms[1].(*Storage).Center)
	}
}

func TestStorageTileAdd(t *testing.T) {
	st := StorageTile{&m.Tile{
		Idx:            0,
		Resource:       entity.ResourceRock,
		ResourceAmount: 0,
	}}
	r := st.Add(entity.ResourceRock, 5)
	if st.ResourceAmount != 5 || r != 0 {
		t.Fatalf("wanted [%v, %v] got [%v %v]", 5, 0, st.ResourceAmount, r)
	}
	r = st.Add(entity.ResourceRock, 5)
	if st.ResourceAmount != 8 || r != 2 {
		t.Fatalf("wanted [%v, %v] got [%v %v]", 8, 2, st.ResourceAmount, r)
	}
}

func TestStorageTileTake(t *testing.T) {
	st := StorageTile{&m.Tile{
		Idx:            0,
		Resource:       entity.ResourceRock,
		ResourceAmount: 10,
	}}
	r := st.Take(5)
	if st.ResourceAmount != 5 || r != 5 {
		t.Fatalf("wanted [%v, %v] got [%v %v]", 5, 5, st.ResourceAmount, r)
	}
	r = st.Take(10)
	if st.ResourceAmount != 0 || r != 5 {
		t.Fatalf("wanted [%v, %v] got [%v %v]", 0, 5, st.ResourceAmount, r)
	}
}
