//
// EPITECH PROJECT, 2025
// my_strlowcase
// File description:
// day06
//

package main

import "fmt"

func my_strlen(str []byte) int {
	count := 0

	for range str {
		count ++;
	}
	return count
}

func my_strlowcase(str []byte) string {
	for i := 0; i < my_strlen(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			str[i] = str[i] + 32
		}
	}
	return string(str)
}

func main() {
	s := []byte("Hello, World!")
	upper := my_strlowcase(s)
	fmt.Println(upper)
}
