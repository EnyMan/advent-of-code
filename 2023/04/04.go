package forth

import (
	helpers "advent/utils"
	"fmt"
	"math"
	"strings"
)

func convert_to_ints(numbers []string) []int {
	var int_numbers []int
	for _, number := range numbers {
		int_numbers = append(int_numbers, helpers.ToInt(number))
	}
	return int_numbers
}

func find_in_list(number int, list []int) bool {
	for _, item := range list {
		if item == number {
			return true
		}
	}
	return false
}

func parse_line(line string, part2 bool) [][]int {
	line_split := strings.Split(line, ":")
	header := line_split[0]
	card_number := helpers.ToInt(strings.Fields(header)[1])
	card := strings.Split(line_split[1], "|")
	winning_numbers := convert_to_ints(strings.Fields(card[0]))
	my_numbers := convert_to_ints(strings.Fields(card[1]))

	matches := []int{}
	for _, winning_number := range winning_numbers {
		if find_in_list(winning_number, my_numbers) {
			matches = append(matches, winning_number)
		}
	}
	clones := []int{}
	for index, _ := range matches {
		clones = append(clones, card_number+index+1)
	}
	return [][]int{{int(math.Pow(2, float64(len(matches)-1)))}, clones}
}

func parse_input(lines []string, part2 bool) []int {
	card_scores := []int{}
	all_clones := map[int]int{}

	for card_number, line := range lines {
		// fmt.Println(line)
		card_number++
		results := parse_line(line, part2)
		card_score, clones := results[0][0], results[1]
		card_scores = append(card_scores, card_score)
		if part2 {
			for _, clone := range clones {
				all_clones[clone]++
			}
			for len(all_clones) > 0 && all_clones[card_number] > 0 {
				card_scores = append(card_scores, card_score)
				for _, clone := range clones {
					all_clones[clone]++
				}
				all_clones[card_number]--
			}
		}
	}
	return card_scores
}

func Forth(filename string, part2 bool) {
	lines := helpers.ReadFile(filename)
	scores := parse_input(lines, part2)
	if part2 {
		fmt.Println(len(scores))
	} else {
		fmt.Println(helpers.Sum(scores))
	}
}
