package hackman

import (
	"fmt"
	"log"

	"github.com/PlayitLOUD73/hackman/api"

	rl "github.com/gen2brain/raylib-go/raylib"
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
// been guessed or not. It is used in
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
	alpha  [26]Letter
}

func (w *WordInfo) fullyGuessed() bool {
	for i := 0; i < len(w.list); i++ {
		if !w.list[i].guessed {
			return false
		}
	}
	return true
}

func (w *WordInfo) Setup(length int, key string) {

	w.word = api.GetWord(length, key)
	w.length = len(w.word)

	// for debugging and testing
	fmt.Printf("%s\n", w.word)

	// initializes w
	for i := 0; i < len(w.word); i++ {
		var temp WordLetter
		temp.abc = w.word[i]
		w.list = append(w.list, temp)
	}

	// initializes list of letters
	for i := 0; i < 26; i++ {
		for j := 0; j < len(w.word); j++ {
			if CharToNum(w.word[j]) == i {
				w.alpha[i].inWord = true
			}
		}
	}

	fmt.Print("\n")

}

func (w *WordInfo) inWord(c int32) bool {

	correct := false

	for i := 0; i < len(w.word); i++ {
		if int(c) == CharToNum(w.word[i]) {
			w.list[i].guessed = true
			correct = true
		}
	}
	return correct
}

// drawWord draws the letters guessed and lines underneath letter.
// No parameters are needed.
func (w *WordInfo) drawWord() {
	var i int32
	for i = 0; i < int32(w.length); i++ {

		rl.DrawRectangle(300+(25*i), 300, 20, 5, rl.Green)
		if w.list[i].guessed {
			drawLetter(w.list[i].abc, 300+(25*i), 265)
		}

	}

}

//var list [26]Letter

//var Ans WordInfo // answer

// CharToNum is a function to convert a letter
// from the alphabet (in byte form) to its
// corresponding position in the alphabet.
func CharToNum(c byte) int {
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

// inWord determines if a particular letter
// is in the word. The function returns true
// if the letter is in the word and false if
// it is not.

func uniToInt(c int32) int32 {
	return c - 65
}
