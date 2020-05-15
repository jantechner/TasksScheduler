package structures

import "fmt"

type Task struct {
	Number 		int
	Duration	int
	ReadyTime	int
	Deadline	int
}

func NewTask(p, r, d, n int) *Task {
	return &Task{n, p, r, d}
}

func (t Task) Str() string {
	return fmt.Sprintf("(%d %d %d)\n", t.Duration, t.ReadyTime, t.Deadline)
}