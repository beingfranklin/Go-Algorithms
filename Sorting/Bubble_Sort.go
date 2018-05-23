package main

import (
	"fmt"
)
func bubbleSort(array_input []int){
	swapped := true;
	for swapped {
		swapped = false
		for i := 0; i < len(array_input) - 1; i++ {
			if array_input[i + 1] < array_input[i] {
				Swap(array_input, i, i + 1)
				swapped = true
			}
		}
	}	
}
func Swap(array_input []int, i, j int) {
	tmp := array_input[j]
	array_input[j] = array_input[i]
	array_input[i] = tmp
}
func main (){

	array_input := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted array: ", array_input)
	bubbleSort(array_input)
	fmt.Println("Sorted array: ", array_input)
	
}
