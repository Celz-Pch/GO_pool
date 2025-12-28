//
// EPITECH PROJECT, 2025
// my_strncmp
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

func my_strncmp(s1 string, s2 string, n int) int {
	i := 0
	len := n

	for i < len && s1[i] == s2[i] {
		i++
	}
	if i == len {
		return 0
	}
	return int(s1[i]) - int(s2[i])
}

// func main() {
// 	s1 := "Hello"
// 	s2 := "Hellcacao"
// 	result := my_strncmp(s1, s2, 4)
// 	fmt.Printf("Comparison result: %d\n", result)
// }
