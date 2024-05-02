package main

import (
	"fmt"
)

func hiking(steps int32, path string) int32 {
	var output int32 = 0
	if steps < 0 {
		fmt.Println("err")
	}
	state := 0
	for i, val := range path {

		if val == rune('D') {
			state--
		}
		if val == rune('U') {
			state++
		}
		if i > 0 {
			if path[i] == byte('U') && state == 0 {
				output += 1
			}
		}

	}
	return output
}
func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	// make leader board

	var ranking []int32
	var p int32 = 0
	for i, val := range ranked {
		if i == 0 {
			p++
			ranking = append(ranking, p)

		}
		if i != 0 {
			if val == ranked[i-1] {
				ranking = append(ranking, p)
			} else {
				p++
				ranking = append(ranking, p)
			}

		}
	}
	// fmt.Println(ranking)
	var hasil []int32

	for _, element := range player {
		for y, value := range ranked {
			if element >= value {
				hasil = append(hasil, ranking[y])
				break
			}
			z := len(ranked) - 1

			if y == z {
				a := ranking[y] + 1
				hasil = append(hasil, a)
				fmt.Println(hasil)
			}

		}

	}
	return hasil

}
func main() {
	fmt.Println(hiking(4, "DDUUDUDUDUDUDUUUD"))

	mapper := make(map[int]string)
	var angka = []int{1, 2, 3, 4}
	var huruf = []string{"A", "B", "C", "D"}
	for i, val := range angka {
		mapper[val] = huruf[i]
	}
	var board = []int32{100, 100, 50, 40, 40, 20, 10}
	var player = []int32{5, 25, 50, 120}
	fmt.Println(climbingLeaderboard(board, player))

}
