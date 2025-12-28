//
// EPITECH PROJECT, 2025
// my_print_params
// File description:
// day07
//

package main

import (
	"os"
)

func my_putchar(c byte) {
	os.Stdout.Write([]byte{c})
}

func my_putstr(str string) int {
	count := 0

	for _, c := range str {
		os.Stdout.Write([]byte(string(c)))
		count++
	}
	return count
}

func my_print_params(args []string) {
	for i := 0; i < len(args); i++ {
		my_putstr(args[i])
		my_putchar('\n')
	}
}

// func main() {
// 	args := os.Args[0:]
// 	my_print_params(args)
// }