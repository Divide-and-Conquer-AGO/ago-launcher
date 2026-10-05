<h1 align="center" style="color:#bf6f00;">
  <a href="https://www.divide-and-conquer-ago.com">Divide and Conquer: V6 Launcher</a>
</h1>

<div align="center">
  <a href="https://discord.gg/yVHm7kBTAY">
    <img src="https://img.shields.io/discord/759414542240972840?style=for-the-badge&label=Discord&color=bf6f00" >
  </a>
  <div>
    <img src="docs/img/image-5.png" >
    <img src="docs/img/image-4.png" >
    <img src="docs/img/image-6.png" >
    <img src="docs/img/image-3.png" >
  </div>
</div>

-----------------
## Development
0. Install [Go](https://go.dev/doc/install) and [Fyne](https://docs.fyne.io/started/)
1. Install `air` for hot reload support

```shell
go install github.com/air-verse/air@latest
```

2. Start the project

```shell
cd src 
air
```

This will build the binary (DaC_Launcher.exe) and run it from `resources/mods/dac_beta` where there are various config files and example folders to use

If you want to test it on an actual mod folder in it's packaged state, you can run

```shell
cd src
fyne package -release -os windows && xcopy /Y "DaC_Launcher.exe" "E:\Steam\steamapps\common\Medieval II Total War\mods\dac-beta\DaC_Launcher.exe" &&"E:\Steam\steamapps\common\Medieval II Total War\mods\dac-beta\DaC_Launcher.exe"
```

Regenerate bundled images and fonts
```shell
fyne bundle -package gui -o gui/bundle.go icon.png && fyne bundle -a -o gui/bundle.go background.png && fyne bundle -a -o gui/bundle.go tolkien.png && fyne bundle -a -o gui/bundle.go favicon.ico && fyne bundle -a -o gui/bundle.go LTMuseum-Black.ttf && fyne bundle -a -o gui/bundle.go LTMuseum-Bold.ttf && fyne bundle -a -o gui/bundle.go LTMuseum-Italic.ttf
```