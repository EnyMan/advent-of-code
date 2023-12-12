package ninth

import (
	helpers "advent/utils"
	"fmt"
	"strings"
)

func subtract(sequence []int) []int {
	result := make([]int, len(sequence)-1)
	for i := 0; i < len(sequence)-1; i++ {
		result[i] = sequence[i+1] - sequence[i]
	}

	return result
}

func parse_input(lines []string, part2 bool) int {
	predictions := []int{}
	for _, line := range lines {
		sequences := [][]int{}
		sequences = append(sequences, helpers.ArrayToInt(strings.Fields(line)))
		for helpers.Sum(sequences[len(sequences)-1]) != 0 {
			sequences = append(sequences, subtract(sequences[len(sequences)-1]))
		}

		for i := len(sequences) - 1; i >= 0; i-- {
			if helpers.Sum(sequences[i]) == 0 {
				if part2 {
					sequences[i] = append([]int{0}, sequences[i]...)
				}
				sequences[i] = append(sequences[i], 0)

			} else {
				next_number := 0
				if part2 {
					next_number = sequences[i][0] - sequences[i+1][0]
					sequences[i] = append([]int{next_number}, sequences[i]...)
				} else {
					next_number = sequences[i+1][len(sequences[i])-1] + sequences[i][len(sequences[i])-1]
					sequences[i] = append(sequences[i], next_number)
				}
				if i == 0 {
					predictions = append(predictions, next_number)
				}
			}
		}
	}
	return helpers.Sum(predictions)
}

func Ninth(input_filename string, part2 bool) {
	lines := helpers.ReadFile(input_filename)
	fmt.Println(parse_input(lines, part2))
}
