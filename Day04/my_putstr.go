//
// EPITECH PROJECT, 2025
// my_putstr
// File description:
// day04
//

package main

import "os"

func my_putstr(str string) int {
	count := 0

	for _, c := range str {
		os.Stdout.Write([]byte(string(c)))
		count++
	}
	return count
}

func main() {
	my_putstr("Hello world")
}
