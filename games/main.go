package main

import (
	"fmt"
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

func run() {
	//Create a config for the window
	cfg := pixelgl.WindowConfig{
		Title:  "My software",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Println(err)
	}

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(200, 500), basicAtlas)

	fmt.Fprintln(basicTxt, "Hello, text!")
	fmt.Fprintln(basicTxt, "I support multiple lines!")
	fmt.Fprintf(basicTxt, "And I'm an %s, yay!", "io.Writer")

	win.Clear(colornames.Black)

	for !win.Closed() {
		basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 4))
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

func OpenWindow() {

}
