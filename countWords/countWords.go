package main

import "strings"

func CountWords(s string) int {
	if s != "" {
		return len(strings.Split(s, " "))
	}
	return 0
}
