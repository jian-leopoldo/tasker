package main

import (
	"log"
	"os/exec"
	"strings"
	"sync"
)

func runScript(module string, file string, params []string) []byte {

	commandToRun := extractParams(params) + file

	data, err := exec.Command(module, commandToRun).Output()

	if err != nil {
		log.Fatal(err)
	}
	return data
}

func extractCommand(task taskStruct, wg *sync.WaitGroup) {
	wg.Add(1)
	mountTaskBody(task)
	wg.Done()
}

func mountTaskBody(task taskStruct) {
	printHeader(task)
	dataToReturn := runScript(task.Module, task.File, task.Params)
	writeLog(task.LogFile, string(dataToReturn))
	printFooter()
}

func extractParams(params []string) string {

	var sb strings.Builder

	for _, param := range params {
		sb.WriteString(param)
		sb.WriteString(" ")
	}

	return sb.String()
}

func runTasks(tasksToRun taskList, wg *sync.WaitGroup) {
	for _, service := range tasksToRun.Task {
		extractCommand(service, wg) //run async
	}
}
