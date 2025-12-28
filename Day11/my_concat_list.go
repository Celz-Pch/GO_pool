//
// EPITECH PROJECT, 2025
// my_concat_list
// File description:
// day11
//

package	main

type ListNode struct {
	Data string
	Next *ListNode
}

func my_concat_list(first *ListNode, second *ListNode) *ListNode {
	if first == nil {
		return second
	}
	current := first
	for current.Next != nil {
		current = current.Next
	}
	current.Next = second
	return first
}
