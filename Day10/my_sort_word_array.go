//
// EPITECH PROJECT, 2025
// my_sort_word_array
// File description:
// day10
//

package main

func my_sort_word_array(array []string, size int) {
	tmp := ""

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if array[i] > array[j] {
				tmp = array[j]
				array[j] = array[i]
				array[i] = tmp
			}
		}
	}
}

// func main() {
// 	words := []string{"banana", "apple", "orange", "grape", "kiwi"}
// 	my_sort_word_array(words, len(words))
// 	for _, word := range words {
// 		fmt.Println(word)
// 	}
// }