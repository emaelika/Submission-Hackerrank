package main

import (
	"fmt"
)

func hurdleRace(k int32, height []int32) int32 {
	// Write your code here
	var max int32

	for _, val := range height {
		if val > max {
			max = val
		}
	}
	if max > k {
		return max - k
	}
	return 0

}
func main() {

	var height = []int32{1, 6, 3, 5, 2}
	var k int32 = 4
	fmt.Println(hurdleRace(k, height))

}
