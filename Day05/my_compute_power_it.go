//
// EPITECH PROJECT, 2025
// my_compute_power_it
// File description:
// day05
//

package main

import "fmt"

func my_compute_power_it(nb int, p int) int {
	result := nb

	if p == 0 {
		return 1
	}
	for i := 1; i < p; i++ {
		result *= nb
	}
	return result
}

// func main() {
// 	n := 2
// 	p := 3
// 	pow := my_compute_power_it(n, p)
// 	fmt.Printf("(%d, %d) => %d\n", n, p, pow)
// }
