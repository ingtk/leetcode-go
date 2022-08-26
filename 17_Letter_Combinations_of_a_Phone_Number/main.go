package main

import (
	"fmt"
	"strconv"
)

var mapLetters = map[int]string{
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	return makeCombination(digits, 0)
}

func makeCombination(digits string, cnt int) []string {
	results := make([]string, 0, 81)
	digit := pint(string(digits[cnt]))
	letters := mapLetters[digit]
	if cnt == len(digits)-1 {
		for _, l := range letters {
			results = append(results, string(l))
		}

		return results
	}

	for _, l := range letters {
		for _, c := range makeCombination(digits, cnt+1) {
			results = append(results, string(l)+c)
		}
	}

	return results
}

func pint(s string) int {
	n, _ := strconv.ParseInt(s, 10, 64)
	return int(n)
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations("234"))
	fmt.Println(letterCombinations("2345"))
}
