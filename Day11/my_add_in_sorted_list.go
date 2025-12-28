//
// EPITECH PROJECT, 2025
// my_add_in_sorted_list
// File description:
// day11
//

package main

type ListNode struct {
	Data string
	Next *ListNode
}

func my_add_in_sorted_list(head **ListNode, new_node *ListNode, cmp func(a, b string) int) {
	if head == nil || new_node == nil {
		return
	}
	if *head == nil || cmp(new_node.Data, (*head).Data) < 0 {
		new_node.Next = *head
		*head = new_node
		return
	}
	current := *head
	for current.Next != nil && cmp(current.Next.Data, new_node.Data) < 0 {
		current = current.Next
	}
	new_node.Next = current.Next
	current.Next = new_node
}
