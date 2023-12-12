package main

import (
	first "advent/01"
	second "advent/02"
	third "advent/03"
	forth "advent/04"
	fifth "advent/05"
	sixth "advent/06"
	seventh "advent/07"
	eighth "advent/08"
	ninth "advent/09"
	"flag"
	"fmt"
)

func PrintPart2() {
	fmt.Println()
	fmt.Println("Part 2")
	fmt.Println()
}

func main() {
	// Define flags
	intPtr := flag.Int("task", 1, "The task to run")
	samplePtr := flag.Bool("sample", false, "Run sample?")
	// partPtr := flag.Bool("part2", false, "Run second part?")

	flag.Parse()

	switch *intPtr {
	case 1:
		if *samplePtr {
			first.First("01/sample.txt", false)
			PrintPart2()
			first.First("01/sample.txt", true)
		} else {
			first.First("01/input.txt", false)
			PrintPart2()
			first.First("01/input.txt", true)
		}
	case 2:
		if *samplePtr {
			second.Second("02/sample.txt", false)
			PrintPart2()
			second.Second("02/sample.txt", true)
		} else {
			second.Second("02/input.txt", false)
			PrintPart2()
			second.Second("02/input.txt", true)
		}
	case 3:
		if *samplePtr {
			third.Third("03/sample.txt", false)
			PrintPart2()
			third.Third("03/sample.txt", true)
		} else {
			third.Third("03/input.txt", false)
			PrintPart2()
			third.Third("03/input.txt", true)
		}
	case 4:
		if *samplePtr {
			forth.Forth("04/sample.txt", false)
			PrintPart2()
			forth.Forth("04/sample.txt", true)
		} else {
			forth.Forth("04/input.txt", false)
			PrintPart2()
			forth.Forth("04/input.txt", true)
		}
	case 5:
		if *samplePtr {
			fifth.Fifth("05/sample.txt", false)
			PrintPart2()
			fifth.Fifth("05/sample.txt", true)
		} else {
			fifth.Fifth("05/input.txt", false)
			PrintPart2()
			fifth.Fifth("05/input.txt", true)
		}
	case 6:
		if *samplePtr {
			sixth.Sixth("06/sample.txt", false)
			PrintPart2()
			sixth.Sixth("06/sample.txt", true)
		} else {
			sixth.Sixth("06/input.txt", false)
			PrintPart2()
			sixth.Sixth("06/input.txt", true)
		}
	case 7:
		if *samplePtr {
			seventh.Seventh("07/sample.txt", false)
			PrintPart2()
			seventh.Seventh("07/sample.txt", true)
		} else {
			seventh.Seventh("07/input.txt", false)
			PrintPart2()
			seventh.Seventh("07/input.txt", true)
		}
	case 8:
		if *samplePtr {
			eighth.Eighth("08/sample.txt", false)
			PrintPart2()
			eighth.Eighth("08/sample2.txt", true)
		} else {
			eighth.Eighth("08/input.txt", false)
			PrintPart2()
			eighth.Eighth("08/input.txt", true)
		}
	case 9:
		if *samplePtr {
			ninth.Ninth("09/sample.txt", false)
			PrintPart2()
			ninth.Ninth("09/sample.txt", true)
		} else {
			ninth.Ninth("09/input.txt", false)
			PrintPart2()
			ninth.Ninth("09/input.txt", true)
		}
	default:
		fmt.Println("No task selected")
	}
}
