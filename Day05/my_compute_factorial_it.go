//
// EPITECH PROJECT, 2025
// my_compute_factorial_it
// File description:
// day05
//

package main

import "fmt"

func my_compute_factorial_it(nb int) int {
	if nb < 0 {
		return 0
	}
	result := 1
	for i := 1; i <= nb; i++ {
		result *= i
	}
	return result
}

// func main() {
// 	n := 5
// 	fact := my_compute_factorial_it(n)
// 	fmt.Printf("(%d, %d)\n", n, fact)
// }
