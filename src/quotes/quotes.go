package quotes

import (
	"ago-launcher/utils"
	"encoding/json"
	"io"
	"net/http"

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

	utils.Logger().Println("[Quoter] Loading quotes.json")
	resp, err := http.Get("https://raw.githubusercontent.com/Divide-and-Conquer-AGO/ago-launcher/refs/heads/main/src/resources/quotes.json")
	if err != nil {
		utils.Logger().Println("could not fetch modVersions file from GitHub")
		return Quotes{}, err
	}
	jsonFile := resp.Body
	defer resp.Body.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		utils.Logger().Println("[Quoter] could not read quote file:", err)
		return Quotes{}, err
	}

	err = json.Unmarshal(byteValue, &quotes)
	if err != nil {
		utils.Logger().Println("[Quoter] could not unmarshal quote file:", err)
		return Quotes{}, err
	}

	numQuotes := len(quotes.Quotes)
	if numQuotes <= 0 {
		utils.Logger().Println("[Quoter] no quotes found")
		return quotes, nil
	}

	utils.Logger().Printf("[Quoter] Found %d quotes\n", numQuotes)

	return quotes, nil
}

func (qouter *Qouter) GetRandomQuote() (Quote, error) {
	utils.Logger().Println("[Quoter] Getting random quote")
	if len(qouter.Quotes.Quotes) == 0 {
		quotes, err := qouter.LoadQuotes()
		if err != nil {
			utils.Logger().Println("[Quoter] error loading quotes:", err)
		}
		qouter.Quotes = quotes
	}
	quote, err := utils.RandomElement(qouter.Quotes.Quotes)
	if err != nil {
		utils.Logger().Println("[Quoter] error getting random quote:", err)
	} else {
		utils.Logger().Printf("[Quoter] Found random quote: %s\n", quote.Quote)
	}
	return quote, err
}

func (qouter *Qouter) PrintAllQuotes() {
	utils.Logger().Println("[Quoter] Getting all LOTR quotes")

	quotes, err := quote.GetAllQuotes()
	if err != nil {
		utils.Logger().Println("[Quoter] Failed to retrieve quotes")
		utils.Logger().Println("[Quoter]", err)
	}

	for _, v := range quotes {
		utils.Logger().Printf("[Quoter] %s - %s\n", v.Dialog, v.CharacterID)
	}

	utils.Logger().Printf("[Quoter] Found %d quotes\n", len(quotes))
}
