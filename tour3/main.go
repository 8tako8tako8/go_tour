package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	split := strings.Fields(s)
	for _, word := range split {
		m[word] = m[word] + 1
	}
	return m
}

//Exercise: Maps
func main() {
	wc.Test(WordCount)
}
