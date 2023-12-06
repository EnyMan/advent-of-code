package seventh

import (
	helpers "advent/utils"
	"fmt"
)

func parse_input(lines []string, part2 bool) int {
	for _, line := range lines {
		fmt.Println(line)
	}
	return 0
}

func Seventh(input_filename string, part2 bool) {
	lines := helpers.ReadFile(input_filename)
	fmt.Println(parse_input(lines, part2))
}
