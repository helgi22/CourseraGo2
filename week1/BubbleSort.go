package main

import (
	"fmt"
	"strconv"
)

func main() {
	var slice []int
	var input string

	for i := 0; i < 10; i++ {
		fmt.Printf("Enter yours number[%v]:", i+1)
		_, _ = fmt.Scanln(&input)
		if num, err := strconv.Atoi(input); err == nil {
			slice = append(slice, num)
		}
	}
	fmt.Printf("Original slice is: %v\n", slice)
	sort(&slice)
	fmt.Printf("Sorted slice slice is: %v\n", slice)
}

func sort(slice *[]int) {
	fmt.Printf("Slice include: %v\n", slice)
}

func swap(arr []int, idx int) {

}
