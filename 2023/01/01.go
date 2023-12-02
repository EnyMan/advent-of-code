package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func read_file() []string {
	file, err := os.Open("01-input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	// get file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file size:", err)
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	lines := string(bs)

	// split lines on new line to create a array of strings
	return strings.Split(lines, "\n")
}

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

func parse_input(input []string) []int {

	calibration := []int{}

	// iterate over input
	for _, line := range input {
		line_numbers := []int{}
		// replace words with numbers
		line := replace_words_with_numbers(line)
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

func main() {
	lines := read_file()
	coords := parse_input(lines)
	sum := 0
	for _, coord := range coords {
		sum += coord
	}
	fmt.Println(sum)
}
