//
// EPITECH PROJECT, 2025
// my_swap
// File description:
// day04
//

package main

import "fmt"

func mySwap(a, b *int) {
	*a, *b = *b, *a
}

// func main() {
// 	x := 10
// 	y := 20

// 	fmt.Println("Avant swap :", x, y)
// 	mySwap(&x, &y)
// 	fmt.Println("Après swap  :", x, y)
// }
