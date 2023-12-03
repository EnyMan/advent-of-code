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
