package inout

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	. "structures"
)

func WriteScheduling(filename string, scheduling Scheduling) {
	err := ioutil.WriteFile(filename, scheduling.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func WriteTasks(filename string, tasks TaskList) {
	err := ioutil.WriteFile(filename, tasks.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func ReadScheduling(filename string, tasks TaskList) (penalty int, processors Processors) {
	var fd *os.File
	fd, err := os.Open(filename)
	if err != nil { panic(err) }

	reader := bufio.NewReader(fd)

	bytes, _,  err := reader.ReadLine()
	if err != nil { panic(err) }

	penalty, _ = strconv.Atoi(string(bytes))

	for i := 0; i < 4; i++ {
		line, err := reader.ReadString('\n')
		if err != nil { panic(err) }
		tasksStr := strings.Fields(line)
		for _, taskNo := range tasksStr {
			index, _ := strconv.Atoi(taskNo)
			processors[i].Append(tasks[index-1])
		}
	}
	return
}

func ReadTasks(filename string) (tasks TaskList) {
	fi, err := os.Open(filename)
	if err != nil { panic(err) }
	reader := bufio.NewReader(fi)
	instanceSizeStr, _,  err := reader.ReadLine()
	if err != nil { panic(err) }

	instanceSize, _ := strconv.Atoi(string(instanceSizeStr))
	for i := 0; i < instanceSize; i++ {
		line, err := reader.ReadString('\n')
		if err != nil { panic(err) }
		tasksStr := strings.Fields(line)
		p, _ := strconv.Atoi(tasksStr[0])
		r, _ := strconv.Atoi(tasksStr[1])
		d, _ := strconv.Atoi(tasksStr[2])
		tasks = append(tasks, Task{i+1,p,r,d})
	}
	return
}
