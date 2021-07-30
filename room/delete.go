package room

import (
	"fmt"

	"github.com/Holmqvist1990/WARF2/entity"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

func (s *Service) DeleteRoomAtMousePos(mp *m.Map, currentMousePos int) {
	pointer := mp.Tiles[currentMousePos].Room
	if pointer == nil {
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
	case *SleepHall:
		s.DeleteSleepHall(mp, id)
	case *Farm:
		s.DeleteFarm(mp, id)
	case *Library:
		s.DeleteLibrary(mp, id)
	default:
		panic(fmt.Sprintf("unknown room type: %v", rm.String()))
	}
}

func (s *Service) DeleteLibrary(mp *m.Map, id int) {
	l, idx := getLibrary(s, id)
	for _, t := range l.tiles {
		ResetGroundTile(mp, t.Idx)
	}
	s.Libraries = append(s.Libraries[:idx], s.Libraries[idx+1:]...)
}

func (s *Service) DeleteStorage(mp *m.Map, id int) {
	st, idx := getStorage(s, id)
	for _, t := range st.StorageTiles {
		ResetGroundTile(mp, t.Idx)
	}
	s.Storages = append(s.Storages[:idx], s.Storages[idx+1:]...)
}

func (s *Service) DeleteFarm(mp *m.Map, id int) {
	f, idx := getFarm(s, id)
	for _, idx := range f.AllTileIdxs {
		ResetGroundTile(mp, idx)
	}
	s.Farms = append(s.Farms[:idx], s.Farms[idx+1:]...)
}

func (s *Service) DeleteSleepHall(mp *m.Map, id int) {
	h, idx := getSleepHall(s, id)
	for _, t := range h.tiles {
		ResetGroundTile(mp, t.Idx)
	}
	s.SleepHalls = append(s.SleepHalls[:idx], s.SleepHalls[idx+1:]...)
}

func ResetGroundTile(mp *m.Map, idx int) {
	mp.Tiles[idx].Sprite = m.Ground
	mp.Tiles[idx].Room = nil
	if entity.IsCarriable(mp.Items[idx].Sprite) {
		return
	}
	mp.Items[idx].Sprite = m.None
}

func getLibrary(s *Service, id int) (Library, int) {
	var lib Library
	var idx int
	for i, l := range s.Libraries {
		if l.ID == id {
			lib = l
			idx = i
			break
		}
	}
	return lib, idx
}

func getStorage(s *Service, id int) (Storage, int) {
	var st Storage
	var idx int
	for i, storage := range s.Storages {
		if storage.ID == id {
			st = storage
			idx = i
			break
		}
	}
	return st, idx
}

func getFarm(s *Service, id int) (Farm, int) {
	var f Farm
	var idx int
	for i, farm := range s.Farms {
		if farm.ID == id {
			f = farm
			idx = i
			break
		}
	}
	return f, idx
}

func getSleepHall(s *Service, id int) (SleepHall, int) {
	var h SleepHall
	var idx int
	for i, hall := range s.SleepHalls {
		if hall.ID == id {
			h = hall
			idx = i
			break
		}
	}
	return h, idx
}
