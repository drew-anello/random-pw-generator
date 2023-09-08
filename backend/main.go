package main

import (
	"fmt"
	"strings"
)
func main()  {
	letters := make([]rune, 26)

	for i := 0; i < 26; i++ {
		letters[i] = rune('a' + i)
	}
	fmt.Println("letters", strings.Split(string(letters), ""))
}