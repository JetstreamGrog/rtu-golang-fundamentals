package main

import (
	"fmt"
)

// Autors: Toms Gudlevskis
// Studenta apliecības numurs: 251RDB303

func main() {
	synonyms := make(map[string]map[string]bool)

	for {
		fmt.Print("> ")
		var command string
			if _, err := fmt.Scan(&command); err != nil {
			break
		}

			if command == "done" {
			break
		}

		switch command {
		case "add":
			var w1, w2 string
			fmt.Scan(&w1, &w2)
			
			if synonyms[w1] == nil {
				synonyms[w1] = make(map[string]bool)
			}
			synonyms[w1][w2] = true

			if synonyms[w2] == nil {
				synonyms[w2] = make(map[string]bool)
			}
			
			synonyms[w2][w1] = true



		case "count":

			var w string
			fmt.Scan(&w)
			fmt.Println(len(synonyms[w]))

		case "check":

			var w1, w2 string
			fmt.Scan(&w1, &w2)
			
			if synonyms[w1] != nil && synonyms[w1][w2] {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}
