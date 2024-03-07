//MUST UPDATE ALL FUNCTIONS WITH NEW CODE

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type LevelSelection struct {
	levToPlasmaButton  Button
	levToCyto1Button   Button
	levToNucleusButton Button
	levToCyto2Button   Button
	levToMenuButton    Button
	levSelBg           *ebiten.Image
}

func newLevelSelection(g *Game) {
	g.stateMachine.state = LevelSelection{
		levToPlasmaButton:  newButton("levToPlasmaBtn.png", newRect(400, 125, 232, 129), ToPlasma),
		levToCyto1Button:   newButton("levToCyto1Btn.png", newRect(750, 125, 232, 129), ToCyto1),
		levToNucleusButton: newButton("levToNucleusBtn.png", newRect(400, 250, 232, 129), ToNucleus),
		levToCyto2Button:   newButton("levToCyto2Btn.png", newRect(750, 250, 232, 129), ToCyto2),
		levToMenuButton:    newButton("menuButton.png", newRect(775, 400, 232, 129), ToMenu),
		levSelBg:           nil,
	}
}

func (l LevelSelection) Init() {
	l.levSelBg, _, err = ebitenutil.NewImageFromFile(loadFile("levSelBg.png"))
	if err != nil {
		log.Fatal(err)
	}
}

func (l LevelSelection) Update(g *Game) {
	ebiten.SetWindowTitle("Cell Signaling Pathway - Level Selection")
	ebiten.SetWindowSize(screenWidth, screenHeight)
	l.levToPlasmaButton.update(g)
	l.levToCyto1Button.update(g)
	l.levToNucleusButton.update(g)
	l.levToCyto2Button.update(g)
	l.levToMenuButton.update(g)
}

func (l LevelSelection) Draw(g *Game, screen *ebiten.Image) {
	screen.DrawImage(l.levSelBg, nil)
	l.levToPlasmaButton.draw(screen)
	l.levToCyto1Button.draw(screen)
	l.levToNucleusButton.draw(screen)
	l.levToCyto2Button.draw(screen)
	l.levToMenuButton.draw(screen)
}
