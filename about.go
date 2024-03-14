package main

import (
	//"log"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type About struct {
	aboutBg          	StillImage
	aboutToMenuButton 	Button
	message				string
}

func newAbout(g *Game) {
	g.stateMachine.state = &About{}
	aboutSprites = []GUI{}
}

func (a *About) Init(g *Game) {
	a.aboutBg = newStillImage("AboutBg.png", newRect(0, 0, 1250, 750))
	a.aboutToMenuButton = newButton("menuButton.png", newRect(350, 450, 300, 200), ToMenu)
	a.message = "WELCOME TO THE CELL\nSIGNALING PATHWAY\nSIMULATOR!\n"
	a.message += "This simulator will\nguide you through the\ncomplete cell signaling\n"
	a.message += "pathway from reception\nthrough translation!\nClick the play "
	a.message += "button\nor select a level\nto begin."
	aboutSprites = []GUI{&a.aboutBg, &a.aboutToMenuButton}
	state_array = aboutSprites
}

func (a *About) Update(g *Game) {
	ebiten.SetWindowTitle("Cell Signaling Pathway - About")
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
