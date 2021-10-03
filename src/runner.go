package main

import (
	"log"
	"os/exec"
	"strings"
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

func extractCommand(task taskStruct) {
	if task.Module == "bash" {
		mountTaskBodyShell(task)
	} else {
		mountTaskBody(task)
	}
}

func mountTaskBody(task taskStruct) {
	printHeader(task)
	dataToReturn := runScript(task.Module, task.File, task.Params)
	writeLog(task.LogFile, string(dataToReturn))
	printFooter()
}

func mountTaskBodyShell(task taskStruct) {
	printHeader(task)
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
