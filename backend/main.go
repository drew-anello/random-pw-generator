package main

import (
	"fmt"
	"strings"
)
func main()  {
	letters := make([]rune, 52)

	for i := 0; i < 26; i++ {
		letters[i] = rune('a' + i)
	}

	for i := 0; i < 26; i++ {
		letters[26 + i] = rune('A' + i)
	}

	numbers := make([]rune, 10)

	for i := 0; i < 10; i++ {
		numbers[i] = rune('0' + i)
	}

	specialChars := "!@#$%^&*()-_=+[]{}|;:'\",.<>?/"

	fmt.Println( strings.Split(string(specialChars), ""))
	fmt.Println( strings.Split(string(letters), ""))
	fmt.Println( strings.Split(string(numbers), ""))
}