package main

import (
	"fmt"
	_ "image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1250
	screenHeight = 750
)

var (
	err error
)

type Game struct {
	stateMachine StateMachine
	defaultFont  Font
	codonFont    Font
	temp_tfa     TFA
}

func loadFile(image string) string {
	// Construct path to file
	imagepath := filepath.Join("Assets", "Images", image)
	// Open file
	file, err := os.Open(imagepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "Error"
	}

	defer file.Close()

	return file.Name()
}

func (g *Game) init() {
	g.defaultFont = newFont("CourierPrime-Regular.ttf", 32)
	g.codonFont = newFont("Honk-Regular.ttf", 90)
	var s_map = SceneConstructorMap{
		"menu": newMainMenu,
		"info": newInfoPage,
	}
	g.stateMachine = newStateMachine(s_map)
}

func (g *Game) Update() error {
	g.stateMachine.update(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.stateMachine.draw(g, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &Game{}
	game.init()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
