package hackman

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func GameLoop(key string) {

	var controller GameController = *NewGameController()
	controller.state = MAINMENU

	//controller.newGame(key)

	//rl.BeginDrawing()
	for !rl.WindowShouldClose() {
		rl.ClearBackground(rl.RayWhite)
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
