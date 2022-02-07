package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GameLoop() {
	var c int32
	numGuesses := 5

	//rl.BeginDrawing()
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

		if numGuesses <= 0 {
			if !gameOver() {
				rl.CloseWindow()
			}
		}
	}

	rl.CloseWindow()

}

func gameOver() bool {

	var option int32 = 0

	for option == 0 {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Gameover! The answer was "+Ans.word, 100, 200, 30, rl.Red)
		rl.DrawText("Play again? y/n", 200, 300, 30, rl.Blue)
		rl.EndDrawing()
		c := rl.GetKeyPressed()
		//c = uniToInt(c)
		// is c y or n?
		if c == 89 || c == 78 {
			option = c
		}

	}

	return !(option == 78)

}
