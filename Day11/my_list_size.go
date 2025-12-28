//
// EPITECH PROJECT, 2025
// my_list_size
// File description:
// day11
//

package main

type ListNode struct {
	Data string
	Next *ListNode
}

func my_list_size(head *ListNode) int {
	count := 0
	current := head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

// func main() {
// 	head := &ListNode{Data: "first"}
// 	head.Next = &ListNode{Data: "second"}
// 	head.Next.Next = &ListNode{Data: "third"}

// 	size := my_list_size(head)
// 	fmt.Println("Size of the list:", size) // Output: Size of the list: 3
// }