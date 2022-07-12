package main

import (
	"fmt"
)

func main() {
	//create slice
	// its length is not part of its type
	var s1 []int // len(s) == 0, s == nil
	fmt.Printf("Length: %d and Capacity: %d and Is nil: %t\n", len(s1), cap(s1), s1 == nil)
	s1 = nil // len(s) == 0, s == nil
	fmt.Printf("Length: %d and Capacity: %d and Is nil: %t\n", len(s1), cap(s1), s1 == nil)
	s1 = []int(nil) // len(s) == 0, s == nil
	fmt.Printf("Length: %d and Capacity: %d and Is nil: %t\n", len(s1), cap(s1), s1 == nil)
	s1 = []int{} // len(s) == 0, s != nil
	fmt.Printf("Length: %d and Capacity: %d and Is nil: %t\n", len(s1), cap(s1), s1 == nil)

	fmt.Println("\n**************")
	var a [5]int = [...]int{1, 3, 5, 7, 9}
	var s []int = a[1:3]
	fmt.Println(s)                                              //[3 5]
	fmt.Printf("Length: %d and Capacity: %d\n", len(s), cap(s)) //Length: 2 and Capacity: 4

	//set different values
	a[0] = 10
	a[1] = 20
	a[2] = 30
	a[3] = 40
	a[4] = 50
	fmt.Println(s) // [20 30]

	// take another slice as whole array
	fmt.Println("Taking Slice As Whole Array")
	s = a[:]
	fmt.Println(s) // [10 20 30 40 50]
	fmt.Printf("Length: %d and Capacity: %d\n", len(s), cap(s))

	fmt.Println("\nModifying Slice")
	for i := 0; i < len(s); i++ {
		s[i] = 100 + i
	}
	fmt.Println(s) //[100 101 102 103 104]
	fmt.Println(a) //[100 101 102 103 104]

	modifyArray(a) //pass by value
	fmt.Println(a)

	fmt.Println("\n**************")
	modifySlice(s) // pass by reference
	fmt.Println(s)
	fmt.Println(a)

}

func modifyArray(a [5]int) {
	fmt.Println("Modifying Array")
	for i := 0; i < len(a); i++ {
		a[i] = -i
	}
}

func modifySlice(a []int) {
	fmt.Println("Modifying Slice")
	for i := 0; i < len(a); i++ {
		a[i] = -i
	}
}
