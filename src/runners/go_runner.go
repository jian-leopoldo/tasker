package runner

import (
	"log"
	"os/exec"
)

// func mountBody(task taskStruct) {

// }

func RunScriptGo(module string, file string, params []string) []byte {

	goCommand := "go run"

	// commandToRun := core.extractParams(params) + file

	data, err := exec.Command(goCommand, "asd").Output()

	if err != nil {
		log.Fatal(err)
	}
	return data
}
