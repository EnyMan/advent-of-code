package helpers

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(input_filename string) []string {
	file, err := os.Open(input_filename)
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

func ToInt(s string) int {
	s = strings.TrimSpace(s)
	var i int
	fmt.Sscanf(s, "%d", &i)
	return i
}

func IsNumber(s string) bool {
	_, err := fmt.Sscanf(s, "%d", new(int))
	return err == nil
}

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func Multiply(numbers []int) int {
	multiply := 1
	for _, number := range numbers {
		multiply *= number
	}
	return multiply
}

func ArrayToInt(numbers []string) []int {
	int_numbers := []int{}
	for _, number := range numbers {
		int_numbers = append(int_numbers, ToInt(number))
	}
	return int_numbers
}

func FindInList(needle string, haystack []string) bool {
	for _, element := range haystack {
		if needle == element {
			return true
		}
	}
	return false
}

func Enqueue(queue []string, element string) []string {
	queue = append(queue, element)
	// fmt.Println("Enqueued:", element)
	return queue
}

func Dequeue(queue []string) (string, []string) {
	element := queue[0]
	if len(queue) == 1 {
		var tmp = []string{}
		return element, tmp

	}

	return element, queue[1:]
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a int, b int, integers []int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i], integers[i+1:])
	}

	return result
}
