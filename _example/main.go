package main

import (
	"fmt"
	"github.com/yousifnimah/NumToWordsGo/NumToWords"
)

func main() {
	input := 1011001
	words, err := NumToWords.Convert(input, "ar")
	if err != nil {
		return
	}
	fmt.Println(words)
}
