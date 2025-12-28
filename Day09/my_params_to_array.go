//
// EPITECH PROJECT, 2025
// my_params_to_array
// File description:
// day09
//

package main

import "fmt"

type InfoParam struct {
	Length    int
	Str       string
	Copy      string
	WordArray []string
}

func my_str_to_word_array(str string) []string {
	var words []string
	start := -1
	for i, c := range str {
		if c != ' ' && c != '\t' && c != '\n' && start == -1 {
			start = i
		}
		if (c == ' ' || c == '\t' || c == '\n') && start != -1 {
			words = append(words, str[start:i])
			start = -1
		}
	}
	if start != -1 {
		words = append(words, str[start:])
	}
	return words
}


func my_params_to_array(ac int, av []string) []InfoParam {
	params := make([]InfoParam, ac+1)
	for i := 0; i < ac; i++ {
		params[i].Length = len(av[i])
		params[i].Str = av[i]
		params[i].Copy = av[i]
		params[i].WordArray = my_str_to_word_array(av[i])
	}
	params[ac].Str = ""
	return params
}

// func main() {
// 	args := []string{"Hello world", "This is Go", "TEST"}
// 	paramsArray := my_params_to_array(len(args), args)

// 	for _, param := range paramsArray {
// 		if param.Str == "" {
// 			break
// 		}
// 		fmt.Printf("Length: %d\n", param.Length)
// 		fmt.Printf("Str: %s\n", param.Str)
// 		fmt.Printf("Copy: %s\n", param.Copy)
// 		fmt.Printf("WordArray: %v\n", param.WordArray)
// 	}
// }