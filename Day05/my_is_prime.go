//
// EPITECH PROJECT, 2025
// my_is_prime
// File description:
// day05
//

package main

import "fmt"

func my_is_prime(nb int) int {
    is_prime := 0;

    for i := 1; i <= nb; i++ {
        if nb % i == 0 {
            is_prime++;
        }
    }
    if is_prime == 2 {
        return 1;
    } else {
        return 0;
    }
}

// func main() {
// 	n := 7
// 	prime := my_is_prime(n)
// 	fmt.Printf("(%d, %d)\n", n, prime)
// }
