//
// EPITECH PROJECT, 2025
// do-op
// File description:
// day10
//

package main

import (
	"fmt"
	"os"
	"strconv"
)

func do_op(nb1 int, nb2 int, operator string) int {
	switch operator {
	case "+":
		return nb1 + nb2
	case "-":
		return nb1 - nb2
	case "*":
		return nb1 * nb2
	case "/":
		if nb2 == 0 {
			fmt.Println("divise by 0")
			return 0
		}
		return nb1 / nb2
	case "%":
		if nb2 == 0 {
			fmt.Println("modulo by 0")
			return 0
		}
		return nb1 % nb2
	default:
		fmt.Println("invalid operator")
		return 0
	}
}

// func main() {
// 	if len(os.Args) != 4 {
// 		return
// 	}

// 	num1, err1 := strconv.Atoi(os.Args[1])
// 	operator := os.Args[2]
// 	num2, err2 := strconv.Atoi(os.Args[3])

// 	if err1 != nil || err2 != nil {
// 		return
// 	}

// 	result := do_op(num1, num2, operator)
// 	fmt.Println(result)
// }
