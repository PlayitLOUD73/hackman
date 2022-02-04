package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type HangmanWord struct {
	Word string
}

func main() {

	var word HangmanWord

	var API_ENDPOINT string
	var API_KEY string

	API_ENDPOINT = "http://clemsonhackman.com/api/word"
	API_KEY = "" // placeholder

	resp, err := http.Get(API_ENDPOINT + "?key=" + API_KEY)
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
