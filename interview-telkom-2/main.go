package main

import "fmt"

func levenshtein(str1, str2 []rune) int {
	s1len := len(str1)
	s2len := len(str2)
	column := make([]int, len(str1)+1)

	for y := 1; y <= s1len; y++ {
		column[y] = y
	}
	for x := 1; x <= s2len; x++ {
		column[0] = x
		lastkey := x - 1
		for y := 1; y <= s1len; y++ {
			oldkey := column[y]
			var incr int
			if str1[y-1] != str2[x-1] {
				incr = 1
			}

			column[y] = minimum(column[y]+1, column[y-1]+1, lastkey+incr)
			lastkey = oldkey
		}
	}
	return column[s1len]
}

func minimum(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

func main() {
	text1 := ""
	text2 := ""
	fmt.Print("Masukan text 1: ")
	fmt.Scan(&text1)
	fmt.Print("Masukan text 2: ")
	fmt.Scan(&text2)
	var str1 = []rune(text1)
	var str2 = []rune(text2)

	distance := levenshtein(str1, str2)
	if distance > 1 {
		fmt.Println("False")
	} else {
		fmt.Println("True")
	}
}
