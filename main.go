package main

import (
	_ "embed"
	"strings"
)

//go:embed secrets.txt
var rawKey string

var key string

func main() {

	keyInfo := strings.Split(rawKey, "=")
	key = keyInfo[1]

	Init()
	mainMenu()

}
