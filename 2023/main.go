package main

import (
	first "advent/01"
	second "advent/02"
	third "advent/03"
	forth "advent/04"
	"flag"
	"fmt"
)

func main() {
	// Define flags
	intPtr := flag.Int("task", 1, "The task to run")
	samplePtr := flag.Bool("sample", false, "Run sample?")
	partPtr := flag.Bool("part2", false, "Run second part?")

	flag.Parse()

	switch *intPtr {
	case 1:
		if *samplePtr {
			first.First("01/sample.txt", false)
			fmt.Println()
			fmt.Println("Part 2")
			fmt.Println()
			first.First("01/sample.txt", true)
		} else {
			first.First("01/input.txt", *partPtr)
		}
	case 2:
		if *samplePtr {
			second.Second("02/sample.txt", false)
			fmt.Println()
			fmt.Println("Part 2")
			fmt.Println()
			second.Second("02/sample.txt", true)
		} else {
			second.Second("02/input.txt", *partPtr)
		}
	case 3:
		if *samplePtr {
			third.Third("03/sample.txt", false)
			fmt.Println()
			fmt.Println("Part 2")
			fmt.Println()
			third.Third("03/sample.txt", true)
		} else {
			third.Third("03/input.txt", *partPtr)
		}
	case 4:
		if *samplePtr {
			forth.Forth("04/sample.txt", false)
			fmt.Println()
			fmt.Println("Part 2")
			fmt.Println()
			forth.Forth("04/sample.txt", true)
		} else {
			forth.Forth("04/input.txt", *partPtr)
		}
	default:
		fmt.Println("No task selected")
	}
}
