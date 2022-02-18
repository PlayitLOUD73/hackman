package hackman

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Hitbox struct {
	xMax int
	xMin int
	yMax int
	yMin int
}

// HitboxLetter is used for on-screen keyboard
type HitboxLetter struct {
	dim Hitbox // dimensions
	abc byte   // letter
}

type Button struct {
	texture  rl.Texture2D
	isKey    bool
	key      byte
	newState GameState
	x        int
	y        int
	dim      Hitbox
	tint     rl.Color
	text     string
}

func NewButton(t rl.Texture2D, xPos int, yPos int, c rl.Color, s string, g GameState) *Button {

	b := new(Button)
	b.isKey = false
	b.texture = t
	b.x = xPos
	b.y = yPos
	b.tint = c
	b.text = s
	b.dim.xMax = xPos + int(t.Width)
	b.dim.xMin = xPos
	b.dim.yMax = yPos + int(t.Height)
	b.dim.yMin = yPos
	b.newState = g

	return b
}

func NewKey(t rl.Texture2D, xPos int, yPos int, c rl.Color, s string, k byte) *Button {

	b := new(Button)
	b.isKey = true
	b.texture = t
	b.x = xPos
	b.y = yPos
	b.tint = c
	b.text = s
	b.dim.xMax = xPos + int(t.Width)
	b.dim.xMin = xPos
	b.dim.yMax = yPos + int(t.Height)
	b.dim.yMin = yPos
	b.key = k

	return b
}

func (b *Button) drawButtonText() {
	font := 30
	width := rl.MeasureText(b.text, int32(font))
	x := b.dim.xMax - ((int(width) + (b.dim.xMax - b.x)) / 2)
	y := b.dim.yMax - ((font + (b.dim.yMax - b.y)) / 2)
	rl.DrawText(b.text, int32(x), int32(y), int32(font), rl.Black)
}

func (b *Button) changeState() GameState { return b.newState }

/*
func clickedLetters() byte {
	x := rl.GetMouseX()
	y := rl.GetMouseY()

	return 0
}
*/

func clickedButtons(b []Button) int {

	x := int(rl.GetMouseX())
	y := int(rl.GetMouseY())

	for i := 0; i < len(b); i++ {
		if x < b[i].dim.xMax && x > b[i].dim.xMin {
			if y < b[i].dim.yMax && y > b[i].dim.yMin {
				return i
			}
		}
	}

	return -1
}

func drawButtons(b []Button) {

	for i := 0; i < len(b); i++ {
		rl.DrawTexture(b[i].texture, int32(b[i].x), int32(b[i].y), b[i].tint)
		b[i].drawButtonText()
	}

}
