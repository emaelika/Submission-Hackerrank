package main

import "fmt"

func saveThePrisoner(n int32, m int32, s int32) int32 {
	// Write your code here
	var x int32

	if m < n {
		if (s-1)+m > n {
			x = s + m - 1 - n
			return x
		}
		x = s + m - 1
		return x
	}
	if m%n == 0 {

		if s-1 < 1 {
			return s
		}
		return s - 1
	}
	x = (m % n) + s - 1
	return x
}
func main() {

	fmt.Println("begin")
	fmt.Println(saveThePrisoner(4, 6, 2))
	fmt.Println(saveThePrisoner(4, 12, 2))

	fmt.Println(saveThePrisoner(999999999, 999999999, 1)) //23525398
	fmt.Println(saveThePrisoner(4, 13, 2))

}
