package first

import "testing"

func TestReplaceWordsWithNumbers(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "2tqbxgrrpmxqfglsqjkqthree6nhjvbxpflhr1eightwohr",
			expected: "2tqbxgrrpmxqfglsqjkqthr3ee6nhjvbxpflhr1eig8htw2ohr",
		},
		{
			input:    "fourfivesix",
			expected: "fo4urfi5ves6ix",
		},
		{
			input:    "seveneightnine",
			expected: "sev7eneig8htni9ne",
		},
		{
			input:    "onetwothreefourfivesixseveneightnine",
			expected: "o1netw2othr3eefo4urfi5ves6ixsev7eneig8htni9ne",
		},
		{
			input:    "onehundredtwenty",
			expected: "o1nehundredtwenty",
		},
	}

	for _, test := range tests {
		result := replace_words_with_numbers(test.input)
		if result != test.expected {
			t.Errorf("Input: %s, Expected: %s, Got: %s", test.input, test.expected, result)
		}
	}
}
