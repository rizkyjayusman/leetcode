package main

import "fmt"

func main() {
	fmt.Println(countConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))
	fmt.Println(countConsistentStrings("abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}))
	fmt.Println(countConsistentStrings("cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}))
}

func countConsistentStrings(allowed string, words []string) int {
	count := 0
	for _, r := range words {
		for i := 0; i < len(r); i++ {
			if r[i] == allowed[0] {
				l := i + len(allowed)
				if l > len(r) {
					break
				}

				if l == len(r) {
					if allowed == string(r[i:]) {
						count++
					}
				} else {
					if allowed == string(r[i:i+len(allowed)]) {
						count++
					}
				}
			}
		}
	}

	return count
}
