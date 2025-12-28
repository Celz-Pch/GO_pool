//
// EPITECH PROJECT, 2025
// concat_params
// File description:
// day08
//

package main

import (
	"fmt"
	"os"
)

func concat_params(args []string) string {
	result := ""
	for i, arg := range args {
		result += arg
		if i < len(args)-1 {
			result += " "
		}
	}
	return result
}

func main() {
	args := os.Args[1:]
	concatenated := concat_params(args)
	fmt.Print(concatenated)
}