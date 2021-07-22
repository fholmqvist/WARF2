package room

import (
	"github.com/Holmqvist1990/WARF2/globals"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

func (s *Service) DeleteRoomAtMousePos(mp *m.Map, currentMousePos int) {
	/////////////////////////////////////////////
	// TODO
	// Are you beginning to smell an abstraction?
	/////////////////////////////////////////////
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
	for i, f := range s.Farms {
		for _, idx := range f.AllTileIdxs {
			if idx == currentMousePos {
				s.DeleteFarm(mp, i)
				return
			}
		}
	}
	for i, sh := range s.SleepHalls {
		for _, t := range sh.tiles {
			if t.Idx == currentMousePos {
				s.DeleteSleepHall(mp, i)
				return
			}
		}
	}
}

func (s *Service) DeleteLibrary(mp *m.Map, idx int) {
	l := s.Libraries[idx]
	for _, t := range l.tiles {
		mp.Tiles[t.Idx].Sprite = m.Ground
		mp.Items[t.Idx].Sprite = m.None
	}
	s.Libraries = append(s.Libraries[:idx], s.Libraries[idx+1:]...)
}

func (s *Service) DeleteStorage(mp *m.Map, idx int) {
	st := s.Storages[idx]
	for _, t := range st.StorageTiles {
		mp.Tiles[t.Idx].Sprite = m.Ground
		mp.Items[t.Idx].Sprite = m.None
	}
	s.Storages = append(s.Storages[:idx], s.Storages[idx+1:]...)
}

func (s *Service) DeleteFarm(mp *m.Map, idx int) {
	f := s.Farms[idx]
	for _, idx := range f.AllTileIdxs {
		mp.Tiles[idx].Sprite = m.Ground
		if mp.Items[idx].Sprite == globals.Wheat {
			continue
		}
		mp.Items[idx].Sprite = m.None
	}
	s.Farms = append(s.Farms[:idx], s.Farms[idx+1:]...)
}

func (s *Service) DeleteSleepHall(mp *m.Map, idx int) {
	for _, t := range s.SleepHalls[idx].tiles {
		mp.Tiles[t.Idx].Sprite = m.Ground
		mp.Items[t.Idx].Sprite = m.None
	}
	s.SleepHalls = append(s.SleepHalls[:idx], s.SleepHalls[idx+1:]...)
}
