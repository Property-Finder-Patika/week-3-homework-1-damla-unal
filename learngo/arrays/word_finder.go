package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

// search for the words inside the corpus
// to run this program, please run with this command -> go run main.go Cat Beginning
// output :
//     #2 : "cat"
//     #11: "beginning"
func main() {
	//get the search query from the command-line
	query := os.Args[1:]
	if len(query) == 0 {
		fmt.Println("Please enter word(s) to search.")
		return
	}

	// created filter array, not search for them
	filter := [...]string{
		"and", "or", "was", "the", "since", "very",
	}

	//converts a given string to a slice
	words := strings.Fields(corpus)

queries:
	for _, q := range query {
		q = strings.ToLower(q)

		for _, v := range filter {
			if q == v {
				continue queries
			}
		}

		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				break
			}
		}
	}
}
