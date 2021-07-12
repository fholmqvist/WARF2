package room

import (
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

func (s *Service) DeleteRoomAtMousePos(mp *m.Map, currentMousePos int) {
	for i, l := range s.Libraries {
		for _, lt := range l.tiles {
			if lt.Idx == currentMousePos {
				s.DeleteLibrary(mp, i)
				return
			}
		}
	}
	for i, st := range s.Storages {
		for _, t := range st.StorageTiles {
			if t.Idx == currentMousePos {
				s.DeleteStorage(mp, i)
				return
			}
		}
	}
}

func (s *Service) DeleteLibrary(mp *m.Map, idx int) {
	l := s.Libraries[idx]
	for _, t := range l.tiles {
		mp.Tiles[t.Idx].Sprite = m.Ground
		mp.Items[t.Idx].Sprite = 0
	}
	s.Libraries = append(s.Libraries[:idx], s.Libraries[idx+1:]...)
}

func (s *Service) DeleteStorage(mp *m.Map, idx int) {
	st := s.Storages[idx]
	for _, t := range st.StorageTiles {
		mp.Tiles[t.Idx].Sprite = m.Ground
		mp.Items[t.Idx].Sprite = 0
	}
	s.Storages = append(s.Storages[:idx], s.Storages[idx+1:]...)
}
