//MUST UPDATE ALL FUNCTIONS WITH NEW CODE

package main

import (
	//"log"
	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type MainMenu struct {
	protoStartBg StillImage
	startBg      Parallax
	startP1      Parallax
	startP2      Parallax
	startP3      Parallax
	startP4      Parallax
	fixedStart   StillImage
	playbutton   Button
	aboutButton  Button
	levSelButton Button
	volButton    VolButton
}

func newMainMenu(g *Game) {
	g.stateMachine.state = MainMenu{}
	menuSprites = []GUI{}
}

func (m MainMenu) Init(g *Game) {
	m.protoStartBg= newStillImage("MenuBg.png", newRect(0, 0, 1250, 750))
	m.startBg= newParallax("StartBg.png", newRect(0, 0, 1250, 750), 5)
	m.startP1= newParallax("parallax-Start2.png", newRect(0, 0, 1250, 750), 4)
	m.startP2= newParallax("parallax-Start3.png", newRect(0, 0, 1250, 750), 3)
	m.startP3= newParallax("parallax-Start4.png", newRect(0, 0, 1250, 750), 2)
	m.startP4= newParallax("parallax-Start5.png", newRect(0, 0, 1250, 750), 1)
	m.fixedStart= newStillImage("fixed-Start.png", newRect(0, 0, 1250, 750))
	m.playbutton= newButton("PlayButton.png", newRect(750, 100, 300, 200), ToPlasma)
	m.aboutButton= newButton("aboutButton.png", newRect(770, 260, 300, 200), ToAbout)
	m.levSelButton= newButton("levSelButton.png", newRect(700, 450, 300, 200), ToLevelSelect)
	m.volButton= newVolButton("volButtonOn.png", newRect(100, 100, 165, 165), m.volButton.SwitchVol, *audioPlayer)
	menuSprites = []GUI{m.protoStartBg, &m.startBg, &m.startP1, &m.startP2, &m.startP3, &m.startP4, m.fixedStart, m.playbutton, m.aboutButton, m.levSelButton, m.volButton}
	state_array = menuSprites
	m.volButton.SwitchVol(g)
}

func (m MainMenu) Update(g *Game) {
	ebiten.SetWindowTitle("Cell Signaling Pathway - Main Menu")
		for _, element := range menuSprites {
		switch element.(type) {
		case (Button):
			element.update(g)
		case (VolButton):
			element.update(g)
		default:
			element.update()
		}
	}
}

func (m MainMenu) Draw(g *Game, screen *ebiten.Image) {
	for _, element := range menuSprites {
		if scaleChange > 0 {
			element.scaleToScreen(screen)
		}
		element.draw(screen)
	}
}
