package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Word struct {
	Text string `json:"text"`
}

func main() {
	if len(os.Args) < 1 {
		fmt.Print("\nPlease pass in a word to define.\n\n  grok <word>\n\n")
		return
	}

	word := os.Args[1]
	apiKey := os.Getenv("WORDNIK_API")
	url := fmt.Sprintf("http://api.wordnik.com/v4/word.json/%v/definitions?api_key=%v", word, apiKey)

	res, err := http.Get(url)
	if err != nil {
		panic("Error from request")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic("Error while reading response")
	}

	var words []Word
	err = json.Unmarshal(body, &words)
	if err != nil {
		panic("Error parsing json")
	}
	fmt.Printf("%v", words[0].Text)
}
