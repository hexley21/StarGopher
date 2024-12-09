package main

import (
	"github.com/hexley21/star-gopher/internal/engine"
	"github.com/hexley21/star-gopher/internal/renderer"
)

func main() {
    width := 20
    height := 20

    renderer.HideCursor()
    defer renderer.ShowCursor()

    game := engine.NewGameEngine(width, height)
    game.Run()
}
