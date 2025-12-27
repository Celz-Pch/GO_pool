//
// EPITECH PROJECT, 2025
// my_sort_int_array
// File description:
// day04
//

package main

import "fmt"

func my_sort_int_array(array []int, size int) {
	tmp := 0

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if array[i] > array[j] {
				tmp = array[j]
				array[j] = array[i]
				array[i] = tmp
			}
		}
	}
}

// func main() {
// 	arr := []int{5, 3, 8, 1, 2}
// 	my_sort_int_array(arr, len(arr))
// 	for _, v := range arr {
// 		fmt.Printf("%d ", v)
// 	}
// }
