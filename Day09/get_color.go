//
// EPITECH PROJECT, 2025
// get_color
// File description:
// day09
//

package main

import "fmt"

func get_color(red, green, blue uint8) int {
	return int(red)<<16 | int(green)<<8 | int(blue)
}

// func main() {
// 	red := uint8(255)
// 	green := uint8(165)
// 	blue := uint8(0)

// 	color := get_color(red, green, blue)
// 	fmt.Printf("Color (RGB: %d, %d, %d) = 0x%06X\n", red, green, blue, color)
// }
