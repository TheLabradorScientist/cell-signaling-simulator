package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type TransductionLevel struct {
	// CYTO 1 SPRITES
	protoCytoBg_1     StillImage
	cytoBg_1          Parallax
	cytoNuc_1         Parallax
	tk1               Kinase
	tk2               Kinase
	tfa               TFA
	infoButton        InfoPage
	otherToMenuButton Button
	message           string
}

var transductionStruct *TransductionLevel

func newTransductionLevel(g *Game) {
	if len(g.transductionSprites) == 0 {
		transductionStruct = &TransductionLevel{
			protoCytoBg_1: newStillImage("CytoBg1.png", newRect(0, 0, 1250, 750)),
			cytoBg_1:      newParallax("ParallaxCyto1.png", newRect(100, 100, 1250, 750), 4),
			cytoNuc_1:     newParallax("ParallaxCyto1.5.png", newRect(100, 100, 1250, 750), 3),

			tk1: newKinase("inact_TK1.png", "act_TK1.png", newRect(500, -100, 150, 150), "tk1"),
			tk2: newKinase("inact_TK2.png", "act_TK2.png", newRect(250, 175, 150, 150), "tk2"),
			tfa: newTFA("inact_TFA.png", "act_TFA.png", newRect(700, 500, 150, 150), "tfa1"),

			message: 
				"WELCOME TO THE CYTOPLASM! \n" +
				"Click when the kinases overlap to \n" +
				"follow the phosphorylation cascade!!",
		}
		transductionStruct.infoButton = infoButton
		transductionStruct.otherToMenuButton = otherToMenuButton

		g.transductionSprites = []GUI{
			&transductionStruct.protoCytoBg_1, &transductionStruct.cytoBg_1,
			&transductionStruct.cytoNuc_1, &transductionStruct.tk1, &transductionStruct.tk2,
			&transductionStruct.tfa, &transductionStruct.otherToMenuButton, &transductionStruct.infoButton,
		}
	}
	g.stateMachine.state = transductionStruct
}

func (t *TransductionLevel) Init(g *Game) {
	g.state_array = g.transductionSprites
	t.tk1.activate()
}

func (t *TransductionLevel) Update(g *Game) {
	for _, element := range g.transductionSprites {
		element.update(g)
	}
	if t.tk1.is_clicked_on {
		t.tk2.activate()
		t.tk1.is_clicked_on = false
	}
	if t.tk2.is_clicked_on {
		t.tfa.activate()
		t.tk2.is_clicked_on = false
	}
	if t.tfa.rect.pos.y > screenHeight {
		ToNucleus(g)
	}
}

func (t *TransductionLevel) Draw(g *Game, screen *ebiten.Image) {
	for _, element := range g.transductionSprites {
		element.draw(screen)
	}
	defaultFont.drawFont(screen, t.message, 100, 50, color.Black)
}
