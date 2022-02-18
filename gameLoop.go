package hackman

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// GameLoop is the main function for the game. It controls each
// scene in the game and initializes a GameController to assist it.
func GameLoop(key string) {

	var controller GameController = *NewGameController()
	controller.state = MAINMENU

	rl.PlayMusicStream(flightpath)

	for !rl.WindowShouldClose() {
		rl.ClearBackground(rl.RayWhite)

		rl.UpdateMusicStream(flightpath)

		// Which part of the game am I in?
		switch controller.state {
		case MAINMENU:
			controller.mainMenu()
		case GAME:
			controller.game()
		case GAMEOVER:
			controller.gameOver()
		case VICTORY:
			controller.vict()
		case NEWGAME:
			controller.newGame(key)
		case EXIT:
			return
		}

	}
}
