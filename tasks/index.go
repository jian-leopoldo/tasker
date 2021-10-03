package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	for i := 0; i < 5; i++ {
		go say("world")
		go say("eae")
		say("hello")
	}

}
