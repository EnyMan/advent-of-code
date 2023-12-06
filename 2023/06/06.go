package sixth

import (
	helpers "advent/utils"
	"fmt"
	"strings"
)

func parse_input(lines []string, part2 bool) int {
	var times []int
	var distances []int
	if part2 {
		times = []int{helpers.ToInt(strings.ReplaceAll(lines[0], " ", "")[5:])}
		distances = []int{helpers.ToInt(strings.ReplaceAll(lines[1], " ", "")[9:])}
	} else {
		times = helpers.ArrayToInt(strings.Fields(lines[0])[1:])
		distances = helpers.ArrayToInt(strings.Fields(lines[1])[1:])
	}
	options := []int{}
	for i := 0; i < len(times); i++ {
		first_valid := 0
		for offset := 0; offset < times[i]; offset++ {
			distance := offset * (times[i] - offset)
			if distance > distances[i] {
				first_valid = offset
				break
			}
		}
		last_valid := times[i] - first_valid
		options = append(options, last_valid-first_valid+1)
	}
	return helpers.Multiply(options)
}

func Sixth(input_filename string, part2 bool) {
	lines := helpers.ReadFile(input_filename)
	fmt.Println(parse_input(lines, part2))
}
