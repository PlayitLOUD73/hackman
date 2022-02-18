package hackman

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// GameState is an enumerator for determining which section
// the game is in. This is used for managing sections of
// game.
type GameState int

const (
	MAINMENU GameState = iota // MAINMENU = 0
	GAME                      // GAME = 1
	GAMEOVER                  // GAMEOVER = 2
	VICTORY                   // VICTORY = 3
	NEWGAME                   // NEWGAME = 4
	EXIT                      // EXIT = 5
)

// GameController is a struct that interfaces with different
// sections of the game. It also holds variables that must be
// changed between frames.
type GameController struct {
	state    GameState
	ans      WordInfo
	guesses  int
	scroll   int
	keyboard Keyboard
}

// NewGameController is a constructor to create a new GameController.
func NewGameController() *GameController {
	g := new(GameController)
	return g
}

// game is a function that contains the main gameplay.
func (g *GameController) game() {
	buttons := g.keyboard.keys[:]

	var darkGreen rl.Color = rl.Color{0x11, 0x28, 0x21, 0xFF}

	c := rl.GetKeyPressed()

	// check for keyboard presses
	if c != 0 {
		buttonMarker := keyConversion(int(uniToInt(c)))
		if g.ans.inWord(uniToInt(c), buttons[buttonMarker]) {
			buttons[buttonMarker].tint = rl.Green
		} else {
			buttons[buttonMarker].tint = rl.Red
			g.guesses--
			rl.PlaySound(errorSound)
		}
	}

	// check to see if state needs to change
	if g.guesses <= 0 {
		g.state = GAMEOVER
	}
	if g.ans.fullyGuessed() {
		g.state = VICTORY
	}

	// check for mouse clicks
	g, buttons = mouseClick(rl.MouseLeftButton, buttons, g)

	// Drawing
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)
	rl.DrawTexture(background, 0, 0, rl.RayWhite)
	rl.DrawRectangle(75, 100, 650, 450, darkGreen)
	drawButtons(buttons)

	g.ans.drawWord()

	rl.DrawText("Lives: "+strconv.Itoa(g.guesses), 575, 100, 30, rl.Green)
	rl.EndDrawing()
}

// gameOver is a function that shows the game over screen and allows
// the player to play again.
func (g *GameController) gameOver() {

	// button initialization
	var buttons []Button
	buttons = append(buttons, *NewButton(button1, (rl.GetScreenWidth()/2-200)/2, ((rl.GetScreenHeight() + 50) / 2), rl.Red, "Exit", EXIT))
	buttons = append(buttons, *NewButton(button1, (rl.GetScreenWidth()/2+200)/2, ((rl.GetScreenHeight() + 50) / 2), rl.Green, "New Game", NEWGAME))

	// check for mouse clicks
	g, buttons = mouseClick(rl.MouseLeftButton, buttons, g)

	// Drawing
	rl.BeginDrawing()

	rl.DrawTexture(background, 0, 0, rl.DarkBrown)
	rl.DrawText("Gameover! The answer was "+g.ans.word, 120, 120, 25, rl.Red)
	drawButtons(buttons)
	rl.EndDrawing()

}

// mainMenu is a function that is the main menu. It is the initial
// screen the player sees.
func (g *GameController) mainMenu() {

	// Button initialization
	var buttons []Button
	buttons = append(buttons, *NewButton(button1, (rl.GetScreenWidth()-200)/2, (rl.GetScreenHeight()+50)/2, rl.SkyBlue, "Start", NEWGAME))

	speed := 2 // horizontal scroll speed

	// check for mouse clicks
	g, buttons = mouseClick(rl.MouseLeftButton, buttons, g)

	// Drawing
	rl.BeginDrawing()

	rl.DrawTexture(background, int32(g.scroll/speed), 0, rl.RayWhite)
	rl.DrawTexture(background, int32((g.scroll/speed - 800)), 0, rl.RayWhite)

	rl.DrawTexture(title, int32((rl.GetScreenWidth()-650)/2), int32((rl.GetScreenHeight()-400)/2), rl.RayWhite)

	drawButtons(buttons)

	rl.EndDrawing()

	// scroll management
	g.scroll++
	if g.scroll > int(background.Width)*speed {
		g.scroll = 0
	}
}

// vict is a function to display the victory screen, with an
// option to restart the game or exit.
func (g *GameController) vict() {

	// Button initialization
	var buttons []Button
	buttons = append(buttons, *NewButton(button1, (rl.GetScreenWidth()/2-200)/2, ((rl.GetScreenHeight() + 50) / 2), rl.Red, "Exit", EXIT))
	buttons = append(buttons, *NewButton(button1, (rl.GetScreenWidth()/2+200)/2, ((rl.GetScreenHeight() + 50) / 2), rl.Green, "New Game", NEWGAME))

	// check for mouse clicks
	g, buttons = mouseClick(rl.MouseLeftButton, buttons, g)

	// Drawing
	rl.BeginDrawing()
	rl.DrawTexture(background, 0, 0, rl.RayWhite)
	drawButtons(buttons)
	rl.DrawText("You guessed the word!", 120, 120, 30, rl.Green)
	g.ans.drawWord()
	rl.EndDrawing()
}

// newGame is a function to reinitialize for another round of hangman.
func (g *GameController) newGame(key string) {
	var Ans WordInfo
	Ans.Setup(-1, key)
	k := NewKeyboard()
	g.guesses = 6
	g.ans = Ans
	g.state = GAME
	g.keyboard = *k

}
