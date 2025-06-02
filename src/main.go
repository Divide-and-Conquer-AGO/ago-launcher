package main

import (
	"ago-launcher/gui"
	"ago-launcher/utils"
	"fmt"
)

func main() {
	var appData gui.AppData
	quotes, err := utils.LoadQuotes()
	if err != nil {
		fmt.Println(err)
	}
	appData.Quotes = quotes
	gui.InitGUI(appData)
}