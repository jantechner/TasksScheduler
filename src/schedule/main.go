package main

import (
	"fmt"

	//"inout"
	"os"
	. "schedulers"
	. "schedulers/heuristic"
	//"strings"
	"verifier"
)

func main() {

	var (
		indexes = []string{"132290", "132324", "132289", "132234", "132311", "132235", "132275", "132332", "132202", "132205", "132217", "132250", "132322", "132212", "116753", "132264", "132078"}
		schedulingFunction SchedulerFunction = GenerateNaiveScheduling
	)

	if args := os.Args[1:]; len(args) > 0 {
		switch args[0] {
		case "linear":
			schedulingFunction = GenerateLinearScheduling
			fmt.Print("Scheduling: linear")
		case "heuristic":
			schedulingFunction = GenerateHeuristicScheduling
			fmt.Print("Scheduling: heuristic")
		default:
			fmt.Print("Scheduling: naive")
		}
		if len(args) >= 2 {
			indexes = append([]string{}, args[1:]...)
			fmt.Println(", Indexes:", indexes)
		} else {
			fmt.Println(", Indexes: all")
		}
	} else {
		fmt.Println("Scheduling: naive, Indexes: all")
	}

	for _, index := range indexes {
		//for i := 50; i <= 500; i += 50 {
		//	filename := fmt.Sprintf("instances/in%s_%d.txt", index, i)
		//	if _, err := os.Stat(filename); os.IsNotExist(err) {
		//		fmt.Println()
		//	} else {
		//		scheduling, t := verifier.CalculatePenaltyForInstance(filename, schedulingFunction)
		//
		//		file := strings.Split(filename, "/")[1]
		//		inout.WriteScheduling(fmt.Sprintf("results/%s", strings.ReplaceAll(file, "in", "out")), scheduling)
		//
		//		//fmt.Println(t)
		//		fmt.Println(scheduling.Penalty, t)
		//	}
		//}
		i := 50
		filename := fmt.Sprintf("instances/in%s_%d.txt", index, i)
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			fmt.Println()
		} else {
			scheduling, t := verifier.CalculatePenaltyForInstance(filename, schedulingFunction)

			//file := strings.Split(filename, "/")[1]
			//inout.WriteScheduling(fmt.Sprintf("results/%s", strings.ReplaceAll(file, "in", "out")), scheduling)

			//fmt.Println(t)
			fmt.Println(scheduling.Penalty, t)
		}

	}
}
