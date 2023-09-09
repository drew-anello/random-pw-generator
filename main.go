package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	length := flag.Int("length", 12, "the length of the password")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	letters := make([]rune, 52)
	for i := 0; i < 26; i++ {
		letters[i] = rune('a' + i)
	}

	for i := 0; i < 26; i++ {
		letters[26+i] = rune('A' + i)
	}

	specialChars := "!@#$%^&*()-_=+[]{}|;:'\",.<>?/"

	numbers := make([]rune, 10)

	for i := 0; i < 10; i++ {
		numbers[i] = rune('0' + i)
	}

	password := make([]rune, *length)

	password[0] = letters[rand.Intn(len(letters))]
	password[1] = numbers[rand.Intn(len(numbers))]
	password[2] = rune(specialChars[rand.Intn(len(specialChars))])

	specialCount := 1

	allChars := append(append(letters, numbers...), []rune(specialChars)...)

	for i := 3; i < 12; i++ {
		chosenChar := allChars[rand.Intn(len(allChars))]

		isSpecial := false
		for _, char := range specialChars {
			if chosenChar == char {
				isSpecial = true
				break
			}
		}

		if isSpecial && specialCount >= 2 {
			i--
			continue
		}
		password[i] = chosenChar

		if isSpecial {
			specialCount++
		}
	}

	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	fmt.Println("Generated password:", string(password))
}
