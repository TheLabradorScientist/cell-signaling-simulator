package main

import (
	//"log"
	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Sprites in Level Selection
type LevelSelection struct {
	levSelBg           StillImage
	levToPlasmaButton  Button
	levToCyto1Button   Button
	levToNucleusButton Button
	levToCyto2Button   Button
	levToMenuButton    Button
}

var levSelStruct *LevelSelection

// Initialize level selection struct and levSelSprites array if not initialized, then set state to levSelStruct
func newLevelSelection(g *Game) {
	if len(levSelSprites) == 0 {
		levSelStruct = &LevelSelection{
			levSelBg: newStillImage("levSelBg.png", newRect(0, 0, 1250, 750)),
			levToPlasmaButton: newButton("levToPlasmaBtn.png", newRect(520, 110, 300, 180), ToPlasma),
			levToCyto1Button: newButton("levToCyto1Btn.png", newRect(820, 110, 300, 180), ToCyto1),
			levToNucleusButton: newButton("levToNucleusBtn.png", newRect(520, 285, 300, 180), ToNucleus),
			levToCyto2Button: newButton("levToCyto2Btn.png", newRect(820, 285, 300, 180), ToCyto2),
			levToMenuButton: newButton("menuButton.png", newRect(250, 190, 300, 200), ToMenu),
		}
		levSelSprites = []GUI{
			&levSelStruct.levSelBg, &levSelStruct.levToPlasmaButton, &levSelStruct.levToCyto1Button,
			&levSelStruct.levToNucleusButton, &levSelStruct.levToCyto2Button, &levSelStruct.levToMenuButton,	
		}
	}
	g.stateMachine.state = levSelStruct
}

func (l *LevelSelection) Init(g *Game) {
	//ebiten.SetWindowTitle("Cell Signaling Pathway - Level Selection")
	state_array = levSelSprites
}

func (l *LevelSelection) Update(g *Game) {
	for _, element := range levSelSprites {
		element.update(g)
	}
}

func (l *LevelSelection) Draw(g *Game, screen *ebiten.Image) {
	for _, element := range levSelSprites {
		element.draw(screen)
	}
}
