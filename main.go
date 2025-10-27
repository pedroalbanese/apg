package main

import (
	"crypto/hmac"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"strings"

	"golang.org/x/crypto/sha3"
)

// Conjuntos de caracteres
const (
	lowerChars     = "abcdefghijklmnopqrstuvwxyz"
	upperChars     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numericChars   = "0123456789"
	specialChars   = "!@#$%&*()-_=+[]{};:,.<>/?"
	ambiguousChars = "l1IO0"
)

// Flags globais
var (
	useLower       bool
	useUpper       bool
	useNumeric     bool
	useSpecial     bool
	avoidAmbiguous bool
	spell          bool
	seed           string
)

// ----- Mecanismo determinístico seguro -----
// HMAC-DRBG-like gerador determinístico a partir de uma seed arbitrária
type deterministicPRF struct {
	seed    []byte
	counter uint64
	buffer  []byte
	pos     int
}

func newDeterministicPRF(seed []byte) *deterministicPRF {
	return &deterministicPRF{seed: seed, counter: 1}
}

func (d *deterministicPRF) refill() {
	h := hmac.New(func() hash.Hash { return sha3.New512() }, d.seed)
	var ctr [8]byte
	for i := 0; i < 8; i++ {
		ctr[7-i] = byte((d.counter >> (8 * i)) & 0xFF)
	}
	h.Write(ctr[:])
	d.buffer = h.Sum(nil)
	d.pos = 0
	d.counter++
}

func (d *deterministicPRF) Read(p []byte) (n int, err error) {
	off := 0
	for off < len(p) {
		if d.pos >= len(d.buffer) || len(d.buffer) == 0 {
			d.refill()
		}
		toCopy := len(d.buffer) - d.pos
		if toCopy > len(p)-off {
			toCopy = len(p) - off
		}
		copy(p[off:off+toCopy], d.buffer[d.pos:d.pos+toCopy])
		d.pos += toCopy
		off += toCopy
	}
	return len(p), nil
}

// Gera índice sem viés (rejection sampling)
func secureIndex(r interface{}, charsetLen int) int {
	max := 256 - (256 % charsetLen)
	buf := make([]byte, 1)
	for {
		switch src := r.(type) {
		case *deterministicPRF:
			src.Read(buf)
		default: // crypto/rand
			rand.Read(buf)
		}
		if int(buf[0]) < max {
			return int(buf[0]) % charsetLen
		}
	}
}

// ----- Função principal -----
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
	flag.StringVar(&seed, "seed", "", "Optional hex seed for deterministic password generation")
	flag.Parse()

	if !(useLower || useUpper || useNumeric || useSpecial) {
		fmt.Println("Please select at least one character set.")
		return
	}

	// Monta o conjunto de caracteres
	charset := ""
	if useLower {
		charset += lowerChars
	}
	if useUpper {
		charset += upperChars
	}
	if useNumeric {
		charset += numericChars
	}
	if useSpecial {
		charset += specialChars
	}
	if avoidAmbiguous {
		charset = removeCharacters(charset, ambiguousChars)
	}

	// Define fonte de aleatoriedade
	var reader interface{}
	if seed != "" {
		seedBytes, err := hex.DecodeString(seed)
		if err != nil {
			fmt.Println("Invalid seed format. Use hex (e.g., 64 hex chars for 32 bytes).")
			return
		}
		reader = newDeterministicPRF(seedBytes)
	} else {
		reader = rand.Reader
	}

	// Gera senhas
	for i := 0; i < numPasswords; i++ {
		password := generatePassword(length, charset, reader)
		if spell {
			fmt.Printf("%s %s\n", password, spellPassword(password))
		} else {
			fmt.Println(password)
		}
	}
}

// ----- Geração de senha -----
func generatePassword(length int, charset string, reader interface{}) string {
	out := make([]byte, length)
	for i := 0; i < length; i++ {
		out[i] = charset[secureIndex(reader, len(charset))]
	}
	return string(out)
}

// ----- Utilitários -----
func removeCharacters(set string, charsToRemove string) string {
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune(charsToRemove, r) {
			return -1
		}
		return r
	}, set)
}

func spellPassword(password string) string {
	phonetic := map[rune]string{
		'a': "alfa", 'A': "Alpha", 'b': "bravo", 'B': "Bravo",
		'c': "charlie", 'C': "Charlie", 'd': "delta", 'D': "Delta",
		'e': "echo", 'E': "Echo", 'f': "foxtrot", 'F': "Foxtrot",
		'g': "golf", 'G': "Golf", 'h': "hotel", 'H': "Hotel",
		'i': "india", 'I': "India", 'j': "juliett", 'J': "Juliett",
		'k': "kilo", 'K': "Kilo", 'l': "lima", 'L': "Lima",
		'm': "mike", 'M': "Mike", 'n': "november", 'N': "November",
		'o': "oscar", 'O': "Oscar", 'p': "papa", 'P': "Papa",
		'q': "quebec", 'Q': "Quebec", 'r': "romeo", 'R': "Romeo",
		's': "sierra", 'S': "Sierra", 't': "tango", 'T': "Tango",
		'u': "uniform", 'U': "Uniform", 'v': "victor", 'V': "Victor",
		'w': "whiskey", 'W': "Whiskey", 'x': "x-ray", 'X': "X-ray",
		'y': "yankee", 'Y': "Yankee", 'z': "zulu", 'Z': "Zulu",
		'0': "ZERO", '1': "ONE", '2': "TWO", '3': "THREE",
		'4': "FOUR", '5': "FIVE", '6': "SIX", '7': "SEVEN",
		'8': "EIGHT", '9': "NINE", '!': "EXCLAMATION", '@': "AT",
		'#': "HASH", '$': "DOLLAR", '%': "PERCENT", '^': "CARET",
		'&': "AMPERSAND", '*': "ASTERISK", '(': "LEFT_PARENTHESIS",
		')': "RIGHT_PARENTHESIS", '-': "HYPHEN", '_': "UNDERSCORE",
		'=': "EQUAL", '+': "PLUS", '[': "LEFT_BRACKET", ']': "RIGHT_BRACKET",
		'{': "LEFT_CURLY_BRACE", '}': "RIGHT_CURLY_BRACE", '|': "PIPE",
		';': "SEMICOLON", ':': "COLON", ',': "COMMA", '.': "PERIOD",
		'<': "LESS_THAN", '>': "GREATER_THAN", '/': "SLASH",
		'?': "QUESTION_MARK",
	}
	parts := []string{}
	for _, c := range password {
		if p, ok := phonetic[c]; ok {
			parts = append(parts, p)
		} else {
			parts = append(parts, string(c))
		}
	}
	return strings.Join(parts, "-")
}
