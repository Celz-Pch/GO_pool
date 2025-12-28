//
// EPITECH PROJECT, 2025
// my_strncat
// File description:
// day07
//

package main

import "fmt"

func my_strncat(dest *string, src string, n int) {
	if n > len(src) {
		n = len(src)
	}
	*dest += src[:n]
}

// func main() {
// 	dest := "Hello, "
// 	src := "World!"
// 	my_strncat(&dest, src, 3)
// 	fmt.Println(dest) // Output: Hello, Wor
// }