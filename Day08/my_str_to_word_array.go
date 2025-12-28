//
// EPITECH PROJECT, 2025
// my_str_to_word_array
// File description:
// day08
//

package main

import "fmt"

func my_str_to_word_array(str string) []string {
	var words []string
	start := -1
	for i, c := range str {
		if c != ' ' && c != '\t' && c != '\n' && start == -1 {
			start = i
		}
		if (c == ' ' || c == '\t' || c == '\n') && start != -1 {
			words = append(words, str[start:i])
			start = -1
		}
	}
	if start != -1 {
		words = append(words, str[start:])
	}
	return words
}

func main() {
	input := "Hello world this is Go"
	words := my_str_to_word_array(input)
	for _, word := range words {
		fmt.Println(word)
	}
}