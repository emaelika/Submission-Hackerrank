package main

import (
	"fmt"
	"strconv"
	"strings"
)

func beautifulDays(i int32, j int32, k int32) int32 {
	// Write your code here
	var counter int32
	for c := i; c <= j; c++ {
		var selisih int32
		if c > pembalik(c) {
			selisih = c - pembalik(c)
		} else if pembalik(c) > c {
			selisih = pembalik(c) - c
		} else {
			selisih = 0
		}
		if selisih%k == 0 {
			counter++
		}
	}
	return counter
}
func pembalik(n int32) int32 {
	base := strconv.Itoa(int(n))
	process := strings.Split(base, "")
	x := 0
	for i, val := range process {
		num, _ := strconv.Atoi(val)
		x += (num * (pangkat(10, i)))

	}
	return int32(x)
}
func pangkat(base int, pangkat int) int {
	if pangkat == 0 {
		return 1
	}
	n := 1
	for i := 0; i < pangkat; i++ {
		n *= base
	}
	return n
}

func main() {
	fmt.Println("begin")
	var x int32 = 123
	fmt.Println(pembalik(x))
	// var i,j,k int32 = 5,6,7
	// fmt.Println(beautifulDays(i,j,k))

}
