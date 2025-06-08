package quotes

import (
	"ago-launcher/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/adam-macioszek/lotr-sdk/quote"
)

type Qouter struct {
	Quotes Quotes
}

type Quotes struct {
	Quotes []Quote `json:"quotes"`
}

type Quote struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

func (qouter *Qouter) LoadQuotes() (Quotes, error) {
	var quotes Quotes

	fmt.Println("[Quoter] Loading quotes.json")
	jsonFile, err := os.Open("resources/quotes.json")
	if err != nil {
		fmt.Println("[Quoter] could not load quote file:", err)
		return Quotes{}, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("[Quoter] could not read quote file:", err)
		return Quotes{}, err
	}

	err = json.Unmarshal(byteValue, &quotes)
	if err != nil {
		fmt.Println("[Quoter] could not unmarshal quote file:", err)
		return Quotes{}, err
	}

	numQuotes := len(quotes.Quotes)
	if numQuotes <= 0 {
		fmt.Println("[Quoter] no quotes found")
		return quotes, nil
	}

	fmt.Println("[Quoter] Found", numQuotes, "quotes")

	return quotes, nil
}

func (qouter *Qouter) GetRandomQuote() (Quote, error) {
	fmt.Println("[Quoter] Getting random quote")
	if len(qouter.Quotes.Quotes) == 0 {
		quotes, err := qouter.LoadQuotes()
		if err != nil {
			fmt.Println("[Quoter] error loading quotes:", err)
		}
		qouter.Quotes = quotes
	}
	quote, err := utils.RandomElement(qouter.Quotes.Quotes)
	if err != nil {
		fmt.Println("[Quoter] error getting random quote:", err)
	} else {
		fmt.Println("[Quoter] Found random quote:", quote.Quote)
	}
	return quote, err
}

func (qouter *Qouter) PrintAllQuotes() {
	fmt.Println("[Quoter] Getting all LOTR quotes")

	quotes, err := quote.GetAllQuotes()
	if err != nil {
		log.Println("[Quoter] Failed to retrieve quotes")
		log.Println("[Quoter]", err)
	}

	for _, v := range quotes {
		fmt.Println("[Quoter]", v.Dialog, "-", v.CharacterID)
	}

	fmt.Printf("[Quoter] Found %v quotes\n", len(quotes))
}