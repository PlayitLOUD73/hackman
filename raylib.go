package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var background rl.Texture2D
var title rl.Texture2D
var button1 rl.Texture2D

// Init initializes a raylib window and sets the appropiate FPS target.
func Init() {
	rl.InitWindow(800, 600, "Hackman")
	rl.SetTargetFPS(60)
	background = rl.LoadTexture("assets/background.png")
	title = rl.LoadTexture("assets/title.png")
	button1 = rl.LoadTexture("assets/button1.png")
}

// drawLetter draws an individual letter at the given position.
func drawLetter(c byte, x int32, y int32) {

	var fontsize int32 = 30
	rl.DrawText(string(c), x, y, fontsize, rl.Blue)

}

// drawWord draws the letters guessed and lines underneath letter.
// No parameters are needed.
func drawWord() {
	var i int32
	for i = 0; i < int32(Ans.length); i++ {

		rl.DrawRectangle(200+(25*i), 225, 20, 5, rl.Black)
		if Ans.list[i].guessed {
			drawLetter(Ans.list[i].abc, 200+(25*i), 200)
		}

	}

}

func DeInit() {

	rl.UnloadTexture(background)
	rl.UnloadTexture(title)
	rl.UnloadTexture(button1)
	rl.CloseWindow()
}
