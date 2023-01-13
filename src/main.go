package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"gopkg.in/yaml.v3"
)

const (
	DefaultSchema = "tasker.yaml"
)

type taskStruct struct {
	Name    string   `yaml:"name"`
	Module  string   `yaml:"module"`
	File    string   `yaml:"file"`
	LogFile string   `yaml:"log_file"`
	Params  []string `yaml:"params"`
}

type taskList struct {
	Task []taskStruct `yaml:"task"`
}

func readConf(filename string) (*taskList, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &taskList{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}

func main() {

	tasksToRun, err := readConf(DefaultSchema)

	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup

	runTasks(*tasksToRun, &wg)

	wg.Wait()
}
