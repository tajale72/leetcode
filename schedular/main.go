package main

import (
    "github.com/faiface/pixel"
    "github.com/faiface/pixel/pixelgl"
    "github.com/faiface/pixel/text"
    "golang.org/x/image/font/basicfont"
    "golang.org/x/image/colornames"
    "fmt"
	"time"
)

func run() {
    cfg := pixelgl.WindowConfig{
        Title:  "My First Game",
        Bounds: pixel.R(0, 0, 1024, 768),
        VSync:  true,
    }
    win, err := pixelgl.NewWindow(cfg)
    if err != nil {
        panic(err)
    }

    basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(120, 500), basicAtlas)

	fmt.Fprintln(basicTxt, "Hello, Evryone!")
	time.Sleep( 2 *time.Second)
	fmt.Fprintln(basicTxt, "I support multiple lines!")
	time.Sleep( 2 *time.Second)
	fmt.Fprintf(basicTxt, "And I'm an %s, yay!", "io.Writer")
	time.Sleep( 2 *time.Second)

    for !win.Closed() {
        win.Clear(colornames.Black)
		basicTxt.Draw(win, pixel.IM)
		win.Update()
    }
}

func main() {
    pixelgl.Run(run)
}