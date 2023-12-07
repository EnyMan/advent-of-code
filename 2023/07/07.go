package seventh

import (
	helpers "advent/utils"
	"fmt"
	"slices"
	"strings"
)

var CARD_VALUES = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var PART2 = false

func hand_type(hand string) int {
	cards := map[rune]int{}
	for _, card := range hand {
		cards[card] += 1
	}

	if PART2 && cards['J'] > 0 && len(cards) > 1 {
		max := 0
		var max_key rune
		for k, v := range cards {
			if k != 'J' && (v > max || (v == max && CARD_VALUES[k] > CARD_VALUES[max_key])) {
				max = v
				max_key = rune(k)
			}
		}
		cards[max_key] += cards['J']
		delete(cards, 'J')
	}

	card_type := 0
	for _, v := range cards {
		if len(cards) == 1 { // Five of a kind
			card_type = 7
			break
		} else if len(cards) == 2 && (v == 4 || v == 1) { // Four of a kind
			card_type = 6
			break
		} else if len(cards) == 2 && (v == 3 || v == 2) { // Full house
			card_type = 5
			break
		} else if len(cards) == 3 && v == 3 { // Three of a kind TTT98 QQQJA
			card_type = 4
		} else if len(cards) == 3 && card_type <= 3 { // Two pair 23432
			card_type = 3
		} else if len(cards) == 4 { // One pair
			card_type = 2
			break
		} else if card_type == 0 { // High card
			card_type = 1
		}
	}
	return card_type
}

func compare(a map[string]string, b map[string]string) int {
	if hand_type(a["hand"]) < hand_type(b["hand"]) {
		return -1
	} else if hand_type(a["hand"]) > hand_type(b["hand"]) {
		return 1
	} else {
		for i := 0; i < 5; i++ {
			if CARD_VALUES[rune(a["hand"][i])] < CARD_VALUES[rune(b["hand"][i])] {
				return -1
			} else if CARD_VALUES[rune(a["hand"][i])] > CARD_VALUES[rune(b["hand"][i])] {
				return 1
			}
		}
	}
	return 0
}

func parse_input(lines []string, part2 bool) []int {
	hands := []map[string]string{}
	for _, line := range lines {
		line_split := strings.Fields(line)
		hands = append(hands, map[string]string{"hand": line_split[0], "bid": line_split[1]})
	}
	position := []int{}
	slices.SortStableFunc(hands, compare)
	for index, hand := range hands {
		// fmt.Println(index+1, hand["bid"])
		position = append(position, (index+1)*helpers.ToInt(hand["bid"]))
	}
	return position
}

func Seventh(input_filename string, part2 bool) {
	if part2 {
		PART2 = true
		CARD_VALUES['J'] = 1
	}
	lines := helpers.ReadFile(input_filename)
	fmt.Println(helpers.Sum(parse_input(lines, part2)))
}
