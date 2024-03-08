//MUST UPDATE ALL FUNCTIONS WITH NEW CODE

package main

import (
	"image/color"
	"log"
	//"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Level1 struct {
	plasmaBg   *ebiten.Image
	seedSignal int
	signal     Signal
	receptorA  Receptor
	receptorB  Receptor
	receptorC  Receptor
	receptorD  Receptor
	temp_tk1A  Kinase
	temp_tk1B  Kinase
	temp_tk1C  Kinase
	temp_tk1D  Kinase
}

//func newLevel1(g *Game) {
//	g.stateMachine.state = Level1{
//		seedSignal: rand.Intn(4) + 1,
//	}
//}

func (l Level1) Init() {
	l.plasmaBg, _, err = ebitenutil.NewImageFromFile(loadFile("PlasmaBg.png"))
	if err != nil {
		log.Fatal(err)
	}

	switch l.seedSignal {
	case 1:
		l.signal = newSignal("signalA.png", newRect(500, 100, 75, 75))
		l.signal.signalType = "signalA"
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ACT"}
	case 2:
		l.signal = newSignal("signalB.png", newRect(500, 100, 75, 75))
		l.signal.signalType = "signalB"
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATT"}
	case 3:
		l.signal = newSignal("signalC.png", newRect(500, 100, 75, 75))
		l.signal.signalType = "signalC"
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATC"}
	case 4:
		l.signal = newSignal("signalD.png", newRect(500, 100, 75, 75))
		l.signal.signalType = "signalD"
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATT"}
		// PLACEHOLDER IN CASE WE DO NOT GET TIME TO CODE RANDOM CODONS
		//template = [5]string{"TAC", "GTC", "CGG", "ACA", "ACT"}
	}

	l.receptorA = newReceptor("inact_receptorA.png", newRect(50, 400, 100, 150), "receptorA")
	l.receptorB = newReceptor("inact_receptorB.png", newRect(350, 400, 100, 150), "receptorB")
	l.receptorC = newReceptor("inact_l.receptorC.png", newRect(650, 400, 100, 150), "l.receptorC")
	l.receptorD = newReceptor("inact_receptorD.png", newRect(950, 400, 100, 150), "receptorD")

	l.temp_tk1A = newKinase("inact_TK1.png", newRect(50, 600, 150, 150), "temp_tk1")
	l.temp_tk1B = newKinase("inact_TK1.png", newRect(350, 600, 150, 150), "temp_tk1")
	l.temp_tk1C = newKinase("inact_TK1.png", newRect(650, 600, 150, 150), "temp_tk1")
	l.temp_tk1D = newKinase("inact_TK1.png", newRect(950, 600, 150, 150), "temp_tk1")
}

func (l Level1) Update(g *Game) {
	ebiten.SetWindowTitle("Cell Signaling Pathway - Signal Reception")
	ebiten.SetWindowSize(screenWidth, screenHeight)
	l.signal.update()
	l.receptorA.update()
	l.receptorB.update()
	l.receptorC.update()
	l.receptorD.update()
	l.temp_tk1A.update(l.temp_tk1B.rect)
	l.temp_tk1B.update(l.temp_tk1C.rect)
	l.temp_tk1C.update(l.temp_tk1D.rect)
	l.temp_tk1D.update(l.temp_tk1A.rect)
	if l.receptorA.is_touching_signal {
		if matchSR(l.signal.signalType, l.receptorA.receptorType) {
			l.receptorA.animate("act_receptorA.png")
			l.temp_tk1A.activate()
		}
	}
	if l.receptorB.is_touching_signal {
		if matchSR(l.signal.signalType, l.receptorB.receptorType) {
			l.receptorB.animate("act_receptorB.png")
			l.temp_tk1B.activate()
		}
	}
	if l.receptorC.is_touching_signal {
		if matchSR(l.signal.signalType, l.receptorC.receptorType) {
			l.receptorC.animate("act_l.receptorC.png")
			l.temp_tk1C.activate()
		}
	}
	if l.receptorD.is_touching_signal {
		if matchSR(l.signal.signalType, l.receptorD.receptorType) {
			l.receptorD.animate("act_receptorD.png")
			l.temp_tk1D.activate()
		}
	}
	if l.temp_tk1A.rect.pos.y >= 750 || l.temp_tk1B.rect.pos.y >= 750 || l.temp_tk1C.rect.pos.y >= 750 || l.temp_tk1D.rect.pos.y >= 750 {
		ToCyto1(g)
	}
}

func (l Level1) Draw(g *Game, screen *ebiten.Image) {
	screen.DrawImage(l.plasmaBg, nil)
	l.receptorA.draw(screen)
	l.receptorB.draw(screen)
	l.receptorC.draw(screen)
	l.receptorD.draw(screen)
	l.signal.draw(screen)
	l.temp_tk1A.draw(screen)
	l.temp_tk1B.draw(screen)
	l.temp_tk1C.draw(screen)
	l.temp_tk1D.draw(screen)
	defaultFont.drawFont(screen, "WELCOME TO THE PLASMA MEMBRANE! \n Drag the signal to the matching receptor \n to enter the cell!", 100, 50, color.White)
}
