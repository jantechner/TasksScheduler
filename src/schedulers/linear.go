package schedulers

import (
	. "structures"
	"time"
)

type SchedulerFunction func (TaskList) (Scheduling, int)

func GenerateLinearScheduling(tasks TaskList) (scheduling Scheduling, schedulingTime int) {
	var processors Processors
	start := time.Now()
	readyTime := func(t1, t2 *Task) bool { return t1.ReadyTime < t2.ReadyTime }
	deadline := func(t1, t2 *Task) bool { return t1.Deadline < t2.Deadline }
	//name := func(t1, t2 *Task) bool { return t1.Number < t2.Number }

	OrderedBy(deadline, readyTime).Sort(tasks)

	var minProc int
	for _, task := range tasks {
		minProc = 0
		for i := 0; i < 4; i++ {
			if processors[i].Time < processors[minProc].Time {
				minProc = i
			}
		}
		processors[minProc].Append(task)
	}
	scheduling = processors.GenerateScheduling()

	schedulingTime = int(time.Since(start).Microseconds())
	return
}
