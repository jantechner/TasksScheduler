package structures

type Processor struct {
	Time    int
	Tasks   TaskList
	Penalty int
}

func (p Processor) TasksNumbersList() []int {
	res := make([]int, 0, len(p.Tasks))
	for _, task := range p.Tasks {
		res = append(res, task.Number)
	}
	return res
}

func (p *Processor) Append(task Task) {
	p.Tasks = append(p.Tasks, task)
	p.Time = Max(p.Time, task.ReadyTime) + task.Duration
	if task.Deadline < p.Time {
		p.Penalty += p.Time - task.Deadline
	}
}

func (p *Processor) CalculatePenalty() (penalty int) {
	time := 0
	for _, task := range p.Tasks {
		time = Max(time, task.ReadyTime) + task.Duration
		if task.Deadline < time { penalty += time - task.Deadline }
	}
	return
}

func Max(a, b int) int {
	if a >= b { return a } else {return b}
}

type Processors [4]Processor


func (p Processors) GenerateScheduling() (scheduling Scheduling) {
	for i, processor := range p {
		scheduling.Tasks[i] = processor.TasksNumbersList()
	}
	scheduling.Penalty = p.TotalPenalty()
	return
}

func (p Processors) TotalPenalty() (totalPenalty int) {
	for _, processor := range p {
		totalPenalty += processor.CalculatePenalty()
	}
	return
}

func (p Processors) CheckPermutations(processedTasks []Task, permutation1 []int, permutation2 []int) (penalty int) {
	for index, processor := range p {
		task1 := processedTasks[permutation1[index]]
		task2 := processedTasks[permutation1[index]]

		if task2.Deadline < task1.Deadline {
			task1, task2 = task2, task1
		}


		time := Max(processor.Time, task1.ReadyTime) + task1.Duration
		if task1.Deadline < time {
			penalty += time - task1.Deadline
		}

		time = Max(time, task2.ReadyTime) + task2.Duration
		if task2.Deadline < time {
			penalty += time - task2.Deadline
		}
	}
	return
}
