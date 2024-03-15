package main

import (
	//"log"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Sprites in About
type About struct {
	aboutBg          	StillImage
	aboutToMenuButton 	Button
	message				string
}

var aboutStruct *About

// Initialize about struct and aboutSprites array if not initialized, then set state to aboutStruct
func newAbout(g *Game) {
	if len(aboutSprites ) == 0 {
		aboutStruct = &About{
			aboutBg: newStillImage("AboutBg.png", newRect(0, 0, 1250, 750)),
			aboutToMenuButton: newButton("menuButton.png", newRect(350, 450, 300, 200), ToMenu),
			message:	"WELCOME TO THE CELL\nSIGNALING PATHWAY\nSIMULATOR!\n" +
						"This simulator will\nguide you through the\ncomplete cell signaling\n" +
						"pathway from reception\nthrough translation!\nClick the play button\n" +
						"or select a level\nto begin.",
		}
		aboutSprites = []GUI{&aboutStruct.aboutBg, &aboutStruct.aboutToMenuButton}
	}
	g.stateMachine.state = aboutStruct
}

func (a *About) Init(g *Game) {
	ebiten.SetWindowTitle("Cell Signaling Pathway - About")
	state_array = aboutSprites
}

func (a *About) Update(g *Game) {
	for _, element := range aboutSprites {
		element.update(g)
	}
}

func (a *About) Draw(g *Game, screen *ebiten.Image) {
	for _, element := range aboutSprites {
		element.draw(screen)
	}
	defaultFont.drawFont(screen, a.message, 775, 275, color.RGBA{50, 5, 5, 250})
}
