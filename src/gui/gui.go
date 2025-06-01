package gui

import (
	"ago-launcher/src/config"
	"fmt"

	"cogentcore.org/core/core"
)

func InitGUI() *core.Body {
	b := core.NewBody()
	return b
}

func RenderInputs(body *core.Body) {
	cfgFile := config.LoadConfigFile()
	cfg := config.ParseConfig(cfgFile)

	form := core.NewForm(body).SetStruct(&cfg)

	fmt.Println(form)
}
