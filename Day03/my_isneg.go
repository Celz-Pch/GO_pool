//
// EPITECH PROJECT, 2025
// my_isneg
// File description:
// day03
//

package main

import "os"

func my_putchar(c byte) {
	os.Stdout.Write([]byte{c})
}

func my_isneg(n int) {
	if n < 0 {
		my_putchar('N')
	} else {
		my_putchar('P')
	}
	my_putchar('\n')
}

// func main() {
// 	my_isneg(10)
// 	my_isneg(-10)
// }
