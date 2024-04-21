package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Sprites in Main Menu
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

var menuStruct *MainMenu

// Initialize menu struct and menuSprites array if not initialized, then set state to menuStruct
func newMainMenu(g *Game) {
	if len(g.menuSprites) == 0 {
		menuStruct = &MainMenu{
			protoStartBg: newStillImage("MenuBg.png", newRect(0, 0, 1250, 750)),
			startBg:      newParallax("StartBg.png", newRect(0, 0, 1250, 750), 5),
			startP1:      newParallax("parallax-Start2.png", newRect(0, 0, 1250, 750), 4),
			startP2:      newParallax("parallax-Start3.png", newRect(0, 0, 1250, 750), 3),
			startP3:      newParallax("parallax-Start4.png", newRect(0, 0, 1250, 750), 2),
			startP4:      newParallax("parallax-Start5.png", newRect(0, 0, 1250, 750), 1),
			fixedStart:   newStillImage("fixed-Start.png", newRect(0, 0, 1250, 750)),
			playbutton:   newButton("PlayButton.png", newRect(750, 100, 300, 200), ToPlasma),
			aboutButton:  newButton("aboutButton.png", newRect(770, 260, 300, 200), ToAbout),
			levSelButton: newButton("levSelButton.png", newRect(700, 450, 300, 200), ToLevelSelect),
		}
		menuStruct.volButton = newVolButton("volButtonOn.png", newRect(100, 100, 165, 165), menuStruct.volButton.Toggle, *audioPlayer)
		
		g.menuSprites = []GUI{
			&menuStruct.protoStartBg, &menuStruct.startBg, &menuStruct.startP1, &menuStruct.startP2,
			&menuStruct.startP3, &menuStruct.startP4, &menuStruct.fixedStart, &menuStruct.playbutton,
			&menuStruct.aboutButton, &menuStruct.levSelButton, &menuStruct.volButton}
	}
	
	g.stateMachine.state = menuStruct
}

func (m *MainMenu) Init(g *Game) {
	g.state_array = g.menuSprites
}

func (m *MainMenu) Update(g *Game) {
	for _, element := range g.menuSprites {
		element.update(g)
	}
}

func (m *MainMenu) Draw(g *Game, screen *ebiten.Image) {
	for _, element := range g.menuSprites {
		element.draw(screen)
	}
}
