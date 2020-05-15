package structures

import (
	"fmt"
)

type Scheduling struct {
	Penalty int
	Tasks 	[4][]int
	Size 	int
}

func NewScheduling(size int) Scheduling {
	var newScheduling Scheduling
	newScheduling.Penalty = 0
	newScheduling.Tasks[0] = make([]int, 0, size)
	newScheduling.Tasks[1] = make([]int, 0, size)
	newScheduling.Tasks[2] = make([]int, 0, size)
	newScheduling.Tasks[3] = make([]int, 0, size)
	newScheduling.Size = size
	return newScheduling
}

func (s Scheduling) Bytes() []byte {
	res := fmt.Sprintf("%d\n", s.Penalty)
	for _, p := range s.Tasks {
		for _, taskNumber := range p {
			res += fmt.Sprintf("%d ", taskNumber)
		}
		res = res[:len(res)-1]
		res += "\n"
	}
	return []byte(res)
}

func (s *Scheduling) Copy() Scheduling {
	newScheduling := NewScheduling(s.Size)
	for j := 0; j < 4; j++ {
		newScheduling.Tasks[j] = make([]int, len(s.Tasks[j]))
		copy(newScheduling.Tasks[j], s.Tasks[j])
	}
	return newScheduling
}
