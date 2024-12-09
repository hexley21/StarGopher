package renderer

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hexley21/star-gopher/internal/object"
)

type Renderer struct {
    width, height int
}

func NewRenderer(width, height int) *Renderer {
    return &Renderer{width: width, height: height}
}

func (r *Renderer) Render(gameObjects []object.GameObject) {
    ClearScreen()

    buffer := r.initBuffer()

    for _, obj := range gameObjects {
        if obj.IsActive() {
            x, y := obj.GetX(), obj.GetY()
            if x >= 0 && x < r.width && y >= 0 && y < r.height {
                buffer[y][x] = obj.GetAppearance()
            }
        }
    }

    for _, row := range buffer {
        fmt.Println(string(row))
    }
}

func (r *Renderer) initBuffer() [][]rune {
	buffer := make([][]rune, r.height)
    for i := range buffer {
        buffer[i] = make([]rune, r.width)
        for j := range buffer[i] {
            buffer[i][j] = ' '
        }
    }

	return buffer
}

func ClearScreen() {
    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func HideCursor() {
    fmt.Print("\033[?25l")
}

func ShowCursor() {
    fmt.Print("\033[?25h")
}

func Kbhit() bool {
    return false // Implement this function to check for keyboard input
}

func Getch() rune {
    return ' ' // Implement this function to get a single character from keyboard input
}
