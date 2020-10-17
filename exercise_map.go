package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	word := strings.Fields(s)
	word_map := make(map[string]int)
	for _, w := range word {
		word_map[w]++
	}
	return word_map
}

func main() {
	wc.Test(WordCount)
}
