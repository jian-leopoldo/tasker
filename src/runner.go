package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"tasker/runner"
)

func runShell(params []string, file string) []byte {

	commandToRun := extractParams(params) + file

	data, err := exec.Command("sh", commandToRun).Output()

	if err != nil {
		log.Fatal(err)
	}
	return data
}

func runScript(module string, file string, params []string) []byte {

	commandToRun := extractParams(params) + file

	data, err := exec.Command(module, commandToRun).Output()

	if err != nil {
		log.Fatal(err)
	}
	return data
}

func extractCommand(task TaskStruct) {
	if task.Module == "bash" {
		mountTaskBodyShell(task)
	} else if task.Module == "go" {
		mountTaskBodyGo(task)
	} else {
		mountTaskBody(task)

	}
}

func test(params string) string {
	fmt.Println(params)
	return params
}

func mountTaskBodyGo(task TaskStruct) {
	printHeader(task, runner.RunScriptGo)
	dataToReturn := runner.RunScriptGo(task.Module, task.File, task.Params)
	writeLog(task.LogFile, string(dataToReturn))
	printFooter()
}

func mountTaskBody(task TaskStruct) {
	printHeader(task, runner.RunScriptGo)
	dataToReturn := runScript(task.Module, task.File, task.Params)
	writeLog(task.LogFile, string(dataToReturn))
	printFooter()
}

func mountTaskBodyShell(task TaskStruct) {
	printHeader(task, runner.RunScriptGo)
	dataToReturn := runShell(task.Params, task.File)
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
