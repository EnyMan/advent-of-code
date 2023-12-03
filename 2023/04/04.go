package forth

import (
	helpers "advent/utils"
	"fmt"
)

func parse_input(lines []string, part2 bool) {
	fmt.Println(lines)
}

func Forth(filename string, part2 bool) {
	lines := helpers.ReadFile(filename)
	parse_input(lines, part2)
}
