package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println("== Find Word in Sentence ==")
	var (
		s, k string
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Sentence: ")
	s, _ = reader.ReadString('\n')

	fmt.Print("Enter keyword: ")
	k, _ = reader.ReadString('\n')

	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	k = strings.TrimSpace(k)
	k = strings.ToLower(k)

	if s == "" || k == "" {
		fmt.Println("Invalid input, please enter a valid sentence and keyword")
		return
	}

	splitted := strings.Split(s, " ")
	found := slices.Contains(splitted, k)

	if found {
		fmt.Println("keyword is found in the sentence")
	} else {
		fmt.Println("keyword is not found in the sentence")
	}
}
