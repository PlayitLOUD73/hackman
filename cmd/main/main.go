package main

import (
	_ "embed"
	"strings"

	game "github.com/PlayitLOUD73/hackman"
)

//go:embed "secrets.txt"
var rawKey string

var key string

func main() {

	keyInfo := strings.Split(rawKey, "=")
	key = keyInfo[1]

	game.Init()
	game.GameLoop(key)
	game.DeInit()

}
