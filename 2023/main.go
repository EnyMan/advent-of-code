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
	boolPtr := flag.Bool("sample", false, "Run sample?")
	partPtr := flag.Bool("part2", true, "Run second part?")

	flag.Parse()

	switch *intPtr {
	case 1:
		if *boolPtr {
			first.First("01/sample.txt")
		} else {
			first.First("01/input.txt")
		}
	case 2:
		if *boolPtr {
			second.Second("02/sample.txt", *partPtr)
		} else {
			second.Second("02/input.txt", *partPtr)
		}
	default:
		fmt.Println("No task selected")
	}
}
