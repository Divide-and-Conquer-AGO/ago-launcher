package api

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/adam-macioszek/lotr-sdk/quote"
)

func PrintAllQuotes() {
	fmt.Println("Getting all LOTR quotes")

	quotes, err := quote.GetAllQuotes()
	if err != nil {
		log.Println("Failed to retrieve quotes")
		log.Println(err)
	}

	for _, v := range quotes {
		fmt.Println(v.Dialog, "-", v.CharacterID)
	}

	fmt.Printf("Found %v quotes", len(quotes))
}

func GetRandomQuote() quote.Quote {
	fmt.Println("Getting random LOTR quote")

	quotes, err := quote.GetAllQuotes()
	if err != nil {
		log.Println("Failed to retrieve quotes")
		log.Println(err)
		return quote.Quote{}
	}

	return quotes[rand.Intn(len(quotes))]
}
