//
// EPITECH PROJECT, 2025
// my_str_isprintable
// File description:
// day06
//

package main

import "fmt"

func my_strlen(str string) int {
	count := 0

	for range str {
		count++
	}
	return count
}

func my_str_isprintable(str string) int {
	for i := 0; i < my_strlen(str); i++ {
		if str[i] < 32 || str[i] > 126 {
			return int(0)
		}
	}
	return int(1)
}

// func main() {
// 	s := "Hello, World!\n"
// 	printable := my_str_isprintable(s)
// 	fmt.Println(printable)
// }