//
// EPITECH PROJECT, 2025
// my_show_param_array
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


func my_show_param_array(par []InfoParam) int {
	for i := 0; i < len(par); i++ {
		if par[i].Str == "" {
			break
		}
		fmt.Println(par[i].Str)
		fmt.Println(par[i].Length)
		for _, word := range par[i].WordArray {
			fmt.Println(word)
		}
	}
	return 0
}

func main() {
	args := []string{"Hello world", "This is Go", "TEST"}
	paramsArray := my_params_to_array(len(args), args)
	my_show_param_array(paramsArray)
}