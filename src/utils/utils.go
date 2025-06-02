package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"os"
)

type Quotes struct {
	Quotes []Quote `json:"quotes"`
}

type Quote struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

func LoadQuotes() (Quotes, error) {
	jsonFile, err := os.Open("resources/quotes.json")
	if err != nil {
		return Quotes{}, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return Quotes{}, err
	}
	var quotes Quotes

	if err := json.Unmarshal(byteValue, &quotes); err != nil {
		return Quotes{}, err
	}

	if len(quotes.Quotes) > 0 {
		return quotes, nil
	}

	return Quotes{}, errors.New("failed to read quote file")
}

func RandomElement[T any](slice []T) (T, error) {
	var zero T
	if len(slice) == 0 {
		return zero, errors.New("no elements available")
	}
	return slice[rand.Intn(len(slice))], nil
}
