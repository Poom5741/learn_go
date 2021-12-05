package main

import "strings"

func main() {
	return wordCount("banana")
}

func wordCount(s string) map[string]int {
	split := strings.Split(s, " ")
	result := map[string]int{}
	for i := 0; i < len(split); i++ {
		result[split[i]] = result[split[i]] + i
	}
	return result
}
