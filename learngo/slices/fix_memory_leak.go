package main

import (
	"fmt"
	"io/ioutil"

	"github.com/inancgumus/learngo/16-slices/exercises/24-fix-the-memory-leak/api"
)

func main() {
	// reports the initial memory usage
	api.Report()

	// returns a slice with 10 million elements.
	// it allocates 65 MB of memory space.
	millions := api.Read()

	// -----------------------------------------------------

	//copy the last 10 elements of the millions slice to the new slice.
	// create a new backing array
	last10 := make([]int, 0)
	last10 = append(last10, millions[len(millions)-10:]...)

	//last10 := millions[len(millions)-10:] // memory leak

	fmt.Printf("\nLast 10 elements: %d\n\n", last10)

	//overwrite the millions slice by simply assigning last10 slice to it
	// Make the millions slice lose reference to its backing array
	// so that its backing array can be cleaned up from memory.
	millions = last10
	api.Report()

	_, err := fmt.Fprintln(ioutil.Discard, millions[0])
	if err != nil {
		return
	}
}
