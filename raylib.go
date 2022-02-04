package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Init() {
	rl.InitWindow(800, 450, "Hackman")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Welcome to Hackman!", 0, 0, 20, rl.Black)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
