package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)
	for _, word := range strings.Fields(s) {
		if _, ok := wordCount[word]; ok {
			wordCount[word] += 1
		} else {
			wordCount[word] = 1
		}
	}
	return wordCount
}

func runMapsExercise() {
	wc.Test(WordCount)
}
