package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GameLoop() {
	var c int32
	numGuesses := 5

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Welcome to Hackman!", 0, 0, 20, rl.Black)
		rl.DrawText("The word is "+Ans.word, 0, 50, 20, rl.Black)

		drawWord()
		c = rl.GetKeyPressed()
		if c != 0 {
			if inWord(uniToInt(c)) {
			} else {
				numGuesses--
			}
		}
		rl.DrawText("Lives: "+strconv.Itoa(numGuesses), 400, 400, 20, rl.Black)
		rl.EndDrawing()
	}

	rl.CloseWindow()

}
