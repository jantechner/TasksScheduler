package structures

import (
	"fmt"
	"sort"
)

type TaskList []Task

func (t TaskList) Len() int {
	return len(t)
}

func (t TaskList) Less(i, j int) bool {
	panic("implement me")
}

func (t TaskList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t TaskList) Bytes() []byte {
	var res string
	res += fmt.Sprintln(len(t))
	for _, task := range t {
		res += fmt.Sprintf("%d %d %d\n", task.Duration, task.ReadyTime, task.Deadline)
	}
	return []byte(res)
}

func (t TaskList) ToTasksNumbers() []int {
	res := make([]int, 0, len(t))
	for _, task := range t {
		res = append(res, task.Number)
	}
	return res
}

func (t TaskList) Str() (res string) {
	res += fmt.Sprintln(len(t))
	for _, task := range t {
		res += task.Str()
	}
	return
}


type lessFunc func(t1, t2 *Task) bool

type multiSorter struct {
	tasks 	TaskList
	less    []lessFunc
}

func (ms *multiSorter) Sort(tasks TaskList) {
	ms.tasks = tasks
	sort.Sort(ms)
}

func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{less: less}
}


func (ms *multiSorter) Len() int { return len(ms.tasks) }

func (ms *multiSorter) Swap(i, j int) { ms.tasks[i], ms.tasks[j] = ms.tasks[j], ms.tasks[i] }

func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.tasks[i], &ms.tasks[j]
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return ms.less[k](p, q)
}