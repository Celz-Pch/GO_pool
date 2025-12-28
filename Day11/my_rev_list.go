//
// EPITECH PROJECT, 2025
// my_rev_list
// File description:
// day11
//

package main

type ListNode struct {
	Data string
	Next *ListNode
}

func my_rev_list(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		nextNode := current.Next
		current.Next = prev
		prev = current
		current = nextNode
	}
	return prev
}

// func main() {
// 	head := &ListNode{Data: "first"}
// 	head.Next = &ListNode{Data: "second"}
// 	head.Next.Next = &ListNode{Data: "third"}

// 	newHead := my_rev_list(head)
// 	current := newHead
// 	for current != nil {
// 		fmt.Println(current.Data)
// 		current = current.Next
// 	}
// }