package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	//strings.Split function splits a string and returns a string slice
	namesASlice := strings.Split(namesA, ", ")
	//sort slices
	sort.Strings(namesASlice)
	sort.Strings(namesB)

	if equal(namesB, namesASlice) {
		fmt.Println("They are equal") //expected output!
	} else {
		fmt.Println("They are not equal")
	}

}

func equal(slice1 []string, slice2 []string) bool {
	//check whether their length are the same or not
	if len(slice1) != len(slice1) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
