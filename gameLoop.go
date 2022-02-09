package hackman

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameState int

const (
	MAINMENU GameState = iota // MAINMENU = 0
	GAME                      // GAME = 1
	GAMEOVER                  // GAMEOVER = 2
	VICTORY                   // VICTORY = 3
	NEWGAME                   // NEWGAME = 4
	EXIT                      // EXIT = 5
)

func GameLoop(key string) {

	var c int32
	var state GameState = MAINMENU

	Ans, numGuesses := newGame(key)

	//rl.BeginDrawing()
	for !rl.WindowShouldClose() {
		rl.ClearBackground(rl.RayWhite)
		// Which part of the game am I in?
		switch state {

		case 0:
			state = mainMenu()
		case 1:
			Ans, numGuesses, state = game(Ans, c, numGuesses)
		case 2:
			state = gameOver(Ans)
		case 3:
			state = vict(Ans)
		case 4:
			Ans, numGuesses = newGame(key)
			state = GAME
		case 5:
			return
		}

		c = rl.GetKeyPressed()

	}
}
