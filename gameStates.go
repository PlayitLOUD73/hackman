package hackman

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func game(Ans WordInfo, c int32, guesses int) (WordInfo, int, GameState) {
	state := GAME
	var darkGreen rl.Color = rl.Color{0x11, 0x28, 0x21, 0xFF}
	if c != 0 {
		if Ans.inWord(uniToInt(c)) {
		} else {
			guesses--
		}
	}

	if guesses <= 0 {
		state = GAMEOVER
	}

	if Ans.fullyGuessed() {
		state = VICTORY
	}

	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)
	rl.DrawTexture(background, 0, 0, rl.RayWhite)
	rl.DrawRectangle(100, 100, 600, 400, darkGreen)

	Ans.drawWord()

	rl.DrawText("Lives: "+strconv.Itoa(guesses), 575, 100, 30, rl.Green)
	rl.EndDrawing()

	return Ans, guesses, state
}

func gameOver(Ans WordInfo) GameState {

	g := GAMEOVER
	var buttons []Button
	buttons = append(buttons, *NewButton(button1, (rl.GetScreenWidth()/2-200)/2, ((rl.GetScreenHeight() + 50) / 2), rl.Red, "Exit", EXIT))
	buttons = append(buttons, *NewButton(button1, (rl.GetScreenWidth()/2+200)/2, ((rl.GetScreenHeight() + 50) / 2), rl.Green, "New Game", NEWGAME))

	g = mouseClick(rl.MouseLeftButton, buttons, g)

	rl.BeginDrawing()

	rl.DrawTexture(background, 0, 0, rl.DarkBrown)
	rl.DrawText("Gameover! The answer was "+Ans.word, 120, 120, 25, rl.Red)
	drawButtons(buttons)
	rl.EndDrawing()
	return g

}

func mainMenu() GameState {

	g := MAINMENU
	var buttons []Button
	buttons = append(buttons, *NewButton(button1, (rl.GetScreenWidth()-200)/2, (rl.GetScreenHeight()+50)/2, rl.SkyBlue, "Start", NEWGAME))
	scrollX := 0
	speed := 2
	//rl.SetTextureWrap(title, rl.WrapRepeat)
	rl.BeginDrawing()
	//rl.ClearBackground(rl.RayWhite)
	rl.DrawTexture(background, int32(scrollX/speed), 0, rl.RayWhite)
	rl.DrawTexture(background, int32((scrollX/speed - 800)), 0, rl.RayWhite)

	rl.DrawTexture(title, int32((rl.GetScreenWidth()-650)/2), int32((rl.GetScreenHeight()-400)/2), rl.RayWhite)

	drawButtons(buttons)

	rl.EndDrawing()

	g = mouseClick(rl.MouseLeftButton, buttons, g)
	scrollX++
	if scrollX > int(background.Width)*speed {
		scrollX = 0
	}
	return g

}

func vict(Ans WordInfo) GameState {

	g := VICTORY
	var buttons []Button
	buttons = append(buttons, *NewButton(button1, (rl.GetScreenWidth()/2-200)/2, ((rl.GetScreenHeight() + 50) / 2), rl.Red, "Exit", EXIT))
	buttons = append(buttons, *NewButton(button1, (rl.GetScreenWidth()/2+200)/2, ((rl.GetScreenHeight() + 50) / 2), rl.Green, "New Game", NEWGAME))

	g = mouseClick(rl.MouseLeftButton, buttons, g)

	rl.BeginDrawing()
	drawButtons(buttons)
	rl.DrawText("You guessed the word!", 120, 120, 30, rl.Green)
	Ans.drawWord()
	rl.EndDrawing()

	return g
}

func newGame(key string) (WordInfo, int) {
	var Ans WordInfo
	Ans.Setup(-1, key)
	return Ans, 6
}
