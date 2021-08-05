// Package jobservice manages
// all the life cycles for jobs
// given to in-game characters.
package jobservice

import (
	"math/rand"
	"time"

	"github.com/Holmqvist1990/WARF2/dwarf"
	"github.com/Holmqvist1990/WARF2/job"
	"github.com/Holmqvist1990/WARF2/room"
	m "github.com/Holmqvist1990/WARF2/worldmap"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Service manages all
// ingame jobs for dwarves.
type Service struct {
	Jobs             []job.Job      `json:"jobs"`
	Workers          []*dwarf.Dwarf `json:"workers"`
	AvailableWorkers []*dwarf.Dwarf `json:"-"`
	Map              *m.Map         `json:"-"`
}

func NewService(mp *m.Map) *Service {
	return &Service{Map: mp}
}

// Update runs every frame, handling
// the lifetime cycle of jobs.
// Most likely to be the biggest
// performance hit. If game slows
// down, start here.
func (s *Service) Update(rs *room.Service, mp *m.Map) {
	if len(s.Jobs) > 0 {
		s.removeFinishedJobs(rs)
		s.updateAvailableWorkers()
	}
	if len(s.AvailableWorkers) > 0 {
		///////////////////////////////
		// TODO
		// Keep track of failed checks
		// with backoff to reduce CPU
		// load for complex checks that
		// now fail on every frame.
		///////////////////////////////
		s.checkForJobs(rs)
		s.checkForNeeds(mp, rs)
	}
	if len(s.Jobs) > 0 {
		s.sortJobPriorities()
		s.assignWorkers()
		s.performWork(rs)
	}
}

func (s *Service) removeFinishedJobs(rs *room.Service) {
	var jobs []job.Job
	for _, job := range s.Jobs {
		if job.Remove() {
			job.Finish(s.Map, rs)
			continue
		}
		jobs = append(jobs, job)
	}
	s.Jobs = jobs
}

func (s *Service) assignWorkers() {
	for _, job := range s.Jobs {
		if HasWorker(job) {
			continue
		}
	lookingForWorker:
		for _, worker := range s.AvailableWorkers {
			if worker.HasJob() {
				continue
			}
			if !SetWorkerAndMove(job, worker, s.Map) {
				continue
			}
			break lookingForWorker
		}
	}
}

func (s *Service) updateAvailableWorkers() {
	var available []*dwarf.Dwarf
	for _, dwarf := range s.Workers {
		if dwarf.Available() {
			available = append(available, dwarf)
		}
	}
	s.AvailableWorkers = available
}

func (s *Service) performWork(rs *room.Service) {
	for _, jb := range s.Jobs {
		dw := jb.GetWorker()
		if dw == nil {
			continue
		}
		if !dw.HasJob() {
			continue
		}
		for _, destination := range jb.GetDestinations() {
			if dw.Idx == destination {
				dw.State = dwarf.Arrived
				break
			}
		}
		if dw.State != dwarf.Arrived {
			if len(dw.Path) == 0 && jb.Remove() {
				dw.SetToAvailable()
				continue
			}
			if !jb.HasInternalMove() {
				continue
			}
		}
		finished := jb.PerformWork(s.Map, s.Workers, rs)
		if !finished {
			continue
		}
		dw.SetToAvailable()
	}
}
