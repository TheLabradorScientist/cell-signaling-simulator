//MUST UPDATE ALL FUNCTIONS WITH NEW CODE

package main

import (
	//"log"
	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	// MENU SPRITES
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
)

type MainMenu struct {
	menuSprites []GUI
}

func newMainMenu(g *Game) {
	g.stateMachine.state = MainMenu{
		menuSprites: []GUI{},
	}
	SwitchVol(g)
}

func (m MainMenu) Init() {
	protoStartBg= newStillImage("MenuBg.png", newRect(0, 0, 1250, 750))
	startBg= newParallax("StartBg.png", newRect(0, 0, 1250, 750), 5)
	startP1= newParallax("parallax-Start2.png", newRect(0, 0, 1250, 750), 4)
	startP2= newParallax("parallax-Start3.png", newRect(0, 0, 1250, 750), 3)
	startP3= newParallax("parallax-Start4.png", newRect(0, 0, 1250, 750), 2)
	startP4= newParallax("parallax-Start5.png", newRect(0, 0, 1250, 750), 1)
	fixedStart= newStillImage("fixed-Start.png", newRect(0, 0, 1250, 750))
	playbutton= newButton("PlayButton.png", newRect(750, 100, 300, 200), ToPlasma)
	aboutButton= newButton("aboutButton.png", newRect(770, 260, 300, 200), ToAbout)
	levSelButton= newButton("levSelButton.png", newRect(700, 450, 300, 200), ToLevelSelect)
	volButton= newVolButton("volButtonOn.png", newRect(100, 100, 165, 165), SwitchVol, *audioPlayer)
	menuSprites = []GUI{protoStartBg, &startBg, &startP1, &startP2, &startP3, &startP4, fixedStart, playbutton, aboutButton, levSelButton, volButton}
	state_array = menuSprites
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
		element.draw(screen)
	}
}