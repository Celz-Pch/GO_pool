//
// EPITECH PROJECT, 2025
// my_put_nbr
// File description:
// day03
//

package main

import "os"

func my_putchar(c byte) {
	os.Stdout.Write([]byte{c})
}

func my_put_nbr(nb int) {
	if nb < 0 {
		my_putchar('-')
		nb = -nb
	}
	if nb >= 10 {
		my_put_nbr(nb / 10)
	}
	my_putchar(byte('0' + (nb % 10)))
}

// func main() {
// 	my_put_nbr(123456)
// 	my_putchar('\n')
//     my_put_nbr(-2147483648)
// }
