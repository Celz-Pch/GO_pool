//
// EPITECH PROJECT, 2025
// my_advanced_sort_word_array
// File description:
// day10
//

package	main

func my_advanced_sort_word_array(array []string, size int, f func(a, b string) int) {
	tmp := ""

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if f(array[i], array[j]) > 0 {
				tmp = array[j]
				array[j] = array[i]
				array[i] = tmp
			}
		}
	}
}

// func main() {
// 	words := []string{"banana", "apple", "orange", "grape", "kiwi"}
// 	compare := func(a, b string) int {
// 		if a < b {
// 			return -1
// 		} else if a > b {
// 			return 1
// 		}
// 		return 0
// 	}
// 	my_advanced_sort_word_array(words, len(words), compare)
// 	for _, word := range words {
// 		fmt.Println(word)
// 	}
// }