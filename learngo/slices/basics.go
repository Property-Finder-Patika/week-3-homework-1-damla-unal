package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := "2 4 6 1 3 5"
	splitted := strings.Split(data, " ")

	var nums []int
	for _, s := range splitted {
		n, _ := strconv.Atoi(s)
		nums = append(nums, n)
	}

	//print slice
	fmt.Println("nums        :", nums)

	// Slice it for the even and odd numbers
	evens, odds := nums[:3], nums[3:]

	fmt.Println("evens       :", evens)
	fmt.Println("odds        :", odds)

	// the two numbers at the middle of slice
	fmt.Println("middle      :", nums[2:4])
	fmt.Println("first 2     :", nums[:2])
	fmt.Println("last 2      :", nums[len(nums)-2:])

	fmt.Println("evens last 1:", evens[len(evens)-1:])
	fmt.Println("odds last 2 :", odds[len(odds)-2:])

}
