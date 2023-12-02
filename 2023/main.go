package main

import (
	first "advent/01"
	second "advent/02"
	"flag"
	"fmt"
)

func main() {
	// Define flags
	intPtr := flag.Int("task", 1, "The task to run")

	flag.Parse()

	fmt.Println("int flag: ", *intPtr)

	switch *intPtr {
	case 1:
		first.First("01/01-input.txt")
	case 2:
		second.Second()
	default:
		fmt.Println("No task selected")
	}
}
