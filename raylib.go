package main

import (
	"embed"
	_ "embed"
	"io/fs"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var background rl.Texture2D
var title rl.Texture2D
var button1 rl.Texture2D

//go:embed "assets/*.png"
var files embed.FS

// Init initializes a raylib window and sets the appropiate FPS target.
func Init() {
	rl.InitWindow(800, 600, "Hackman")
	rl.SetTargetFPS(60)

	//var Background []byte = assets[0].Name()

	assets, _ := fs.ReadDir(files, "assets")
	//var Background []byte
	//var name string = assets[0].Name()
	backgroundRaw, _ := fs.ReadFile(files, "assets/"+assets[0].Name())
	var backgroundImg *rl.Image = rl.LoadImageFromMemory(".png", backgroundRaw, int32(len(backgroundRaw)))
	background = rl.LoadTextureFromImage(backgroundImg)

	button1Raw, _ := fs.ReadFile(files, "assets/"+assets[1].Name())
	var button1Img *rl.Image = rl.LoadImageFromMemory(".png", button1Raw, int32(len(button1Raw)))
	button1 = rl.LoadTextureFromImage(button1Img)

	titleRaw, _ := fs.ReadFile(files, "assets/"+assets[2].Name())
	var titleImg *rl.Image = rl.LoadImageFromMemory(".png", titleRaw, int32(len(titleRaw)))
	title = rl.LoadTextureFromImage(titleImg)

}

// drawLetter draws an individual letter at the given position.
func drawLetter(c byte, x int32, y int32) {

	var fontsize int32 = 30
	rl.DrawText(string(c), x, y, fontsize, rl.Green)

}

func mouseClick(mouse int32, b []Button) {
	if rl.IsMouseButtonReleased(mouse) {
		clicked := clickedButtons(b)
		if clicked != -1 {
			b[clicked].onClick()
		}
	}
}

func DeInit() {

	rl.UnloadTexture(background)
	rl.UnloadTexture(title)
	rl.UnloadTexture(button1)
	rl.CloseWindow()
	os.Exit(0)
}
