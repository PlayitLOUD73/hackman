package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type HangmanWord struct {
	Word string
}

func main() {

	var word HangmanWord

	Init()

	key, ok := os.LookupEnv("HACKMAN_API_KEY")
	if !ok {
		log.Fatalln("missing API key in env")
	}

	API_ENDPOINT := "http://clemsonhackman.com/api/word"

	resp, err := http.Get(API_ENDPOINT + "?key=" + key)
	if err != nil {
		log.Fatalln("cannot connect to api: ", err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bodyBytes, &word)
	if err != nil {
		log.Fatalln("unable to unmarshall json: ", err)
	}
	fmt.Printf("%s\n", word.Word)

}
