//
// EPITECH PROJECT, 2025
// my_str_islower
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

func my_str_islower(str string) int {
	for i := 0; i < my_strlen(str); i++ {
		if str[i] < 'a' || str[i] > 'z' {
			return int(0)
		}
	}
	return int(1)
}

// func main() {
// 	s := "helloworld"
// 	lower := my_str_islower(s)
// 	fmt.Println(lower)
// }