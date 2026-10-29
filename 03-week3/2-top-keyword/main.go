package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

func main() {
	fmt.Println(getTopKeyword("vibe coding is a coding while you are vibing along with coding")) // coding
	fmt.Println(getTopKeyword("go makes programming fun fun"))                                   // fun
	fmt.Println(getTopKeyword("go go go"))                                                       // go
	fmt.Println(getTopKeyword("cats chase dogs and dogs chase dogs"))                            // dogs
	fmt.Println(getTopKeyword("the quick brown fox"))                                            // brown
	fmt.Println(getTopKeyword(""))                                                               // ""
}

// getTopKeyword retrieve most frequent words in a sentence
func getTopKeyword(sentence string) string {
	var words []string = strings.Split(sentence, " ")
	var group map[string]int = map[string]int{}

	for _, w := range words {
		group[w]++
	}

	skeys := slices.Collect(maps.Keys(group))
	slices.Sort(skeys)

	var topKeyword string
	var highestFreq int = 0

	for _, key := range skeys {
		freq := group[key]
		if freq > highestFreq {
			highestFreq = freq
			topKeyword = key
		}
	}

	return topKeyword
}
