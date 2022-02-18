package hackman

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Hitbox is a struct used to define a region for a button.
type Hitbox struct {
	xMax int
	xMin int
	yMax int
	yMin int
}

// Button is a struct that is used for both buttons and for keys for the virtual keyboard.
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

// NewButton is a constructor to create a button and initialize it.
// This is not used for keys.
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

// NewKey is a constructor to create a button specificially for the on-screen keyboard.
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

// drawButtonText draws centered text on a button.
func (b *Button) drawButtonText() {
	font := 30
	width := rl.MeasureText(b.text, int32(font))
	x := b.dim.xMax - ((int(width) + (b.dim.xMax - b.x)) / 2)
	y := b.dim.yMax - ((font + (b.dim.yMax - b.y)) / 2)
	rl.DrawText(b.text, int32(x), int32(y), int32(font), rl.Black)
}

// changeState is a helper function to change gamestates based
// on a button's function.
func (b *Button) changeState() GameState { return b.newState }

// clickedButtons is a function to find if a button is clicked
// when given an array of buttons.
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

// drawButtons is a function to draw all buttons in a button array.
func drawButtons(b []Button) {

	for i := 0; i < len(b); i++ {
		rl.DrawTexture(b[i].texture, int32(b[i].x), int32(b[i].y), b[i].tint)
		b[i].drawButtonText()
	}

}
