package hackman

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MAINMENU GameState = iota // MAINMENU = 0
	GAME                      // GAME = 1
	GAMEOVER                  // GAMEOVER = 2
	VICTORY                   // VICTORY = 3
	NEWGAME                   // NEWGAME = 4
	EXIT                      // EXIT = 5
)

type GameController struct {
	state   GameState
	ans     WordInfo
	guesses int
	scroll  int
}

func NewGameController(key string) *GameController {
	g := new(GameController)
	g.newGame(key)
	return g
}

func (g *GameController) game(c int32) {

	var darkGreen rl.Color = rl.Color{0x11, 0x28, 0x21, 0xFF}
	if c != 0 {
		if g.ans.inWord(uniToInt(c)) {
		} else {
			g.guesses--
		}
	}

	if g.guesses <= 0 {
		g.state = GAMEOVER
	}

	if g.ans.fullyGuessed() {
		g.state = VICTORY
	}

	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)
	rl.DrawTexture(background, 0, 0, rl.RayWhite)
	rl.DrawRectangle(100, 100, 600, 400, darkGreen)

	g.ans.drawWord()

	rl.DrawText("Lives: "+strconv.Itoa(g.guesses), 575, 100, 30, rl.Green)
	rl.EndDrawing()
}

func (g *GameController) gameOver() {

	var buttons []Button
	buttons = append(buttons, *NewButton(button1, (rl.GetScreenWidth()/2-200)/2, ((rl.GetScreenHeight() + 50) / 2), rl.Red, "Exit", EXIT))
	buttons = append(buttons, *NewButton(button1, (rl.GetScreenWidth()/2+200)/2, ((rl.GetScreenHeight() + 50) / 2), rl.Green, "New Game", NEWGAME))

	g.state = mouseClick(rl.MouseLeftButton, buttons, g.state)

	rl.BeginDrawing()

	rl.DrawTexture(background, 0, 0, rl.DarkBrown)
	rl.DrawText("Gameover! The answer was "+g.ans.word, 120, 120, 25, rl.Red)
	drawButtons(buttons)
	rl.EndDrawing()

}

func (g *GameController) mainMenu() {

	var buttons []Button
	buttons = append(buttons, *NewButton(button1, (rl.GetScreenWidth()-200)/2, (rl.GetScreenHeight()+50)/2, rl.SkyBlue, "Start", NEWGAME))
	speed := 2
	//rl.SetTextureWrap(title, rl.WrapRepeat)
	rl.BeginDrawing()
	//rl.ClearBackground(rl.RayWhite)
	rl.DrawTexture(background, int32(g.scroll/speed), 0, rl.RayWhite)
	rl.DrawTexture(background, int32((g.scroll/speed - 800)), 0, rl.RayWhite)

	rl.DrawTexture(title, int32((rl.GetScreenWidth()-650)/2), int32((rl.GetScreenHeight()-400)/2), rl.RayWhite)

	drawButtons(buttons)

	rl.EndDrawing()

	g.state = mouseClick(rl.MouseLeftButton, buttons, g.state)
	g.scroll++
	if g.scroll > int(background.Width)*speed {
		g.scroll = 0
	}
}

func (g *GameController) vict() {

	var buttons []Button
	buttons = append(buttons, *NewButton(button1, (rl.GetScreenWidth()/2-200)/2, ((rl.GetScreenHeight() + 50) / 2), rl.Red, "Exit", EXIT))
	buttons = append(buttons, *NewButton(button1, (rl.GetScreenWidth()/2+200)/2, ((rl.GetScreenHeight() + 50) / 2), rl.Green, "New Game", NEWGAME))

	g.state = mouseClick(rl.MouseLeftButton, buttons, g.state)

	rl.BeginDrawing()
	rl.DrawTexture(background, 0, 0, rl.RayWhite)
	drawButtons(buttons)
	rl.DrawText("You guessed the word!", 120, 120, 30, rl.Green)
	g.ans.drawWord()
	rl.EndDrawing()
}

func (g *GameController) newGame(key string) {
	var Ans WordInfo
	Ans.Setup(-1, key)
	g.guesses = 6
	g.ans = Ans
	g.state = GAME
}
