package schedulers

import (
	"math"
	. "structures"
	"time"
)

func GenerateNaiveScheduling(tasks TaskList) (scheduling Scheduling, schedulingTime int) {
	var processors Processors
	start := time.Now()
	instanceSize := len(tasks)
	tpr := Ceil(instanceSize,4)
	for i := 0; i < 4; i++ {
		for j := i * tpr; j < (i + 1) * tpr && j < instanceSize; j++ {
			scheduling.Tasks[i] = append(scheduling.Tasks[i], j)
			processors[i].Append(tasks[j])
		}
	}
	scheduling.Penalty = processors.TotalPenalty()
	schedulingTime = int(time.Since(start).Microseconds())
	return
}

func Ceil(a, b int) int {
	return int(math.Ceil(float64(a)/float64(b)))
}