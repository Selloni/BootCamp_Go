package main

import (
	"errors"
	"fmt"
	"log"
	"unsafe"
)

func main() {
	arr := []int{2, 45, 6, 32, 17, 13, 5}
	res, err := getElement(arr, 6)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func getElement(arr []int, idx int) (int, error) {
	if idx > len(arr) || idx < 0 {
		return 0, errors.New("idx out of range")
	}
	// An array of contiguous uint32 values stored in memory.
	var a int
	// The number of bytes each uint32 occupies: 4.
	const size = unsafe.Sizeof(int(0))
	// Take the initial memory address of the array and begin iteration.
	p := uintptr(unsafe.Pointer(&arr[0]))
	for i := 0; i < idx; i++ {
		// Print the integer that resides at the current address and then
		// increment the pointer to the next value in the array.
		a = (*(*int)(unsafe.Pointer(p)))
		p += size
	}
	return a, nil
}
