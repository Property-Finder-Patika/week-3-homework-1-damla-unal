package main

import (
	"fmt"
	"strings"
)

func main() {
	defineBasicArray()
	getAssignmentExampleArray()
	declareAndSetArray()
	fix()
	wizardPrinter()

}

func defineBasicArray() {
	// this declares an array with five integer elements
	// its length      : 5
	// its element type: int
	var tags [5]int        // initially set to the zero values
	fmt.Println(tags)      // [0 0 0 0 0]
	fmt.Println(len(tags)) //5

	// this array doesn't occupy any memory space (its length is zero)
	// its length      : 0
	// its element type: byte
	var zero [0]byte
	fmt.Printf("zero              : %#v\n", zero)

	// if an ellipsis ‘‘...’’ appears in place of the length, the array length is determined by the number of initializers
	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "9", GBP: "!", RMB: "s"}
	fmt.Printf("%T\n", symbol)    //[4]string
	fmt.Println(RMB, symbol[RMB]) //"3 "s"

	// an array r with 100 elements, all zero except for the last, which has value −1.
	r := [...]int{99: -1}
	fmt.Printf("%T\n", r) //[100]int

	// When comparing two arrays, their types should be identical
	// not comparable (type mismatch: length):
	// var blue = [3]int{6, 9, 3}
	// var red  = [5]int{6, 9, 3}
	// fmt.Println("Are they equal?", blue == red)

}

func getAssignmentExampleArray() {
	prev := [3]string{"Kafka's Revenge", "Stay Golden", "Everything Ship"}

	// You can't do this:
	// books = prev

	var books [4]string

	for i, b := range prev {
		books[i] += b + " 2nd Ed."
	}

	books[3] = "Awesomeness"

	fmt.Printf("last year:\n%#v\n", prev)
	fmt.Printf("\nthis year:\n%#v\n", books)
}

func declareAndSetArray() {

	var (
		bestFriends = [3]string{}
		distances   = [5]int{}
		data        = [5]byte{}
		ratios      = [1]float64{}
		alives      = [4]bool{}
		zero        = [0]byte{}
	)

	fmt.Println("******** Declare Empty Array ***********")
	// Declare and print the arrays with their types.
	fmt.Printf("names    : %#v\n", bestFriends)
	fmt.Printf("distances: %#v\n", distances)
	fmt.Printf("data     : %#v\n", data)
	fmt.Printf("ratios   : %#v\n", ratios)
	fmt.Printf("alives   : %#v\n", alives)
	fmt.Printf("zero     : %#v\n", zero)

	// Print only the types of the same arrays.
	fmt.Println("only the types of the same arrays")
	fmt.Printf("names    : %T\n", bestFriends)
	fmt.Printf("distances: %T\n", distances)
	fmt.Printf("data     : %T\n", data)
	fmt.Printf("ratios   : %T\n", ratios)
	fmt.Printf("alives   : %T\n", alives)
	fmt.Printf("zero     : %T\n", zero)

	// Print only the elements of the same arrays.
	fmt.Println("only the elements of the same arrays")
	fmt.Printf("names    : %q\n", bestFriends)
	fmt.Printf("distances: %d\n", distances)
	fmt.Printf("data     : %d\n", data)
	fmt.Printf("ratios   : %.2f\n", ratios)
	fmt.Printf("alives   : %t\n", alives)
	fmt.Printf("zero     : %d\n", zero)

	fmt.Println("******** Set Array ***********")

	distances[0] = 50
	distances[1] = 40
	distances[2] = 75
	distances[3] = 30
	distances[4] = 125

	separator := "\n" + strings.Repeat("=", 20) + "\n"

	fmt.Print("\ndistances", separator)
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}

	data[0] = 'H'
	data[1] = 'E'
	data[2] = 'L'
	data[3] = 'L'
	data[4] = 'O'

	fmt.Print("\ndata", separator)
	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%d]: %d\n", i, data[i])
	}

}

func fix() {
	// for the fixing, add comma between the elements of the names array
	// simplify this code(remove redundant empty string element from books)

	/*	var names [3]string = [3]string{
			"Einstein" "Shepard"
			"Tesla"
		}

		var books [5]string = [5]string{
			"Kafka's Revenge",
			"Stay Golden",
			"",
			"",
			""
		}

		fmt.Printf("%q\n", names)
		fmt.Printf("%q\n", books)

	*/

	var names = [3]string{
		"Einstein", "Shepard",
		"Tesla",
	}

	var books = [5]string{
		"Kafka's Revenge",
		"Stay Golden",
	}

	fmt.Printf("%q\n", names)
	fmt.Printf("%q\n", books)
}

func wizardPrinter() {
	fmt.Println("***************** Wizard Printer ********************")
	//created multi-dimensional array
	scientists := [...][3]string{
		{"First Name", "Last Name", "Nickname"},
		{"Albert", "Einstein", "emc2"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	// print this scientists array in pretty table form
	for i := range scientists {
		n := scientists[i]
		fmt.Printf("%-15s %-15s %-15s\n", n[0], n[1], n[2])

		if i == 0 {
			fmt.Println(strings.Repeat("=", 50))
		}
	}

}
