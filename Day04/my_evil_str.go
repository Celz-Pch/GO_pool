//
// EPITECH PROJECT, 2025
// my_evil_str
// File description:
// day04
//

package main

import "os"

func my_strlen(str string) int {
	count := 0

	for range str {
		count ++;
	}
	return count
}

func my_evil_str(str string) {
	for i := my_strlen(str) - 1; i >= 0; i-- {
		os.Stdout.Write([]byte{str[i]})
	}
	os.Stdout.Write([]byte{'\n'})
}

// func main() {
// 	s := "Roh Roh"
// 	my_evil_str(s)
// }