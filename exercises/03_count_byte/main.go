package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountByte(s string, c byte) int {
	var r int
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			r++
		}
	}
	return r
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Print("Now byte: ")
	var c byte
	fmt.Scan(&c)
	fmt.Printf("%v\n", CountByte(text, c))
}
