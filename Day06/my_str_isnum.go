//
// EPITECH PROJECT, 2025
// my_str_isnum
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

func my_str_isnum(str string) int {
	for i := 0; i < my_strlen(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			return int(0)
		}
	}
	return int(1)
}

// func main() {
// 	s := "123456"
// 	num := my_str_isnum(s)
// 	fmt.Println(num)
// }