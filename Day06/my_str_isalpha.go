//
// EPITECH PROJECT, 2025
// my_str_isalpha
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

func my_str_isalpha(str string) int {
	for i := 0; i < my_strlen(str); i++ {
		if (str[i] < 'A' || (str[i] > 'Z' && str[i] < 'a') || str[i] > 'z') {
			return int(0)
		}
	}
	return int(1)
}

// func main() {
// 	s := "HelloWorld"
// 	alpha := my_str_isalpha(s)
// 	fmt.Println(alpha)
// }