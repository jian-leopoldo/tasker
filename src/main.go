package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type TaskStruct struct {
	Name    string   `yaml:"name"`
	Module  string   `yaml:"module"`
	File    string   `yaml:"file"`
	LogFile string   `yaml:"log_file"`
	Params  []string `yaml:"params"`
}

type TaskList struct {
	Task []TaskStruct `yaml:"task"`
}

func readConf(filename string) (*TaskList, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &TaskList{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}

func main() {

	tasksToRun, err := readConf("tasker.yaml")

	if err != nil {
		log.Fatal(err)
	}

	runTasks(*tasksToRun)
}

func runTasks(tasksToRun TaskList) {
	for _, service := range tasksToRun.Task {
		extractCommand(service) //run async
	}
}
