package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	lowerChars    = "abcdefghijklmnopqrstuvwxyz"
	upperChars    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numericChars  = "0123456789"
	specialChars  = "!@#$%^&*()-_=+[]{}|;:'\",.<>/?"
	ambiguousChars = "l1IO0"
	excludedChars = "\\|^"
)

func main() {
	var (
		length       int
		numPasswords int
		useLower     bool
		useUpper     bool
		useNumeric   bool
		useSpecial   bool
		avoidAmbiguous bool
	)

	flag.IntVar(&length, "l", 12, "Password length")
	flag.IntVar(&numPasswords, "n", 6, "Number of passwords to generate")
	flag.BoolVar(&useLower, "L", true, "Use lowercase characters")
	flag.BoolVar(&useUpper, "U", true, "Use uppercase characters")
	flag.BoolVar(&useNumeric, "N", true, "Use numeric characters")
	flag.BoolVar(&useSpecial, "S", false, "Use special characters")
	flag.BoolVar(&avoidAmbiguous, "H", false, "Avoid ambiguous characters")
	flag.Parse()

	if !(useLower || useUpper || useNumeric || useSpecial) {
		fmt.Println("Please select at least one character set.")
		return
	}

	rand.Seed(time.Now().UnixNano())

	characters := ""
	if useLower {
		characters += lowerChars
	}
	if useUpper {
		characters += upperChars
	}
	if useNumeric {
		characters += numericChars
	}
	if useSpecial {
		characters += specialChars
	}

	// Remove excluded characters
	characters = removeCharacters(characters, excludedChars)

	if avoidAmbiguous {
		// Remove ambiguous characters only once if avoidAmbiguous is set
		characters = removeCharacters(characters, ambiguousChars)
	}

	for i := 0; i < numPasswords; i++ {
		password := generatePassword(length, characters)
		fmt.Println(password)
	}
}

func generatePassword(length int, charset string) string {
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		password[i] = charset[randomIndex]
	}
	return string(password)
}

func removeCharacters(set string, charsToRemove string) string {
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune(charsToRemove, r) {
			return -1
		}
		return r
	}, set)
}
