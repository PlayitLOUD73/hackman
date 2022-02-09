package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// HangmanWord is a struct that contains a
// word. This is used for easy JSON unmarshaling.
type HangmanWord struct {
	Word string
}

// GetWord returns a word of size length by
// accessing the Clemson Hackman API.
// If length is -1, the word size is random.
// Length can be from 5 to 14 or -1.
func GetWord(length int, key string) string {

	var word HangmanWord

	API_ENDPOINT := "https://clemsonhackman.com/api/word"

	url := API_ENDPOINT + "?key=" + key
	if length != -1 {
		url += "&length=" + strconv.Itoa(length)
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("cannot connect to api: ", err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bodyBytes, &word)
	if err != nil {
		log.Fatalln("unable to unmarshall json: ", err)
	}

	return word.Word
}
