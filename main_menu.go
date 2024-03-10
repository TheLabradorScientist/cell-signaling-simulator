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
	g.stateMachine.state = MainMenu{
	protoStartBg: newStillImage("MenuBg.png", newRect(0, 0, 1250, 750)),
	startBg: newParallax("StartBg.png", newRect(0, 0, 1250, 750), 5.0),
	startP1: newParallax("parallax-Start2.png", newRect(0, 0, 1250, 750), 4.0),
	startP2: newParallax("parallax-Start3.png", newRect(0, 0, 1250, 750), 3.0),
	startP3: newParallax("parallax-Start4.png", newRect(0, 0, 1250, 750), 2.0),
	startP4: newParallax("parallax-Start5.png", newRect(0, 0, 1250, 750), 1.0),
	fixedStart: newStillImage("fixed-Start.png", newRect(0, 0, 1250, 750)),
	playbutton: newButton(newRect(750, 100, 300, 200), ToPlasma, "PlayButton.png"),
	aboutButton: newButton(newRect(770, 260, 300, 200), ToAbout, "aboutButton.png"),
	levSelButton: newButton(newRect(700, 450, 300, 200), ToLevelSelect, "levSelButton.png"),
	volButton: newVolButton("volButtonOn.png", "volButtonOff.png", newRect(100, 100, 165, 165), SwitchVol, *audioPlayer),
	}
}

func (m MainMenu) Init(g *Game) {

	menuSprites = []GUI{&m.protoStartBg, &m.startBg, &m.startP1, &m.startP2, &m.startP3, &m.startP4, &m.fixedStart, &m.playbutton, &m.aboutButton, &m.levSelButton, &m.volButton}

	state_array = menuSprites

	curr_volBtn = m.volButton

	SwitchVol(g)
}

func (m MainMenu) Update(g *Game) {
	ebiten.SetWindowTitle("Cell Signaling Pathway - Main Menu")
	for _, element := range menuSprites {
		element.update(g)
/*  		switch element.(type) {
		case (Button):
			element.update(g)
		case (VolButton):
			element.update(g)
		default:
			//element.update()
		}  */
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
