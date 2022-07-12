package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"unsafe"
)

const size = 1e7

func main() {
	// don't worry about this code.
	// it stops the garbage collector: prevents cleaning up the memory.
	debug.SetGCPercent(-1)

	// run the program to see the initial memory usage.
	report("initial memory usage")

	var array [size]int
	report("after declaring an array")

	array2 := array
	report("after copying the array")

	passArray(array)

	slice1 := array[:]
	slice2 := array[:1e3]
	slice3 := array[1e3:1e4]
	// the size of the arrays and slices should be the same
	report("after slicings")

	passSlice(slice3)

	fmt.Println()
	fmt.Printf("Array's size : %d bytes.\n", unsafe.Sizeof(array))
	fmt.Printf("Array2's size: %d bytes.\n", unsafe.Sizeof(array2))
	fmt.Printf("Slice1's size: %d bytes.\n", unsafe.Sizeof(slice1))
	fmt.Printf("Slice2's size: %d bytes.\n", unsafe.Sizeof(slice2))
	fmt.Printf("Slice3's size: %d bytes.\n", unsafe.Sizeof(slice3))
}

// passes [size]int array — about 80MB!
// observe that passing an array to a function (or assigning it to a variable)
// affects the memory usage dramatically
func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

// only passes 24-bytes of slice header
// observe that passing a slice doesn't affect the memory usage
func passSlice(items []int) {
	items[0] = 100
	report("inside passSlice")
}

// reports the current memory usage
func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}
