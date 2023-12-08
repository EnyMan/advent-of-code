package eighth

import (
	helpers "advent/utils"
	"fmt"
	"strings"
)

func part_1(instructions string, nodes map[string][]string) int {
	current_node := "AAA"
	steps_taken := 0
	for current_node != "ZZZ" {
		for instruction := range instructions {
			if instructions[instruction] == 'R' {
				current_node = nodes[current_node][1]
				steps_taken++
			} else if instructions[instruction] == 'L' {
				current_node = nodes[current_node][0]
				steps_taken++
			} else {
				fmt.Println("Unknown instruction:", string(instructions[instruction]))
			}
		}
	}
	return steps_taken
}

func final_nodes(nodes []string) bool {
	for _, node := range nodes {
		if node[2] != 'Z' {
			return false
		}
	}
	return true
}

func part_2(instructions string, nodes map[string][]string) int {
	var current_nodes = make([]string, 0)
	for node_name, _ := range nodes {
		if node_name[2] == 'A' {
			current_nodes = helpers.Enqueue(current_nodes, node_name)
		}
	}
	steps_taken := 0
	sub_steps_taken := []int{}
	for !final_nodes(current_nodes) {
		for instruction := range instructions {
			var new_nodes = make([]string, 0)
			for _, current_node := range current_nodes {
				current_node, current_nodes = helpers.Dequeue(current_nodes)
				if instructions[instruction] == 'R' {
					if nodes[current_node][1][2] == 'Z' {
						sub_steps_taken = append(sub_steps_taken, steps_taken+1)
					} else {
						new_nodes = helpers.Enqueue(new_nodes, nodes[current_node][1])
					}
				} else if instructions[instruction] == 'L' {
					if nodes[current_node][0][2] == 'Z' {
						sub_steps_taken = append(sub_steps_taken, steps_taken+1)
					} else {
						new_nodes = helpers.Enqueue(new_nodes, nodes[current_node][0])
					}
				} else {
					fmt.Println("Unknown instruction:", string(instructions[instruction]))
				}
			}
			steps_taken++
			current_nodes = new_nodes
		}
	}
	return helpers.LCM(sub_steps_taken[0], sub_steps_taken[1], sub_steps_taken[1:])
}

func parse_input(lines []string, part2 bool) int {
	instructions := ""
	nodes := map[string][]string{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		if instructions == "" {
			instructions = line
		} else {
			split := strings.Fields(line)
			nodes[split[0]] = []string{split[2][1:4], split[3][:3]}
		}
	}
	if part2 {
		return part_2(instructions, nodes)
	} else {
		return part_1(instructions, nodes)
	}
}

func Eighth(input_filename string, part2 bool) {
	lines := helpers.ReadFile(input_filename)
	fmt.Println(parse_input(lines, part2))
}
