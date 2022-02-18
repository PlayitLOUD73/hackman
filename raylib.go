package hackman

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
var button2 rl.Texture2D

//go:embed "assets/*.png"
var files embed.FS

// Init initializes a raylib window and sets the appropiate FPS target.
// It also initializes textures
func Init() {
	rl.InitWindow(800, 600, "Hackman")
	rl.SetTargetFPS(60)

	assets, _ := fs.ReadDir(files, "assets")

	backgroundRaw, _ := fs.ReadFile(files, "assets/"+assets[0].Name())
	var backgroundImg *rl.Image = rl.LoadImageFromMemory(".png", backgroundRaw, int32(len(backgroundRaw)))
	background = rl.LoadTextureFromImage(backgroundImg)

	button1Raw, _ := fs.ReadFile(files, "assets/"+assets[1].Name())
	var button1Img *rl.Image = rl.LoadImageFromMemory(".png", button1Raw, int32(len(button1Raw)))
	button1 = rl.LoadTextureFromImage(button1Img)

	button2Raw, _ := fs.ReadFile(files, "assets/"+assets[2].Name())
	var button2Img *rl.Image = rl.LoadImageFromMemory(".png", button2Raw, int32(len(button2Raw)))
	button2 = rl.LoadTextureFromImage(button2Img)

	titleRaw, _ := fs.ReadFile(files, "assets/"+assets[3].Name())
	var titleImg *rl.Image = rl.LoadImageFromMemory(".png", titleRaw, int32(len(titleRaw)))
	title = rl.LoadTextureFromImage(titleImg)

}

// drawLetter draws an individual letter at the given position.
func drawLetter(c byte, x int32, y int32) {

	var fontsize int32 = 30
	rl.DrawText(string(c), x, y, fontsize, rl.Green)

}

func mouseClick(mouse int32, b []Button, g *GameController) (*GameController, []Button) {
	if rl.IsMouseButtonReleased(mouse) {
		clicked := clickedButtons(b)
		if clicked != -1 {
			if b[clicked].isKey {
				if g.ans.inWord(int32(CharToNum(b[clicked].key)), b[clicked]) {
					b[clicked].tint = rl.Green
				} else {
					b[clicked].tint = rl.Red
					g.guesses--
				}
			} else {
				g.state = b[clicked].changeState()
			}
		}
	}
	return g, b
}

func DeInit() {

	rl.UnloadTexture(background)
	rl.UnloadTexture(title)
	rl.UnloadTexture(button1)
	rl.CloseWindow()
	os.Exit(0)
}
