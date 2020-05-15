package schedulers

import (
	"math/rand"
	. "structures"
)

func GenerateRandomScheduling(size int) Scheduling {
	scheduling := NewScheduling(size)

	candidates := make([]int, size)
	for i := 0; i < size; i++ {
		candidates[i] = i + 1
	}
	rand.Shuffle(len(candidates), func(i, j int) { candidates[i], candidates[j] = candidates[j], candidates[i] })

	for i := 0; i < size; i++ {
		j := rand.Intn(4)
		scheduling.Tasks[j] = append(scheduling.Tasks[j], candidates[i])
	}
	return scheduling
}
