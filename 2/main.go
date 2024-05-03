package main

import (
	"fmt"
)

func designerPdfViewer(h []int32, word string) int32 {
	// Write your code here
	// parsing rune
	wordInt := []int{}
	for _, val := range word {
		val -= 97
		wordInt = append(wordInt, int(val))
	}
	letterHeights := []int32{}
	for _, char := range wordInt {
		letterHeights = append(letterHeights, h[char])
	}
	var max int32
	for _, char := range letterHeights {
		if char > max {
			max = char
		}
	}

	var length int32 = int32(len(wordInt))

	return max * length
}

func main() {
	fmt.Println("begin")
	var h = []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}
	word := "az"
	fmt.Println(designerPdfViewer(h, word))

}
