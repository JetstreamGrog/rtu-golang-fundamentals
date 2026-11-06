package main

import (
	"fmt"
)

// Autors: Toms Gudlevskis
// Studenta apliecības numurs: 251RDB303

func main() {
	fmt.Println("First sequence of numbers")
	seq1 := readSequence()

	fmt.Println("Second sequence of numbers")
	seq2 := readSequence()

	result := filterSequences(seq1, seq2)

	printResult(result)
}

func readSequence() []int {
	var sequence []int
	for {
		fmt.Print("Input number: ")
		var n int
		_, err := fmt.Scanf("%d\n", &n)

		if err != nil {
			fmt.Println("input-output error")

			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if n == 0 {
			break
		}
		sequence = append(sequence, n)
	}
	return sequence
}

func filterSequences(seq1, seq2 []int) []int {
	var result []int

	inSeq2 := make(map[int]bool)
	for _, val := range seq2 {
		inSeq2[val] = true
	}

	for _, val := range seq1 {
		if !inSeq2[val] {
			result = append(result, val)
		}
	}
	return result
}

func printResult(result []int) {
	fmt.Print("Result:")
	for _, val := range result {
		fmt.Printf(" %d", val)
	}
	fmt.Println()
}
