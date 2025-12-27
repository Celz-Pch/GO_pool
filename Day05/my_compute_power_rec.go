//
// EPITECH PROJECT, 2025
// my_compute_power_rec
// File description:
// day05
//

package main

import "fmt"

func my_compute_power_rec(nb int, p int) int {
	if p == 0 {
		return 1;
	}
	if p < 0 {
        return 0;
    } else {
        return nb * my_compute_power_rec(nb, p - 1);
    }
}

// func main() {
// 	n := 2
// 	p := 3
// 	pow := my_compute_power_rec(n, p)
// 	fmt.Printf("(%d, %d) => %d\n", n, p, pow)
// }
