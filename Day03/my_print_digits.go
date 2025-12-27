//
// EPITECH PROJECT, 2025
// my_printdigits
// File description:
// day03
//

package main

import "os"

func my_putchar(c byte) {
	os.Stdout.Write([]byte{c})
}

func my_print_digits() {
	for i := byte(48); i <= byte(57); i++ {
		my_putchar(i)
	}
	my_putchar('\n')
}

// func main() {
// 	my_print_digits()
// }
