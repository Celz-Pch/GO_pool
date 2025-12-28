//
// EPITECH PROJECT, 2025
// my_strdup
// File description:
// day08
//

package main

import "fmt"

func my_strdup(src string) string {
	dup := ""
	for _, c := range src {
		dup += string(c)
	}
	return dup
}

func main() {
	original := "Hello, World!"
	duplicate := my_strdup(original)
	fmt.Println("dup:", duplicate)
}