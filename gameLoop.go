package hackman

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GameLoop(key string) {
	var c int32
	numGuesses := 6
	var Ans WordInfo
	Ans.Setup(-1, key)
	var darkGreen rl.Color = rl.Color{0x11, 0x28, 0x21, 0xFF}

	//rl.BeginDrawing()
	for !rl.WindowShouldClose() {

		c = rl.GetKeyPressed()
		if c != 0 {
			if Ans.inWord(uniToInt(c)) {
			} else {
				numGuesses--
			}
		}

		if numGuesses <= 0 {
			gameOver(Ans)
		}

		if Ans.fullyGuessed() {
			vict(Ans)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawTexture(background, 0, 0, rl.RayWhite)
		rl.DrawRectangle(100, 100, 600, 400, darkGreen)

		Ans.drawWord()

		rl.DrawText("Lives: "+strconv.Itoa(numGuesses), 575, 100, 30, rl.Green)
		rl.EndDrawing()

	}

	DeInit()

}

func gameOver(Ans WordInfo) {

	var option int32 = 0

	var buttons []Button
	buttons = append(buttons, *NewButton(button1, (rl.GetScreenWidth()/2-200)/2, ((rl.GetScreenHeight() + 50) / 2), rl.Red, "Exit", DeInit))
	buttons = append(buttons, *NewButton(button1, (rl.GetScreenWidth()/2+200)/2, ((rl.GetScreenHeight() + 50) / 2), rl.Green, "New Game", gameLoop))

	for option == 0 && !rl.WindowShouldClose() {

		mouseClick(rl.MouseLeftButton, buttons)

		rl.BeginDrawing()
		//rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Gameover! The answer was "+Ans.word, 120, 120, 25, rl.Red)
		drawButtons(buttons)
		rl.EndDrawing()

	}

}

func mainMenu() {

	var buttons []Button
	buttons = append(buttons, *NewButton(button1, (rl.GetScreenWidth()-200)/2, (rl.GetScreenHeight()+50)/2, rl.SkyBlue, "Start", gameLoop))
	scrollX := 0
	speed := 2
	//rl.SetTextureWrap(title, rl.WrapRepeat)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		//rl.ClearBackground(rl.RayWhite)
		rl.DrawTexture(background, int32(scrollX/speed), 0, rl.RayWhite)
		rl.DrawTexture(background, int32((scrollX/speed - 800)), 0, rl.RayWhite)

		rl.DrawTexture(title, int32((rl.GetScreenWidth()-650)/2), int32((rl.GetScreenHeight()-400)/2), rl.RayWhite)

		drawButtons(buttons)

		rl.EndDrawing()

		mouseClick(rl.MouseLeftButton, buttons)
		scrollX++
		if scrollX > int(background.Width)*speed {
			scrollX = 0
		}

	}
	DeInit()

}

func vict(Ans WordInfo) {

	var buttons []Button
	buttons = append(buttons, *NewButton(button1, (rl.GetScreenWidth()/2-200)/2, ((rl.GetScreenHeight() + 50) / 2), rl.Red, "Exit", DeInit))
	buttons = append(buttons, *NewButton(button1, (rl.GetScreenWidth()/2+200)/2, ((rl.GetScreenHeight() + 50) / 2), rl.Green, "New Game", gameLoop))

	for !rl.WindowShouldClose() {

		mouseClick(rl.MouseLeftButton, buttons)

		rl.BeginDrawing()
		drawButtons(buttons)
		rl.DrawText("You guessed the word!", 120, 120, 30, rl.Green)
		Ans.drawWord()
		rl.EndDrawing()

	}

}
