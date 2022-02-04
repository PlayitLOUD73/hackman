package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Init() {
	rl.InitWindow(800, 450, "Hackman")
	rl.SetTargetFPS(60)
}
