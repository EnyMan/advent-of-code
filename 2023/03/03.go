package third

import (
	helpers "advent/utils"
	"fmt"
	"strings"
)

func find_left(line string, pos int) string {
	pos--
	number := ""
	for pos >= 0 && line[pos] != '.' {
		number = string(line[pos]) + number
		pos--
	}
	return number
}

func find_right(line string, pos int) string {
	pos++
	number := ""
	for pos < len(line) && line[pos] != '.' {
		number += string(line[pos])
		pos++
	}
	return strings.TrimSpace(number)
}

func find_off(line string, pos int) []string {
	number := ""
	number_right := ""
	number_left := ""
	left_end := false
	right_end := false
	// fmt.Println("  ", line, pos)
	for i := 0; i < len(line); i++ {
		is_direct := helpers.IsNumber(string(line[pos]))
		if is_direct {
			// middle number
			if i == 0 {
				number = string(line[pos])
				continue
			}
			if !left_end && pos-i >= 0 && helpers.IsNumber(string(line[pos-i])) {
				number = string(line[pos-i]) + number
			} else {
				left_end = true
			}
			if !right_end && pos+i < len(line)-1 && helpers.IsNumber(string(line[pos+i])) {
				number = number + string(line[pos+i])
			} else {
				right_end = true
			}
		} else {
			// left and right at the same time
			if !left_end && pos-i-1 >= 0 && helpers.IsNumber(string(line[pos-i-1])) {
				number_left = string(line[pos-i-1]) + number_left
			} else {
				left_end = true
			}
			if !right_end && pos+i+1 < len(line)-1 && helpers.IsNumber(string(line[pos+i+1])) {
				number_right = number_right + string(line[pos+i+1])
			} else {
				right_end = true
			}
		}
		if left_end && right_end {
			continue
		}
	}
	return []string{number, number_left, number_right}
}

func parse_input(lines []string, part2 bool) [][]int {
	global_numbers := [][]int{}
	gears := [][]int{}
	for i := 0; i < len(lines); i++ {
		// fmt.Println(lines[i])
		specials := []int{}
		// go through each character
		for j := 0; j < len(lines[i]); j++ {
			char := lines[i][j]
			// find all special symbols
			if (char < '0' || char > '9') && char != '.' {
				// fmt.Println(string(char), j)
				specials = append(specials, j)
			}
		}
		// find gears
		numbers := []int{}
		for _, special := range specials {
			local_numbers := []int{}
			// find all conected numbers
			// go left
			left_number := helpers.ToInt(find_left(lines[i], special))
			if left_number > 0 {
				local_numbers = append(local_numbers, left_number)
			}
			// go right
			right_number := helpers.ToInt(find_right(lines[i], special))
			if right_number > 0 {
				local_numbers = append(local_numbers, right_number)
			}
			// go down
			if i < len(lines)-1 {
				down_numbers := find_off(lines[i+1], special)
				for _, down_number := range down_numbers {
					down_number := helpers.ToInt(down_number)
					if down_number > 0 {
						local_numbers = append(local_numbers, down_number)
					}
				}
			}
			// go up
			if i-1 >= 0 {
				up_numbers := find_off(lines[i-1], special)
				for _, up_number := range up_numbers {
					up_number := helpers.ToInt(up_number)
					if up_number > 0 {
						local_numbers = append(local_numbers, up_number)
					}
				}
			}
			if lines[i][special] == '*' && len(local_numbers) == 2 {
				// fmt.Println("Found Gear", local_numbers)
				gears = append(gears, local_numbers)
			}
			numbers = append(numbers, local_numbers...)
		}
		global_numbers = append(global_numbers, numbers)
	}
	// fmt.Println(global_numbers)
	// fmt.Println(gears)
	if part2 {
		return gears
	} else {
		return global_numbers
	}

}

func Third(filename string, part2 bool) {
	lines := helpers.ReadFile(filename)
	global_numbers := parse_input(lines, part2)
	sum := 0
	for _, line_numbers := range global_numbers {
		if part2 {
			sum += helpers.Multiply(line_numbers)
		} else {
			sum += helpers.Sum(line_numbers)
		}
	}
	fmt.Println(sum)
}
