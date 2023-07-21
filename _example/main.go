package main

import (
	"fmt"
	"github.com/yousifnimah/NumToWordsGo/NumToWords"
)

func main() {
	input := 1000000000000
	words, err := NumToWords.Convert(input, "en")
	if err != nil {
		return
	}
	fmt.Println(words)
}
