package main

import (
	"fmt"
)

func printHeader(task TaskStruct, receiver func(module string, file string, params []string) []byte) {
	fmt.Println("---------------------------------------------------")
	fmt.Println(task.Name)
	fmt.Print("\n")
	receiver(task.Module, task.File, task.Params)
	fmt.Print("Running Task----------------------------------\n\n")
	fmt.Print("Task Name: ", task.Name)
	fmt.Print("\n")
}

func printFooter() {
	fmt.Print("\n\n")
	fmt.Println("---------------------------------------------------")
}
