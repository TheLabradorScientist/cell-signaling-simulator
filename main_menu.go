package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type MainMenu struct {
	startBg      *ebiten.Image
	playbutton   Button
	infoButton   Button
	levSelButton Button
}

func newMainMenu(g *Game) {
	g.stateMachine.state = MainMenu{
		playbutton:   newButton("PlayButton.png", newRect(750, 100, 242, 138), ToLevel1),
		infoButton:   newButton("infoButton.png", newRect(770, 260, 242, 138), ToInfo),
		levSelButton: newButton("levSelButton.png", newRect(700, 450, 232, 140), ToLevelSelect),
		startBg:      nil,
	}
}

func (m MainMenu) Init() {
	m.startBg, _, err = ebitenutil.NewImageFromFile(loadFile("MenuBg.png"))
	if err != nil {
		log.Fatal(err)
	}
}

func (m MainMenu) Update(g *Game) {
	ebiten.SetWindowTitle("Cell Signaling Pathway - Main Menu")
	ebiten.SetWindowSize(screenWidth, screenHeight)
	m.infoButton.on_click(g)
	m.playbutton.on_click(g)
	m.levSelButton.on_click(g)
}

func (m MainMenu) Draw(g *Game, screen *ebiten.Image) {
	screen.DrawImage(m.startBg, nil)
	m.playbutton.draw(screen)
	m.infoButton.draw(screen)
	m.levSelButton.draw(screen)
}
