package main

import (
	"fmt"
)

func viralAdvertising(n int32) int32 {
	// Write your code here
	// x = n; day 1: n/2=2, 2*3=6 (ini untuk besoknya
	// day 0) kumulatif 2|x = 5/2 + (i)*3
	x := recursiveCount(5, n-1, 0)
	return x

}
func recursiveCount(base int32, n int32, cumulative int32) int32 {
	// 5=> 5?2/2
	if n == 0 {

		return cumulative + base/2
	}
	cumulative += base / 2
	n--
	base2 := base / 2 * 3
	fmt.Println(base, base2)
	return recursiveCount(base2, n, cumulative)
}

func main() {
	fmt.Println("begin")
	var n int32 = 5
	fmt.Println(viralAdvertising(n))

}
