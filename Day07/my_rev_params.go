//
// EPITECH PROJECT, 2025
// my_rev_params
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

func my_rev_params(args []string) {
	for i := len(args) - 1; i != 0; i-- {
		my_putstr(args[i])
		my_putchar('\n')
	}
}

// func main() {
// 	args := os.Args[0:]
// 	my_rev_params(args)
// }