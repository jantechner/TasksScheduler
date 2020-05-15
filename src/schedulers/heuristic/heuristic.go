package schedulers

import (
	"math/rand"
	. "structures"
	"time"
)

func GenerateHeuristicScheduling(tasks TaskList) (scheduling Scheduling, schedulingTime int) {
	start := time.Now()

	rand.Seed(time.Now().UTC().UnixNano())
	population := NewPopulation(0.03,100, tasks)
	//population := NewPopulation(0.05,12, tasks)
	population.Generate()
	for generation := 0; generation < 10000000; generation++ {
		duration := time.Since(start)
		if duration.Milliseconds() > int64(99*len(tasks)) {
			//fmt.Println("Generation", generation)
			break
		}
		population.Mutate()
		population.RecalculateMinPenalty()
	}
	scheduling = population.BestResult()
	//fmt.Println(scheduling.Tasks)

	schedulingTime = int(time.Since(start).Microseconds())
	return
}
