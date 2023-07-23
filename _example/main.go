package main

import (
	"fmt"
	"github.com/yousifnimah/NumToWordsGo/NumToWords"
)

func main() {
	input := 901511
	words, err := NumToWords.Convert(input, "ar")
	if err != nil {
		return
	}
	fmt.Println(words)
}
