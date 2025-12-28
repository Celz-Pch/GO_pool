//
// EPITECH PROJECT, 2025
// my_apply_on_nodes
// File description:
// day11
//

package main

import "fmt"

type ListNode struct {
	Data string
	Next *ListNode
}

func my_apply_on_nodes(head *ListNode, f func(string)) {
	current := head
	for current != nil {
		f(current.Data)
		current = current.Next
	}
}

// func printData(data string) {
// 	fmt.Println(data)
// }

// func main() {
// 	head := &ListNode{Data: "first"}
// 	head.Next = &ListNode{Data: "second"}
// 	head.Next.Next = &ListNode{Data: "third"}

// 	my_apply_on_nodes(head, printData)
// }