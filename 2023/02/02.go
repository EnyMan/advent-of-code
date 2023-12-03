package second

import (
	helpers "advent/utils"
	"fmt"
	"strconv"
	"strings"
)

func violates_rule(trial_sample string, part2 bool) bool {
	counts := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sample_splits := strings.Split(trial_sample, " ")
	count, color := sample_splits[0], sample_splits[1]
	countInt, err := strconv.Atoi(count)
	if err != nil {
		fmt.Println(err)
		return true
	}
	if countInt > counts[color] {
		return true
	}
	return false
}

func calcualte_game_power(trials []string, part2 bool) int {
	min_counts := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, trial := range trials {
		trial_splits := strings.Split(trial, ",")
		for _, trial_split := range trial_splits {
			trial_split := strings.TrimLeft(trial_split, " ")
			sample_splits := strings.Split(trial_split, " ")
			count, color := sample_splits[0], sample_splits[1]
			countInt, err := strconv.Atoi(count)
			if err != nil {
				fmt.Println(err)
			}
			if countInt > min_counts[color] {
				min_counts[color] = countInt
			}
		}
	}
	return min_counts["red"] * min_counts["green"] * min_counts["blue"]
}

func parse_input(input []string, part2 bool) []int {
	ok_ids := []int{}
	game_powers := []int{}
	game_violation := bool(false)
	for _, line := range input {
		// fmt.Println(line)
		line_splits := strings.Split(line, ":")
		header, game_data := line_splits[0], line_splits[1]
		header_splits := strings.Split(header, " ")
		game_id := header_splits[1]
		trials := strings.Split(game_data, ";")
		game_violation = false
		game_power := calcualte_game_power(trials, part2)
		for _, trial := range trials {
			trial_splits := strings.Split(trial, ",")
			for _, trial_split := range trial_splits {
				trial_split := strings.TrimLeft(trial_split, " ")
				if violates_rule(trial_split, part2) {
					game_violation = true
				}
				if game_violation {
					break
				}
			}
			if game_violation && !part2 {
				break
			}
		}
		game_idInt, err := strconv.Atoi(game_id)
		if err != nil {
			fmt.Println(err)
			break
		}
		if !game_violation {
			ok_ids = append(ok_ids, game_idInt)
		}
		game_powers = append(game_powers, game_power)
	}
	if part2 {
		return game_powers
	} else {
		return ok_ids
	}
}

func Second(filename string, part2 bool) {
	lines := helpers.ReadFile(filename)
	results := parse_input(lines, part2)
	sum := 0
	for _, result := range results {
		sum += result
	}
	fmt.Println(sum)
}
