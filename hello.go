package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adam-macioszek/lotr-sdk/quote"
)

func main() {
	fmt.Println("Hello AGO users!")
	log.Println(os.Getenv("LOTR_API_KEY"))

	quoteId := "5cd96e05de30eff6ebccebb1"
	testQuote, err := quote.GetQuoteByID(quoteId)
	if err != nil {
		log.Println(err)
	}
	log.Println(testQuote.Dialog)
}
