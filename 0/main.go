package main

import (
	"fmt"
	"strings"
)

func bracketCheck(teks string) bool {
	//checking bracket value
	var brackets []string
	for _, val := range teks {
		x := fmt.Sprintf("%c", val)
		if x == "(" || x == "[" || x == "{" {
			brackets = append(brackets, x)
		}
		if x == ")" || x == "]" || x == "}" {
			// cek bracket buka
			pass := cekBracketBuka(brackets, x)
			if !pass {
				return false
			}
			brackets = append(brackets, x)
		}
	}

	//zero value handling
	if len(brackets) == 0 {
		return false
	}

	//counting the brackets
	mapper := make(map[string]int)
	for _, value := range brackets {
		_, ok := mapper[value]
		if !ok {
			mapper[value] = 1
		} else {
			mapper[value]++
		}

	}

	//bracket frequency comparison
	if mapper["("] != mapper[")"] {
		return false
	}
	if mapper["{"] != mapper["}"] {
		return false
	}
	if mapper["["] != mapper["]"] {
		return false
	}
	return true
}

func cekBracketBuka(arr []string, value string) bool {
	if len(arr) == 0 {
		return false
	}
	switch value {
	case "}":
		for _, val := range arr {
			if val == "{" {
				return true

			}
		}
	case ")":
		for _, val := range arr {
			if val == "(" {
				return true

			}
		}
	case "]":
		for _, val := range arr {
			if val == "[" {
				return true

			}
		}
	}
	return false
}

func reoccurenceCheck(array []string) string {
	arr1 := strings.Split(array[0], ", ")
	arr2 := strings.Split(array[1], ", ")
	var arr3 []string
	for _, val := range arr1 {
		for _, str := range arr2 {
			if val == str {
				arr3 = append(arr3, val)
				break
			}
		}
	}
	if len(arr3) == 0 {
		return "false"
	}
	result := strings.Join(arr3, ", ")

	return result

}
func main() {
	fmt.Println("Nomor 1")
	fmt.Println(bracketCheck("(hello (world))"))                 //true
	fmt.Println(bracketCheck("{}}"))                             //false
	fmt.Println(bracketCheck("{}({}"))                           //false
	fmt.Println(bracketCheck("(hello {world}) [(test) {case}]")) //true
	fmt.Println(bracketCheck("(hello {world)"))                  //false
	fmt.Println(bracketCheck("(hello} {world)"))
	fmt.Println("Nomor 2")
	fmt.Println(reoccurenceCheck([]string{"1, 3, 4, 7, 13", "1, 2, 4, 13, 15"}))
	fmt.Println(reoccurenceCheck([]string{"1, 3, 4, 7, 13", "2, 5, 11, 15"}))
	fmt.Println("Nomor Bonus")
	fmt.Println(cekPalindrome("katak"))   //true
	fmt.Println(cekPalindrome("jakarta")) //false
	fmt.Println(cekPalindrome("abccba"))  //true

}
func cekPalindrome(kata string) bool {
	var eja []string
	for _, val := range kata {
		x := fmt.Sprintf("%c", val)
		eja = append(eja, x)
	}
	var balik []string
	y := len(eja)
	y--

	for i := y; i >= 0; i-- {
		balik = append(balik, eja[i])

	}
	terbalik := strings.Join(balik, "")
	ok := strings.Compare(terbalik, kata)

	if ok == 0 {
		return true
	}
	return false
}
