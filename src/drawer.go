package main

import (
	"fmt"
)

func printHeader(task taskStruct) {
	fmt.Println("---------------------------------------------------")
	fmt.Println(task.Name)
	fmt.Print("\n")
	fmt.Print("Running Task----------------------------------\n\n")
	fmt.Print("Task Name: ", task.Name)
	fmt.Print("\n")
}

func printFooter() {
	fmt.Print("\n\n")
	fmt.Println("---------------------------------------------------")
}
