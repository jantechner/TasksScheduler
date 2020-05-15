package verifier

import (
	. "inout"
	. "schedulers"
	. "structures"
)

func CalculatePenalty(instanceFilename, schedulingFilename string) (originalPenalty, calculatedPenalty int) {
	tasks := ReadTasks(instanceFilename)
	originalPenalty, processors := ReadScheduling(schedulingFilename, tasks)
	calculatedPenalty = processors.TotalPenalty()
	return
}

func CalculatePenaltyForInstance(instanceFilename string, schedule SchedulerFunction) (Scheduling, int) {
	tasks := ReadTasks(instanceFilename)
	scheduling, schedulingTime := schedule(tasks)
	return scheduling, schedulingTime
}

func CalculatePenaltyForScheduling(scheduling Scheduling, tasks []Task) int {
	penalty := 0
	for j := 0; j < 4; j++ {
		time := 0
		for i := 0; i < len(scheduling.Tasks[j]); i++ {
			taskIndex := scheduling.Tasks[j][i] - 1
			time = Max(time, tasks[taskIndex].ReadyTime) + tasks[taskIndex].Duration
			if tasks[taskIndex].Deadline < time {
				penalty += time - tasks[taskIndex].Deadline
			}
		}
	}
	return penalty
}