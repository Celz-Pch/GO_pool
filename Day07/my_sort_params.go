//
// EPITECH PROJECT, 2025
// my_sort_params
// File description:
// day07
//

package main

import (
	"os"
)

func my_putchar(c byte) {
	os.Stdout.Write([]byte{c})
}

func my_putstr(str string) int {
	count := 0

	for _, c := range str {
		os.Stdout.Write([]byte(string(c)))
		count++
	}
	return count
}

func my_print_params(args []string) {
	for i := 0; i < len(args); i++ {
		my_putstr(args[i])
		my_putchar('\n')
	}
}

func sortStrings(arr []string) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func my_sort_params(args []string) {
	sortStrings(args[1:])
	my_print_params(args)
}

// func main() {
// 	args := os.Args[0:]
// 	my_sort_params(args)
// }