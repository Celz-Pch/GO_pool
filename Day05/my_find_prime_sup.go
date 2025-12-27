//
// EPITECH PROJECT, 2025
// my_find_prime_sup
// File description:
// day05
//

package main

import "fmt"

func my_find_prime_sup(nb int) int {
	is_prime := 0;

    for i := 1; i <= nb; i++ {
        if nb % i == 0 {
            is_prime++;
        }
    }
    if is_prime == 2 {
        return nb;
    } else {
        return my_find_prime_sup(nb + 1);
    }
}

// func main() {
// 	n := 13
// 	prime_sup := my_find_prime_sup(n)
// 	fmt.Printf("(%d, %d)\n", n, prime_sup)
// }
