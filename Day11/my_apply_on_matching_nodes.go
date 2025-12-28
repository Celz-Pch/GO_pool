//
// EPITECH PROJECT, 2025
// my_apply_on_matching_nodes
// File description:
// day11
//

package main

import "fmt"

type ListNode struct {
	Data string
	Next *ListNode
}

func my_apply_on_matching_nodes(head *ListNode, f func(string), cond func(string) bool) {
	current := head
	for current != nil {
		if cond(current.Data) {
			f(current.Data)
		}
		current = current.Next
	}
}

// func is_even_length(data string) bool {
// 	return len(data)%2 == 0
// }

// func print_data(data string) {
// 	fmt.Println(data)
// }

// func main() {
// 	head := &ListNode{Data: "first"}
// 	head.Next = &ListNode{Data: "second"}
// 	head.Next.Next = &ListNode{Data: "third"}
// 	head.Next.Next.Next = &ListNode{Data: "fourth"}

// 	my_apply_on_matching_nodes(head, print_data, is_even_length)
// }