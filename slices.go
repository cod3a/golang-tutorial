package main

import "fmt"

func slicesAndArrays() {
	var arr [5]int
	arr2 := [6]int{24, 65, 11, 13, 27, 39}
	for i, value := range arr {
		fmt.Println(i, value)
		arr[i] = i
	}
	// slices:
	slice1 := make([]int, 5)
	slice2 := arr2[2:5]
	slice3 := arr2[:3]
	slice4 := append(slice2, 101, 202)
	slice5 := copy(slice3, slice4)

	fmt.Println(arr, arr2, slice1, slice2, slice3, slice4, len(slice4), slice5)
}
