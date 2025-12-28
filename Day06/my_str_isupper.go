//
// EPITECH PROJECT, 2025
// my_str_isupper
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

func my_str_isupper(str string) int {
	for i := 0; i < my_strlen(str); i++ {
		if str[i] < 'A' || str[i] > 'Z' {
			return int(0)
		}
	}
	return int(1)
}

// func main() {
// 	s := "HELLOWORLD"
// 	upper := my_str_isupper(s)
// 	fmt.Println(upper)
// }