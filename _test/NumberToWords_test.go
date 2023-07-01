package _test

import (
	"fmt"
	"github.com/yousifnimah/NumToWordsGo/NumToWords"
	"testing"
)

func TestNumberToWords(t *testing.T) {
	tests := []struct {
		number   int
		expected string
	}{
		{0, "zero"},
		{1, "one"},
		{10, "ten"},
		{42, "forty two"},
		{123, "one hundred twenty three"},
		{1000, "one thousand"},
		{9999, "nine thousand nine hundred ninety nine"},
		{1000000, "one million"},
		{123456789, "one hundred twenty three million four hundred fifty six thousand seven hundred eighty nine"},
		// Add more test cases here
	}

	for _, test := range tests {
		result, err := NumToWords.Convert(test.number, "en")
		if err != nil {
			t.Errorf(fmt.Sprintf("Somthing is wrong: %s", err))
		}
		if result != test.expected {
			t.Errorf("NumberToWords(%d) = %s; expected %s", test.number, result, test.expected)
		}
	}
}
