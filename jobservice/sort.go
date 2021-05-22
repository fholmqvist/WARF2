package jobservice

import "math/rand"

func (jb *JobService) Len() int {
	return len(jb.Jobs)
}

func (jb *JobService) Less(i, j int) bool {
	fst := jb.Jobs[i].Priority()
	snd := jb.Jobs[j].Priority()
	// Randomize equally prioritized.
	if fst == snd {
		return rand.Intn(2) == 1
	}
	// Highest priority first.
	return fst > snd
}

func (jb *JobService) Swap(i, j int) {
	jb.Jobs[i], jb.Jobs[j] = jb.Jobs[j], jb.Jobs[i]
}
