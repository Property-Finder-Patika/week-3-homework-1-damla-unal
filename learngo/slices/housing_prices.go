package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)
	//create slices
	var (
		locs                       []string
		sizes, beds, baths, prices []int
	)

	// split it into rows by using the newline character
	rows := strings.Split(data, "\n")

	for _, row := range rows {
		// split it by using the separator (",")
		cols := strings.Split(row, separator)

		locs = append(locs, cols[0])

		n, _ := strconv.Atoi(cols[1])
		sizes = append(sizes, n)

		n, _ = strconv.Atoi(cols[2])
		beds = append(beds, n)

		n, _ = strconv.Atoi(cols[3])
		baths = append(baths, n)

		n, _ = strconv.Atoi(cols[4])
		prices = append(prices, n)
	}

	for _, h := range strings.Split(header, separator) {
		fmt.Printf("%-15s", h)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	for i := range rows {
		fmt.Printf("%-15s", locs[i])
		fmt.Printf("%-15d", sizes[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d", prices[i])
		fmt.Println()
	}

	//print housing prices averages
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))
	// jump over the location column
	fmt.Printf("%-15s", "")

	var sum int

	for _, n := range sizes {
		sum += n
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(len(sizes)))

	sum = 0
	for _, n := range beds {
		sum += n
	}
	fmt.Printf("%-15.2f", float64(sum)/float64(len(beds)))
}
