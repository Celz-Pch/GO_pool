//
// EPITECH PROJECT, 2025
// swap_endian_color
// File description:
// day09
//

package main

import "fmt"

func swap_endian_color(color int) int {
	red := (color >> 16) & 0xFF
	green := (color >> 8) & 0xFF
	blue := color & 0xFF
	return (blue << 16) | (green << 8) | red
}

// func main() {
// 	original_color := 0xFFAA33
// 	swapped_color := swap_endian_color(original_color)
// 	fmt.Printf("O color: 0x%06X\n", original_color)
// 	fmt.Printf("S color: 0x%06X\n", swapped_color)
// }