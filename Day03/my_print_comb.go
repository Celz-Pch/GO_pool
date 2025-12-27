//
// EPITECH PROJECT, 2025
// my_print_comb
// File description:
// day03
//

package main

import "os"

func my_putchar(c byte) {
	os.Stdout.Write([]byte{c})
}

func my_display_all(a, b, c byte) {
	my_putchar(a)
	my_putchar(b)
	my_putchar(c)
	if !(a == 55 && b == 56 && c == 57) {
		my_putchar(',')
		my_putchar(' ')
	} else {
		my_putchar('\n')
	}
}

func loop_c(a, b byte) {
	for c := b + 1; c <= '9'; c++ {
		my_display_all(a, b, c)
	}
}

func loop_b(a byte) {
	for b := a + 1; b <= '8'; b++ {
		loop_c(a, b)
	}
}

func loop_a() {
	for a := byte('0'); a <= '7'; a++ {
		loop_b(a)
	}
}

func my_print_comb() {
	loop_a()
}

// func main() {
// 	my_print_comb()
// }
