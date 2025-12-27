//
// EPITECH PROJECT, 2025
// my_print_realpha
// File description:
// day03
//

package main

import "os"

func my_putchar(c byte) {
	os.Stdout.Write([]byte{c})
}

func my_print_revalpha() {
	for i := byte(122); i >= byte(97); i-- {
		my_putchar(i)
	}
	my_putchar('\n')
}

// func main() {
// 	my_print_revalpha()
// }
