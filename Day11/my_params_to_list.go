//
// EPITECH PROJECT, 2025
// my_params_to_list
// File description:
// day11
//

package main

import (
	"fmt"
	"os"
)

type ParamNode struct {
	Str  string
	Size int
	Next *ParamNode
}


type ListNode struct {
	Data string
	Next *ListNode
}

func my_params_to_list(ac int, av []string) *ListNode {
	var head *ListNode
	for i := 0; i < ac; i++ {
		node := &ListNode{
			Data: av[i],
			Next: head,
		}
		head = node
	}
	return head
}

// func main() {
// 	args := os.Args
// 	head := my_params_to_list(len(args), args)
// 	current := head
// 	for current != nil {
// 		fmt.Println(current.Data)
// 		current = current.Next
// 	}
// }