package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type InfoPage struct {
	infoBg           *ebiten.Image
	infoToMenuButton Button
}

func newInfoPage(g *Game) {
	g.stateMachine.state = InfoPage{
		infoToMenuButton: newButton("menuButton.png", newRect(300, 375, 242, 138), ToMenu),
		infoBg:           nil,
	}
}

func (i InfoPage) Init() {
	i.infoBg, _, err = ebitenutil.NewImageFromFile(loadFile("InfoBg.png"))
	if err != nil {
		log.Fatal(err)
	}
}

func (i InfoPage) Update(g *Game) {
	ebiten.SetWindowTitle("Cell Signaling Pathway - About")
	ebiten.SetWindowSize(screenWidth, screenHeight)
	i.infoToMenuButton.on_click(g)
}

func (i InfoPage) Draw(g *Game, screen *ebiten.Image) {
	screen.DrawImage(i.infoBg, nil)
	i.infoToMenuButton.draw(screen)
	m := "WELCOME TO THE CELL\nSIGNALING PATHWAY\nSIMULATOR!\n"
	m += "This simulator will\nguide you through the\ncomplete cell signaling\n"
	m += "pathway from reception\nthrough translation!\nClick the play "
	m += "button\nor select a level\nto begin."
	g.defaultFont.drawFont(screen, m, 775, 275, color.RGBA{50, 5, 5, 250})
}
