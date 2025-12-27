//
// EPITECH PROJECT, 2025
// my_compute_factorial_rec
// File description:
// day05
//

package main

import "fmt"

func my_compute_factorial_rec(n int) int {
    if n < 0 || n > 12 {
        return 0;
    }
    if n == 0 {
        return 1;
    } else {
        return n * my_compute_factorial_rec(n - 1);
    }
}

func main() {
	n := 5
	fact := my_compute_factorial_rec(n)
	fmt.Printf("(%d, %d)\n", n, fact)
}
