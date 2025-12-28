//
// EPITECH PROJECT, 2025
// my_find_node
// File description:
// day11
//

package main

type ListNode struct {
	Data string
	Next *ListNode
}

func my_find_node(head *ListNode, to_find string) *ListNode {
	current := head
	for current != nil {
		if current.Data == to_find {
			return current
		}
		current = current.Next
	}
	return nil
}
