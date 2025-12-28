//
// EPITECH PROJECT, 2025
// my_show_word_array
// File description:
// day08
//

package main

import "os"

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

func my_show_word_array(arr []string) {
	for _, word := range arr {
		my_putstr(word)
		my_putchar('\n')
	}
}

func main() {
	words := []string{"Hello", "World", "This", "Is", "Go"}
	my_show_word_array(words)
}