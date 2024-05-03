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

func countOperations(num1 int, num2 int) int {
	if num2 == 0 || num1 == 0 {
		return 0
	}
	if num1 > num2 {
		_, _, count := recurMin(num1, num2, 0)
		return count
	}
	_, _, count := recurMin(num2, num1, 0)

	return count
}

func recurMin(hi int, lo int, count int) (int, int, int) {
	if hi%lo == 0 {
		count += (hi / lo)

		return hi, lo, count
	}
	x := hi % lo
	count += (hi / lo)

	return recurMin(lo, x, count)
}

func main() {

	var height = []int32{1, 6, 3, 5, 2}
	fmt.Println(height)

	// var k int32 = 4
	// fmt.Println(hurdleRace(k, height))
	// fmt.Println(countOperations(10, 10))
	for i := 1; i < len(height); i++ {
		if height[i] < height[i-1] {
			height[i], height[i-1] = height[i-1], height[i]
		}
	}
	fmt.Println(height)
}
