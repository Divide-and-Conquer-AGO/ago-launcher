# Usage: make package modFolder=E:\Steam\steamapps\common\Medieval II Total War\mods\ago_beta

# Name of your executable (edit if your app name is different)
EXE_NAME := AGO_Launcher.exe

.PHONY: run package copy 

run:
	air

package:
	fyne package
	copy /Y $(EXE_NAME) "$(modFolder)\"