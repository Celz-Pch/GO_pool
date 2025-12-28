//
// EPITECH PROJECT, 2025
// my_sort_list
// File description:
// day11
//

package main

type ListNode struct {
	Data string
	Next *ListNode
}

func my_sort_list(head *ListNode, cmp func(a, b string) int) *ListNode {
	if head == nil {
		return nil
	}
	sorted := false
	for !sorted {
		sorted = true
		current := head
		for current != nil && current.Next != nil {
			if cmp(current.Data, current.Next.Data) > 0 {
				current.Data, current.Next.Data = current.Next.Data, current.Data
				sorted = false
			}
			current = current.Next
		}
	}
	return head
}