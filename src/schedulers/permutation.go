package schedulers

import (
	. "structures"
)

func GeneratePermutationScheduling(inputTasks []Task, sortingMode int) Scheduling {
	tasks := append([]Task{}, inputTasks...)

	scheduling := NewScheduling(len(tasks))

	tasks = sortTasks(tasks, sortingMode)

	var processors Processors

	perms4 := permutation(rangeSlice(0, 4))
	perms42 := permutation(rangeSlice(4, 8))

	var processedTasks []Task
	var i int
	for i = 0; i <= (len(tasks)/4)*4; i+=4 {
		if i+8 > len(tasks) {
			break
		}
		processedTasks = tasks[i:i+8]

		bestPermutationIndexes := []int{0,0}
		minPenalty := 4294967295

		for permutation1Index, permutation1 := range perms4 {
			for permutation2Index, permutation2 := range perms42 {
				penalty := processors.CheckPermutations(processedTasks, permutation1, permutation2)
				if penalty < minPenalty {
					minPenalty = penalty
					bestPermutationIndexes = []int{permutation1Index, permutation2Index}
				}
			}
		}
		for index := range processors {
			processors[index].Append(processedTasks[perms4[bestPermutationIndexes[0]][index]])
		}
	}

	processedTasks = tasks[i:]
	var minProc int
	for _, task := range processedTasks {
		minProc = 0
		for i := 0; i < 4; i++ {
			if processors[i].Time < processors[minProc].Time {
				minProc = i
			}
		}
		processors[minProc].Append(task)
	}
	for index := range processors {
		scheduling.Tasks[index] = processors[index].Tasks.ToTasksNumbers()
	}
	scheduling.Penalty = processors.TotalPenalty()
	return scheduling
}

func rangeSlice(start, stop int) []int {
	if start > stop {
		panic("Slice ends before it started")
	}
	xs := make([]int, stop-start)
	for i := 0; i < len(xs); i++ {
		xs[i] = i + start
	}
	return xs
}

func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}

func sortTasks(tasks []Task, mode int) []Task {

	readyTime := func(t1, t2 *Task) bool { return t1.ReadyTime < t2.ReadyTime }
	deadline := func(t1, t2 *Task) bool { return t1.Deadline < t2.Deadline }
	duration := func(t1, t2 *Task) bool { return t1.Duration < t2.Duration }

	sortedTasks := make([]Task, len(tasks))
	copy(sortedTasks, tasks)

	if mode == 0 {
		OrderedBy(deadline, readyTime, duration).Sort(tasks)
	} else if mode == 1 {
		OrderedBy(deadline, duration, readyTime).Sort(tasks)
	} else if mode == 2 {
		OrderedBy(readyTime, duration, deadline).Sort(tasks)
	} else if mode == 3 {
		OrderedBy(readyTime, deadline, duration).Sort(tasks)
	} else if mode == 4 {
		OrderedBy(duration, deadline, readyTime).Sort(tasks)
	} else if mode == 5 {
		OrderedBy(duration, readyTime, deadline).Sort(tasks)
	}
	return tasks
}