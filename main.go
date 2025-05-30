package main

import (
	"ago-launcher/internal/utils"
	"fmt"
)

func main() {
	fmt.Println("Launching AGO launcher..")

	cfg := config.LoadConfig()

	section := cfg.Section("scripts")

	fmt.Println(section.Key("natural_disasters").String())
	fmt.Println(section.Key("random_aa_ai_start").String())
	fmt.Println(section.Key("merge_dol_amroth").String())
	fmt.Println(section.Key("randomized_start").String())

	// quote0 := api.GetRandomQuote()
	// quote1 := api.GetRandomQuote()
	// quote2 := api.GetRandomQuote()
	// quote3 := api.GetRandomQuote()
	// quote4 := api.GetRandomQuote()
	// quote5 := api.GetRandomQuote()

	// b := core.NewBody()

	// core.NewButton(b).SetText(quote0.Dialog)
	// core.NewButton(b).SetText(quote1.Dialog)
	// core.NewButton(b).SetText(quote2.Dialog)
	// core.NewButton(b).SetText(quote3.Dialog)
	// core.NewButton(b).SetText(quote4.Dialog)
	// core.NewButton(b).SetText(quote5.Dialog)
	// b.RunMainWindow()
}
