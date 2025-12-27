//
// EPITECH PROJECT, 2025
// my_print_comb2
// File description:
// day03
//

package main

import "os"

func my_putchar(c byte) {
	os.Stdout.Write([]byte{c})
}

func putchar_all(a, b byte) {
    my_putchar((a / 10) + '0');
    my_putchar((a % 10) + '0');
    my_putchar(' ');
    my_putchar((b / 10) + '0');
    my_putchar((b % 10) + '0');
    if !(a == 98 && b == 99) {
        my_putchar(',');
        my_putchar(' ');
    }
}

func my_print_comb2() {
	for a := 0; a <= 98; a++ {
		for b := a + 1; b <= 99; b++ {
			putchar_all(byte(a), byte(b))
		}
	}
	my_putchar('\n')
}

// func main() {
// 	my_print_comb2()
// }
