package schedulers

import (
	"math/rand"

	. "schedulers"
	"sort"
	. "structures"
	"sync"
	. "verifier"
)

type population struct {
	size           int
	minPenalty     int
	bestScheduling Scheduling
	mutationChance float64
	population     []Scheduling
	tasks          []Task
}

func NewPopulation (mutationChance float64, size int, tasks []Task) *population {
	var newPopulation population
	newPopulation.minPenalty =- 1
	newPopulation.size = size
	newPopulation.population = make([]Scheduling, size*2)
	newPopulation.tasks = tasks
	newPopulation.mutationChance = mutationChance / float64(len(tasks)/50)
	return &newPopulation
}

func (p *population) generateSingle(i int, wg *sync.WaitGroup) {
	var scheduling Scheduling
	if i < 6 {
		scheduling = GeneratePermutationScheduling(p.tasks, i)
	} else {
		scheduling = GenerateRandomScheduling(len(p.tasks))
		scheduling.Penalty = CalculatePenaltyForScheduling(scheduling, p.tasks)
	}
	p.population[i] = scheduling
	wg.Done()
}

func (p *population) Generate() {
	var wg sync.WaitGroup
	for i := 0; i < p.size; i++ {
		wg.Add(1)
		go p.generateSingle(i, &wg)
	}
	wg.Wait()
}

func (p *population) mutateSingle(i int, wg *sync.WaitGroup) {
	s := p.population[i].Copy()
	for j := 0; j < 4; j++ {
		for i := 0; i < len(s.Tasks[j]); i++ {
			if rand.Float64() < p.mutationChance {
				whichProcessor := rand.Intn(4)
				whichPlace := rand.Intn(len(s.Tasks[j]) + 1)
				if whichPlace >= len(s.Tasks[whichProcessor]) {
					s.Tasks[whichProcessor] = append(s.Tasks[whichProcessor],s.Tasks[j][i])
				} else {
					s.Tasks[whichProcessor] = append(s.Tasks[whichProcessor], 0)
					copy(s.Tasks[whichProcessor][whichPlace+1:], s.Tasks[whichProcessor][whichPlace:])
					s.Tasks[whichProcessor][whichPlace] = s.Tasks[j][i]
				}
				s.Tasks[j] = append(s.Tasks[j][:i], s.Tasks[j][i+1:]...)
			}
		}
	}
	s.Penalty = CalculatePenaltyForScheduling(s, p.tasks)
	p.population[i + p.size] = s
	wg.Done()
}

func (p *population) Mutate() {
	var wg sync.WaitGroup
	for i := 0; i < p.size; i++ {
		wg.Add(1)
		go p.mutateSingle(i, &wg)
	}
	wg.Wait()
}

func (p *population) RecalculateMinPenalty() {
	sort.Slice(p.population, func(i, j int) bool {
		if p.population[i].Penalty < p.population[j].Penalty {
			return true
		}
		return false
	})

	//if p.minPenalty == -1 {
	//	defer fmt.Println("Min penalty: ", p.population[0].Penalty, p.population[0])
	//}
	//if p.minPenalty > p.population[0].Penalty {
	//	fmt.Println("Min penalty: ", p.population[0].Penalty)
	//}

	p.minPenalty = p.population[0].Penalty
	p.bestScheduling = p.population[0]
}

func (p *population) BestResult() Scheduling {
	p.bestScheduling.Penalty = p.minPenalty
	return p.bestScheduling
}

