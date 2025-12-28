//
// EPITECH PROJECT, 2025
// my_strupcase
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

func my_strupcase(str []byte) string {
	for i := 0; i < my_strlen(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			str[i] = str[i] - 32
		}
	}
	return string(str)
}

// func main() {
// 	s := []byte("Hello, World!")
// 	upper := my_strupcase(s)
// 	fmt.Println(upper)
// }