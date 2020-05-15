package main

import (
	"fmt"
	"math"
	"math/rand"
	. "structures"
	"time"
	. "inout"
)

var (
	MaxJobTime        = 8
	MaxDeadlineOffset = 10
	DeadlineBuffor    = 0
	JobTimeBuffor     = 5
)

func main() {
	rand.Seed(time.Now().Unix())
	for instanceSize := 50; instanceSize <= 500; instanceSize += 50 {
		var tasks []Task
		for i := 0; i < instanceSize; i++ {
			p := rand.Intn(MaxJobTime) + 1
			base := int(math.Floor(float64(i)/4.0))
			rx := []int{-1,1}[rand.Intn(2)] * rand.Intn(JobTimeBuffor)
			r := base + rx
			if r < 0 { r = 0}
			d := p + r + rand.Intn(MaxDeadlineOffset) + DeadlineBuffor
			tasks = append(tasks, Task{i, p, r, d})
			WriteTasks(fmt.Sprintf("my_instances/in132332_%d.txt", instanceSize), tasks)
		}
	}
}
