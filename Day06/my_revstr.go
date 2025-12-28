//
// EPITECH PROJECT, 2025
// my_revstr
// File description:
// day06
//

package main

import "fmt"

func my_revstr(str string) string {
    runes := []rune(str)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

// func main() {
// 	s := "Hello, World!"
// 	rev := my_revstr(s)
// 	fmt.Println(rev)
// }