package main

import (
	"fmt"
	"log"
)

// Letter is letter information for
// a list of letters. It is meant to
// be used in an array to determine what
// letter it is.
type Letter struct {
	inWord  bool
	guessed bool
}

// WordLetter is a letter and if it has
// been guessed or not. It is uesd in
// WordInfo.
type WordLetter struct {
	abc     byte // letter
	guessed bool
}

// WordInfo is information regarding a word
// to be used for guessing. It includes the
// word in string form, list of letters, and
// word length.
type WordInfo struct {
	word   string
	list   []WordLetter
	length int
}

var list [26]Letter
var Ans WordInfo // answer

// charToNum is a function to convert a letter
// from the alphabet (in byte form) to its
// corresponding position in the alphabet.
func charToNum(c byte) int {
	switch c {

	case 'a':
		return 0
	case 'b':
		return 1
	case 'c':
		return 2
	case 'd':
		return 3
	case 'e':
		return 4
	case 'f':
		return 5
	case 'g':
		return 6
	case 'h':
		return 7
	case 'i':
		return 8
	case 'j':
		return 9
	case 'k':
		return 10
	case 'l':
		return 11
	case 'm':
		return 12
	case 'n':
		return 13
	case 'o':
		return 14
	case 'p':
		return 15
	case 'q':
		return 16
	case 'r':
		return 17
	case 's':
		return 18
	case 't':
		return 19
	case 'u':
		return 20
	case 'v':
		return 21
	case 'w':
		return 22
	case 'x':
		return 23
	case 'y':
		return 24
	case 'z':
		return 25
	default:
		log.Fatalln("charToNum must get a letter")
	}
	return 0
}

// Setup initializes the data necessary
// for the game to run. It takes in a
// length to determine the size of the word.
func Setup(length int) {

	Ans.word = GetWord(length)
	Ans.length = len(Ans.word)
	fmt.Printf("%s\n", Ans.word)

	// initializes Ans
	for i := 0; i < len(Ans.word); i++ {
		var temp WordLetter
		temp.abc = Ans.word[i]
		Ans.list = append(Ans.list, temp)
	}

	// initializes list of letters
	for i := 0; i < 26; i++ {
		for j := 0; j < len(Ans.word); j++ {
			if charToNum(Ans.word[j]) == i {
				list[i].inWord = true
				fmt.Printf("%c ", Ans.word[j])
			}
		}
	}

	fmt.Print("\n")

}

// inWord determines if a particular letter
// is in the word. The function returns true
// if the letter is in the word and false if
// it is not.
func inWord(c byte) bool {

	for i := 0; i < len(Ans.word); i++ {
		if c == Ans.word[i] {
			Ans.list[i].guessed = true
			return true
		}
	}
	return false
}
