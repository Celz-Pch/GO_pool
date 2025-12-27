//
// EPITECH PROJECT, 2025
// my_getnbr
// File description:
// day04
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

func my_getnbr(str string) int {
	result := 0
	negatif := 1
	i := 0

	if my_strlen(str) > 0 && str[0] == '-' {
		negatif = -1
		i++
	}
	for ; i < my_strlen(str); i++ {
		result += int(str[i] - 48)
		if i != my_strlen(str) - 1 {
			result *= 10
		}
	}
	return result * negatif
}

// func main() {
// 	n := my_getnbr("-1234")
// 	fmt.Println(n)
// }
