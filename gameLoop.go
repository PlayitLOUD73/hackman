package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func GameLoop() {

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Welcome to Hackman!", 0, 0, 20, rl.Black)
		rl.DrawText("The word is "+Ans.word, 0, 50, 20, rl.Black)

		rl.EndDrawing()
	}

	rl.CloseWindow()

}
