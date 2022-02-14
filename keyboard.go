package hackman

import rl "github.com/gen2brain/raylib-go/raylib"

type Keyboard struct {
	keys [26]Button
}

func NewKeyboard() *Keyboard {
	k := new(Keyboard)
	start := 105
	for i := 0; i < 10; i++ {
		k.keys[i] = *NewKey(button2, start+(i*60), 350, rl.DarkGray, string(keysHelper(i)), keysHelper(i))
	}
	for i := 10; i < 19; i++ {
		k.keys[i] = *NewKey(button2, start+20+((i-10)*60), 420, rl.DarkGray, string(keysHelper(i)), keysHelper(i))
	}
	for i := 19; i < 26; i++ {
		k.keys[i] = *NewKey(button2, start+40+((i-19)*60), 490, rl.DarkGray, string(keysHelper(i)), keysHelper(i))
	}

	return k
}

func keysHelper(i int) byte {
	switch i {
	case 0:
		return 'q'
	case 1:
		return 'w'
	case 2:
		return 'e'
	case 3:
		return 'r'
	case 4:
		return 't'
	case 5:
		return 'y'
	case 6:
		return 'u'
	case 7:
		return 'i'
	case 8:
		return 'o'
	case 9:
		return 'p'
	case 10:
		return 'a'
	case 11:
		return 's'
	case 12:
		return 'd'
	case 13:
		return 'f'
	case 14:
		return 'g'
	case 15:
		return 'h'
	case 16:
		return 'j'
	case 17:
		return 'k'
	case 18:
		return 'l'
	case 19:
		return 'z'
	case 20:
		return 'x'
	case 21:
		return 'c'
	case 22:
		return 'v'
	case 23:
		return 'b'
	case 24:
		return 'n'
	case 25:
		return 'm'
	}
	return ' '
}

func keyConversion(i int) int {
	switch i {
	case 0:
		return 10
	case 1:
		return 23
	case 2:
		return 21
	case 3:
		return 12
	case 4:
		return 2
	case 5:
		return 13
	case 6:
		return 14
	case 7:
		return 15
	case 8:
		return 7
	case 9:
		return 16
	case 10:
		return 17
	case 11:
		return 18
	case 12:
		return 25
	case 13:
		return 24
	case 14:
		return 8
	case 15:
		return 9
	case 16:
		return 0
	case 17:
		return 3
	case 18:
		return 11
	case 19:
		return 4
	case 20:
		return 6
	case 21:
		return 22
	case 22:
		return 1
	case 23:
		return 20
	case 24:
		return 5
	case 25:
		return 19
	}
	return 0
}
