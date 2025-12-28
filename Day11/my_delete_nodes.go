//
// EPITECH PROJECT, 2025
// my_delete_nodes
// File description:
// day11
//

package main

type ListNode struct {
	Data string
	Next *ListNode
}

func my_delete_nodes(head **ListNode, data_ref string) {
	current := *head
	var prev *ListNode

	for current != nil {
		if current.Data == data_ref {
			if prev == nil {
				*head = current.Next
			} else {
				prev.Next = current.Next
			}
			current = current.Next
		} else {
			prev = current
			current = current.Next
		}
	}
}
