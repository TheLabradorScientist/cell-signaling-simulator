//MUST UPDATE ALL FUNCTIONS WITH NEW CODE

package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Level2 struct {
	cytoBg_1 *ebiten.Image
	tk1      Kinase
	tk2      Kinase
	tfa      TFA
}

func newLevel2(g *Game) {
//	g.stateMachine.state = Level2{}
}

func (l Level2) Init() {
	l.cytoBg_1, _, err = ebitenutil.NewImageFromFile(loadFile("CytoBg1.png"))
	if err != nil {
		log.Fatal(err)
	}
	l.tk1 = newKinase("act_TK1.png", newRect(500, -100, 150, 150), "tk1")
	l.tk2 = newKinase("inact_TK2.png", newRect(250, 175, 150, 150), "tk2")
	l.tfa = newTFA("inact_TFA.png", newRect(700, 500, 150, 150), "tfa1")
}

func (l Level2) Update(g *Game) {
	//ebiten.SetWindowTitle("Cell Signaling Pathway - Signal Transduction")
	ebiten.SetWindowSize(screenWidth, screenHeight)
	l.tk1.activate()
	l.tk1.update(l.tk2.rect)
	l.tk2.update(l.tfa.rect)
	l.tfa.update()
	if l.tk1.is_clicked_on {
		l.tk2.activate()
		l.tk1.is_clicked_on = false
	}
	if l.tk2.is_clicked_on {
		l.tfa.activate()
		l.tk2.is_clicked_on = false
	}
	if l.tfa.rect.pos.y > 750 {
		ToNucleus(g)
	}
}

func (l Level2) Draw(g *Game, screen *ebiten.Image) {
	screen.DrawImage(l.cytoBg_1, nil)
	l.tk1.draw(screen)
	l.tk2.draw(screen)
	l.tfa.draw(screen)
	defaultFont.drawFont(screen, "WELCOME TO THE CYTOPLASM! \n Click when each kinase overlaps to follow \n the phosphorylation cascade!!", 100, 50, color.Black)
}
