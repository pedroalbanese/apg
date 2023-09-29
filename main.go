package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

const (
	lowerChars     = "abcdefghijklmnopqrstuvwxyz"
	upperChars     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numericChars   = "0123456789"
	specialChars   = "!@#$%&*()-_=+[]{};:,.<>/?"
	ambiguousChars = "l1IO0"
)

var (
	useLower       bool
	useUpper       bool
	useNumeric     bool
	useSpecial     bool
	avoidAmbiguous bool
	spell          bool
)

func main() {
	var (
		length       int
		numPasswords int
	)

	flag.IntVar(&length, "l", 12, "Password length")
	flag.IntVar(&numPasswords, "n", 6, "Number of passwords to generate")
	flag.BoolVar(&useLower, "L", true, "Use lowercase characters")
	flag.BoolVar(&useUpper, "U", true, "Use uppercase characters")
	flag.BoolVar(&useNumeric, "N", true, "Use numeric characters")
	flag.BoolVar(&useSpecial, "S", false, "Use special characters")
	flag.BoolVar(&avoidAmbiguous, "H", false, "Avoid ambiguous characters")
	flag.BoolVar(&spell, "spell", false, "Spell passwords using phonetic alphabet")
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

	if avoidAmbiguous {
		// Remove ambiguous characters only once if avoidAmbiguous is set
		characters = removeCharacters(characters, ambiguousChars)
	}

	for i := 0; i < numPasswords; i++ {
		password := generatePassword(length, characters)
		if spell {
			fmt.Printf("%s ", password)
			password = spellPassword(password)
		}
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

func spellPassword(password string) string {
	phoneticAlphabet := map[rune]string{
		'a': "alfa", 'A': "Alpha", 'b': "bravo", 'B': "Bravo", 'c': "charlie", 'C': "Charlie",
		'd': "delta", 'D': "Delta", 'e': "echo", 'E': "Echo", 'f': "foxtrot", 'F': "Foxtrot",
		'g': "golf", 'G': "Golf", 'h': "hotel", 'H': "Hotel", 'i': "india", 'I': "India",
		'j': "juliett", 'J': "Juliett", 'k': "kilo", 'K': "Kilo", 'l': "lima", 'L': "Lima",
		'm': "mike", 'M': "Mike", 'n': "november", 'N': "November", 'o': "oscar", 'O': "Oscar",
		'p': "papa", 'P': "Papa", 'q': "quebec", 'Q': "Quebec", 'r': "romeo", 'R': "Romeo",
		's': "sierra", 'S': "Sierra", 't': "tango", 'T': "Tango", 'u': "uniform", 'U': "Uniform",
		'v': "victor", 'V': "Victor", 'w': "whiskey", 'W': "Whiskey", 'x': "x-ray", 'X': "X-ray",
		'y': "yankee", 'Y': "Yankee", 'z': "zulu", 'Z': "Zulu",
		'0': "ZERO", '1': "ONE", '2': "TWO", '3': "THREE", '4': "FOUR", '5': "FIVE",
		'6': "SIX", '7': "SEVEN", '8': "EIGHT", '9': "NINE",
		'!': "EXCLAMATION", '@': "AT", '#': "HASH", '$': "DOLLAR", '%': "PERCENT",
		'^': "CARET", '&': "AMPERSAND", '*': "ASTERISK", '(': "LEFT_PARENTHESIS",
		')': "RIGHT_PARENTHESIS", '-': "HYPHEN", '_': "UNDERSCORE", '=': "EQUAL",
		'+': "PLUS", '[': "LEFT_BRACKET", ']': "RIGHT_BRACKET", '{': "LEFT_CURLY_BRACE",
		'}': "RIGHT_CURLY_BRACE", '|': "PIPE", ';': "SEMICOLON", ':': "COLON", ',': "COMMA",
		'.': "PERIOD", '<': "LESS_THAN", '>': "GREATER_THAN", '/': "SLASH", '?': "QUESTION_MARK",
	}

	splittedPassword := strings.Split(password, "")
	var spelledPassword []string

	for i, char := range splittedPassword {
		charRune := rune(char[0])
		spelledChar, found := phoneticAlphabet[charRune]
		if !found {
			spelledChar = string(charRune)
		}

		if i == 0 && unicode.IsUpper(charRune) {
			spelledChar = strings.Title(spelledChar)
		}

		spelledPassword = append(spelledPassword, spelledChar)
	}

	return strings.Join(spelledPassword, "-")
}

func removeCharacters(set string, charsToRemove string) string {
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune(charsToRemove, r) {
			return -1
		}
		return r
	}, set)
}
