package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type HangmanWord struct {
	Word string
}

func respToWord(resp *http.Response) string {

	var word HangmanWord

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	err := json.Unmarshal(bodyBytes, &word)
	if err != nil {
		log.Fatalln("unable to unmarshall json: ", err)
	}

	return word.Word
}

func GetWord(length int) string {

	API_ENDPOINT := "https://clemsonhackman.com/api/word"

	url := API_ENDPOINT + "?key=" + key
	if length != 0 {
		url += "&length=" + strconv.Itoa(length)
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("cannot connect to api: ", err)
	}

	return respToWord(resp)
}
