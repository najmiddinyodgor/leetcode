package main

import (
	"strings"
)

func UncommonFromSentences(s1 string, s2 string) []string {
	r, dict := make([]string, 0), make(map[string]int)

	for _, w := range strings.Split(s1+" "+s2, " ") {
		dict[w]++
	}

	for key, value := range dict {
		if value == 1 {
			r = append(r, key)
		}
	}

	return r
}
