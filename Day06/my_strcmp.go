//
// EPITECH PROJECT, 2025
// my_strcmp
// File description:
// day06
//

package main

import "fmt"

func my_strlen(s1 string) int {
	count := 0

	for range s1 {
		count++
	}
	return count
}

func my_strcmp(s1 string, s2 string) int {
	i := 0
	len1 := my_strlen(s1)
	len2 := my_strlen(s2)

	for i < len1 && i < len2 && s1[i] == s2[i] {
		i++
	}
	if i == len1 && i == len2 {
		return 0
	}
	return int(s1[i]) - int(s2[i])
}

// func main() {
// 	s1 := "Hello"
// 	s2 := "Hellcacao"
// 	result := my_strcmp(s1, s2)
// 	fmt.Printf("Comparison result: %d\n", result)
// }
