package main

import (
	"log"
	"os"
)

var key string

func main() {

	Init()
	var ok bool
	// Access environment variable
	key, ok = os.LookupEnv("HACKMAN_API_KEY")
	if !ok {
		log.Fatalln("missing API key in env")
	}

	//fmt.Printf("%s\n", GetWord(-1))

	Setup(5)

	GameLoop()

}
