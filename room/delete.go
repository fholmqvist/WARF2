package room

import (
	"fmt"

	"github.com/Holmqvist1990/WARF2/entity"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

func (s *Service) DeleteRoomAtMousePos(mp *m.Map, currentMousePos int) {
	pointer := mp.Tiles[currentMousePos].Room
	if pointer == nil {
		pointer = mp.Items[currentMousePos].Room
	}
	if pointer == nil {
		fmt.Println("ROOM POINTER WAS NIL AT", currentMousePos)
		return
	}
	rm, ok := pointer.(Room)
	if !ok {
		panic(fmt.Sprintf("unknown room type: %v", rm))
	}
	id := rm.GetID()
	switch rm.(type) {
	case *Storage:
		s.DeleteStorage(mp, id)
	default:
		s.DeleteRoom(mp, id, rm.String())
	}
}

func (s *Service) DeleteRoom(mp *m.Map, id int, roomType string) {
	var room Room
	var idx int
	for i, rm := range s.Rooms {
		if rm.String() == roomType && rm.GetID() == id {
			room = rm
			idx = i
			break
		}
	}
	if room == nil {
		return
	}
	for _, t := range room.Tiles() {
		ResetGroundTile(mp, t)
	}
	s.Rooms = append(s.Rooms[:idx], s.Rooms[idx+1:]...)
}

func (s *Service) DeleteStorage(mp *m.Map, id int) {
	st, idx := getStorage(s, id)
	if st == nil {
		return
	}
	for _, t := range st.StorageTiles {
		ResetGroundTile(mp, t.Idx)
		mp.Items[t.Idx].Resource = t.Resource
		mp.Items[t.Idx].ResourceAmount = t.ResourceAmount
	}
	s.Rooms = append(s.Rooms[:idx], s.Rooms[idx+1:]...)
}

func ResetGroundTile(mp *m.Map, idx int) {
	mp.Tiles[idx].Sprite = m.Ground
	mp.Tiles[idx].Room = nil
	if entity.IsCarriable(mp.Items[idx].Sprite) {
		return
	}
	mp.Items[idx].Sprite = m.None
	mp.Items[idx].Resource = entity.ResourceNone
	mp.Items[idx].ResourceAmount = 0
}

func getStorage(s *Service, id int) (*Storage, int) {
	for i, rm := range s.Rooms {
		storage, ok := rm.(*Storage)
		if !ok {
			continue
		}
		if storage.GetID() == id {
			return storage, i
		}
	}
	return nil, -1
}
