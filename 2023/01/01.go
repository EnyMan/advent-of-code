package first

import (
	helpers "advent/utils"
	"fmt"
	"strings"
	"unicode"
)

func replace_words_with_numbers(line string) string {
	digits := map[string]string{
		"one":   "o1ne",
		"two":   "tw2o",
		"three": "thr3ee",
		"four":  "fo4ur",
		"five":  "fi5ve",
		"six":   "s6ix",
		"seven": "sev7en",
		"eight": "eig8ht",
		"nine":  "ni9ne",
	}
	for word, number := range digits {
		line = strings.ReplaceAll(line, word, number)
	}
	return line
}

func parse_input(input []string, part2 bool) []int {

	calibration := []int{}

	// iterate over input
	for _, line := range input {
		line_numbers := []int{}
		// replace words with numbers
		if part2 {
			line = replace_words_with_numbers(line)
		}
		for _, char := range line {
			// add numbers to the line_numbers array
			if unicode.IsDigit(char) {
				line_numbers = append(line_numbers, int(char-'0'))
			}
		}
		//fmt.Println(line, line_numbers, get_first_and_last_number(line_numbers))
		calibration = append(calibration, get_first_and_last_number(line_numbers))
	}
	return calibration
}

func get_first_and_last_number(line []int) int {
	if len(line) == 0 {
		return 0
	}
	number := line[0]*10 + line[len(line)-1]
	return number

}

func First(filename string, part2 bool) {
	lines := helpers.ReadFile(filename)
	coords := parse_input(lines, part2)
	sum := 0
	for _, coord := range coords {
		sum += coord
	}
	fmt.Println(sum)
}
