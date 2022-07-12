package main

import (
	"fmt"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Fix the backing array problem
//
//  Ensure that changing the elements of the `mine` slice
//  does not change the elements of the `nums` slice.
//
//
// CURRENT OUTPUT (INCORRECT)
//
//  Mine         : [-50 -100 -150 25 30 50]
//  Original nums: [-50 -100 -150]
//
//
// EXPECTED OUTPUT
//
//  Mine         : [-50 -100 -150]
//  Original nums: [56 89 15]
//
// ---------------------------------------------------------
func main() {

	nums := []int{56, 89, 15, 25, 30, 50}

	//create different backing arrays
	mine := append([]int(nil), nums[:3]...)
	fmt.Println("First version: ", mine)

	mine[0], mine[1], mine[2] = -50, -100, -150

	fmt.Println("Mine         :", mine)
	fmt.Println("Original nums:", nums[:3])

	fmt.Println("************** Internals Backing Array Sort **************")
	sortBackingArray()
}

func sortBackingArray() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	fmt.Println("Original slice:", items)

	//find middle element
	mid := len(items) / 2
	smid := items[mid-1 : mid+2]

	// sorting the smid will affect the items as well. their backing array is the same.
	sort.Strings(smid)

	fmt.Println()
	fmt.Println("Sorted  :", items)
}
