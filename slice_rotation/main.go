package main

import "fmt"

//Write a function that rotates given slice by given number of steps to the right or to the left
func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(rotateRight(slice, 3)) //[3 4 5 1 2]
	fmt.Println(rotateLeft(slice, 3))  //[4 5 1 2 3]

}

func rotateRight(slice []int, step int) []int {
	lenSlice := len(slice)

	// check step number is equal to slice len or step is zero
	if lenSlice == step || step == 0 {
		return slice
	}

	var newSlice []int
	//first part
	newSlice = append(newSlice, slice[(lenSlice-step):]...)
	//remaining part
	newSlice = append(newSlice, slice[:(lenSlice-step)]...)

	return newSlice

}

func rotateLeft(slice []int, step int) []int {
	lenSlice := len(slice)

	// check step number is equal to slice len or step is zero
	if lenSlice == step || step == 0 {
		return slice
	}

	var newSlice []int
	// first part
	newSlice = append(newSlice, slice[step:]...)
	// remaining part
	newSlice = append(newSlice, slice[:step]...)

	return newSlice

}
