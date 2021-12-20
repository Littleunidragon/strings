package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ToLower(s string) string {
	s = strings.ToUpper(s)
	return s
}

func ToUpper(s string) string {
	s = strings.ToLower(s)
	return s
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Printf("%v\n", CountRune(text, c))
}
