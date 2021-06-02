package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Definition struct {
	Text         string `json:"text"`
	PartOfSpeech string `json:"partOfSpeech"`
}

var abbreviations = map[string]string{
	"noun":            "n.",
	"verb":            "v.",
	"transitive verb": "v.",
	"adjective":       "adj.",
}

func main() {
	if len(os.Args) < 2 {
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

	var definitions []Definition
	err = json.Unmarshal(body, &definitions)
	if err != nil {
		panic("Error parsing json")
	}
	fmt.Printf("\n%v\n\n", word)
	for i := 0; i < len(definitions) && i <= 3; i++ {
		if len(definitions[i].Text) == 0 {
			continue
		}
		fmt.Printf("%v\n\n", formatDefintion((definitions[i])))
	}
}

func formatDefintion(defintion Definition) string {
	return fmt.Sprintf("%-*v%v", 6, abbreviations[defintion.PartOfSpeech], defintion.Text)
}
