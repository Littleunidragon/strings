package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountRune(s string, c rune) int {
	var r int
	for i := 0; i < len(s); i++ {
		if rune(s[i]) == c {
			r++
		}
	}
	return r
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Print("Now rune: ")
	var c rune
	fmt.Scan(&c)
	fmt.Printf("%v\n", CountRune(text, c))
}
