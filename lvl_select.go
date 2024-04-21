package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Sprites in Level Selection
type LevelSelection struct {
	levSelBg           StillImage
	levToMenuButton    Button
	levToPlasmaButton  Button
	levToCyto1Button   Button
	levToNucleusButton Button
	levToCyto2Button   Button
}

var levSelStruct *LevelSelection

// Initialize level selection struct and levSelSprites array if not initialized, then set state to levSelStruct
func newLevelSelection(g *Game) {
	if len(g.levSelSprites) == 0 {
		levSelStruct = &LevelSelection{
			levSelBg: newStillImage("levSelBg.png", newRect(0, 0, 1250, 750)),
			levToMenuButton: newButton("menuButton.png", newRect(250, 190, 300, 200), ToMenu),
			levToPlasmaButton: newButton("levToPlasmaBtn.png", newRect(520, 110, 300, 180), ToPlasma),
			levToCyto1Button: newButton("levToCyto1Btn.png", newRect(820, 110, 300, 180), ToCyto1),
			levToNucleusButton: newButton("levToNucleusBtn.png", newRect(520, 285, 300, 180), ToNucleus),
			levToCyto2Button: newButton("levToCyto2Btn.png", newRect(820, 285, 300, 180), ToCyto2),
		}
		g.levSelSprites = []GUI{
			&levSelStruct.levSelBg, &levSelStruct.levToMenuButton, &levSelStruct.levToPlasmaButton, &levSelStruct.levToCyto1Button,
			&levSelStruct.levToNucleusButton, &levSelStruct.levToCyto2Button,	
		}
	}
	g.stateMachine.state = levSelStruct
}

func (l *LevelSelection) Init(g *Game) {
	g.state_array = g.levSelSprites
}

func (l *LevelSelection) Update(g *Game) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		for _, element := range g.levSelSprites {
			element.update(g)
		}
	}
}

func (l *LevelSelection) Draw(g *Game, screen *ebiten.Image) {
	for _, element := range g.levSelSprites {
		element.draw(screen)
	}
}
