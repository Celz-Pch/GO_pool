//
// EPITECH PROJECT, 2025
// my_strlen
// File description:
// day04
//

package main

import "fmt"

func my_strlen(str string) int {
	count := 0

	for range str {
		count++
	}
	return count
}

// func main() {
// 	s := "Roh Roh"
// 	n := my_strlen(s)
// 	fmt.Printf("\nNombre de caractères affichés : %d\n", n)
// }
