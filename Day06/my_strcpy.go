//
// EPITECH PROJECT, 2025
// my_strcpy
// File description:
// day06
//

package main

import "fmt"

func my_strcpy(dest []byte, src string) {
    for i := 0; i < len(src); i++ {
        dest[i] = src[i]
    }
}

// func main() {
// 	src := "Hello, World!"
// 	dest := make([]byte, len(src))
// 	my_strcpy(dest, src)
// 	fmt.Println(string(dest))
// }
