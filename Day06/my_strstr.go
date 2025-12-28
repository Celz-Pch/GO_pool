//
// EPITECH PROJECT, 2025
// my_strstr
// File description:
// daty06
//

package main

import "fmt"

func my_strlen(str string) int {
	count := 0

	for range str {
		count ++;
	}
	return count
}

func my_strstr(str string, to_find string) string {
	str_len := my_strlen(str)
	find_len := my_strlen(to_find)

	if my_strlen(to_find) == 0 {
		return str
	}
	for i := 0; i <= str_len-find_len; i++ {
		match := true
		for j := 0; j < find_len; j++ {
			if str[i+j] != to_find[j] {
				match = false
				break
			}
		}
		if match {
			return str[i:]
		}
	}
	return ""
}

// func main() {
// 	s := "Hello, World!"
// 	to_find := "World"
// 	result := my_strstr(s, to_find)
// 	fmt.Println(result)
// }
