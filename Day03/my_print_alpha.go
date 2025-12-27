//
// EPITECH PROJECT, 2025
// day03
// File description:
// my_print_alpha
//

package main

import "os"

func my_putchar(c byte) {
	os.Stdout.Write([]byte{c})
}

func my_print_alpha() {
   for i := byte(97); i <= byte(122); i++ {
	   my_putchar(i)
   }
   my_putchar('\n')
}

// func main() {
// 	my_print_alpha()
// }