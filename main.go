package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 960
	screenHeight = 1120
)

var (
	err error
	scene string = "Main Menu"
	startBg      *ebiten.Image
	startButton	*ebiten.Image
)


type Game struct {
	//title               string
	defaultFont         Font
	switchedMainToStart bool
	switchedStartToMain bool
}


func init() {
	startBg, _, err = ebitenutil.NewImageFromFile("Cell.png")
	if err != nil {
		log.Fatal(err)
	}
	startButton, _, err = ebitenutil.NewImageFromFile("red-start.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	cursorX, cursorY := ebiten.CursorPosition()
	switch scene {
	case "Main Menu":
		ebiten.SetWindowTitle("Main Menu")
		ebiten.SetWindowSize(screenWidth, screenHeight)
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && cursorX >= 420 && cursorX <= 540 && cursorY >= 480 && cursorY <= 620 {
			g.switchedMainToStart = true
			g.switchedStartToMain = false
		}
	case "Signal Reception":
		ebiten.SetWindowTitle("Signal Reception")
		ebiten.SetWindowSize(screenWidth, screenHeight)
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && cursorX >= 25 && cursorX <= 50 && cursorY >= 25 && cursorY <= 50 {
			g.switchedMainToStart = false
			g.switchedStartToMain = true
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch scene {
	case "Main Menu":
		op := &ebiten.DrawImageOptions{}
		//op.GeoM.Translate(300, 250)
		screen.DrawImage(startBg, nil)
		screen.DrawImage(startButton, op)
		vector.DrawFilledRect(screen, 50, 50, 50, 50, color.Alpha16{0xF81F}, false)
		if g.switchedMainToStart {
			scene = "Signal Reception"
		}
	case "Signal Reception":
		vector.DrawFilledRect(screen, 25, 25, 25, 25, color.White, false)
		g.defaultFont.drawFont(screen, "Hello World", 100, 100, color.White)
		if g.switchedStartToMain {
			scene = "Main Menu"
		}
	}
}


func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &Game{}
	game.defaultFont = newFont("ThaleahFat.ttf", 25)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
