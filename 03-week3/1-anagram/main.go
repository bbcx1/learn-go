package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Println(checkAnagram("lemon", "melon"))       // true
	fmt.Println(checkAnagram("listen", "silent"))     // true
	fmt.Println(checkAnagram("triangle", "integral")) // true
	fmt.Println(checkAnagram("hello", "world"))       // false
	fmt.Println(checkAnagram("", "   "))              // false
}

func checkAnagram(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	var group1 map[string]int = map[string]int{}
	var group2 map[string]int = map[string]int{}

	for i := 0; i < len(word1); i++ {
		group1[string(word1[i])]++
		group2[string(word2[i])]++
	}

	return maps.Equal(group1, group2)
}
