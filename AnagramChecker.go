package main

import (
	"fmt"
	"strings"
)

// Autors: Toms Gudlevskis
// Studenta apliecības numurs: 251RDB303

func main() {
	var word1, word2 string

	fmt.Print("First word: ")
	fmt.Scanln(&word1)

	fmt.Print("Second word: ")
	fmt.Scanln(&word2)

	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)

	counts1 := make(map[rune]int)
	counts2 := make(map[rune]int)

	for _, char := range word1 {
		counts1[char]++
	}

	for _, char := range word2 {
		counts2[char]++
	}

	isAnagram := true
	if len(counts1) != len(counts2) {
		isAnagram = false
	} else {
		for char, count := range counts1 {
			if counts2[char] != count {
				isAnagram = false
				break
			}
		}
	}

	if isAnagram {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
