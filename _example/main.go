package main

import (
	"fmt"
	"github.com/yousifnimah/NumToWordsGo/NumToWords"
)

func main() {
	input := 1500000
	words, err := NumToWords.Convert(input, "en")
	if err != nil {
		return
	}
	fmt.Println(words)
}
