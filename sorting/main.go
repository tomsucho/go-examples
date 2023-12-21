package main

import (
	"fmt"
	"sort"
)

type byLen []string

func (b byLen) Len() int {
	return len(b)
}

func (b byLen) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byLen) Less(i, j int) bool {
	return len(b[i]) < len(b[j])
}

func main() {

	test := []string{"apple", "joe", "bananas"}
	sort.Sort(byLen(test))
	fmt.Println(test)

	s := "abcdefg🔥"
	txt := []rune(s)
	reversed := ""

	// for idx, c := range s {
	// 	txt[idx] = c
	// }

	for i := len(txt) - 1; i >= 0; i-- {
		// fmt.Printf("Rune %d: %s\n", i, string(txt[i]))
		reversed = reversed + string(txt[i])
	}
	fmt.Println(s)
	fmt.Println(reversed)

	// Sorting
	nums := []int{9, 5, 7, 1, 4}
	sort.Ints(nums)
	fmt.Println(nums)
}
