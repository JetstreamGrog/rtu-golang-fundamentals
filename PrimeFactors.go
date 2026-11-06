package main

import (
	"fmt"
	"os"
	"strconv"
)

// Autors: Toms Gudlevskis
// Studenta apliecības numurs: 251RDB303

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			n, err := strconv.Atoi(arg)
			if err != nil || n < 2 {
				continue
			}
			factors := getPrimeFactors(n)
			printResult(n, factors)
		}
	} else {
		fmt.Print("Input number: ")
		var inputStr string
		_, err := fmt.Scanln(&inputStr)
		if err != nil {
			fmt.Println("input-output error")
			return
		}

		n, err := strconv.Atoi(inputStr)
		if err != nil || n < 2 {
			fmt.Println("input-output error")
			return
		}

		factors := getPrimeFactors(n)
		printResult(n, factors)
	}
}

func getPrimeFactors(n int) []int {
	var factors []int
	factor := 2
	for factor <= n {
		if n%factor == 0 {
			factors = append(factors, factor)
			n = n / factor
		} else {
			factor++
		}
	}
	return factors
}

func printResult(n int, factors []int) {
	fmt.Printf("%d:", n)
	for _, f := range factors {
		fmt.Printf(" %d", f)
	}
	fmt.Println()
}
