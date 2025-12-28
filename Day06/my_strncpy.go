//
// EPITECH PROJECT, 2025
// my_strncpy
// File description:
// day06
//

package main

import "fmt"

func my_strncpy(dest []byte, src string, n int) {
    for i := 0; i < n; i++ {
        dest[i] = src[i]
    }
}

// func main() {
// 	src := "Hello, World!"
// 	dest := make([]byte, len(src))
// 	my_strncpy(dest, src, 5)
// 	fmt.Println(string(dest))
// }
