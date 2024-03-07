//MUST UPDATE ALL FUNCTIONS WITH NEW CODE

package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Level4 struct {
	cytoBg_2     *ebiten.Image
	ribosome     Ribosome
	rightChoice  CodonChoice
	wrongChoice1 CodonChoice
	wrongChoice2 CodonChoice
	dna          [5]Template
	mrna         [5]Template
	protein      [5]Transcript
	mrna_ptr     int
	rightTrna    CodonChoice
	wrongTrna1   CodonChoice
	wrongTrna2   CodonChoice
	reset        bool
}

func newLevel4(g *Game) {
	g.stateMachine.state = Level4{}
}

func (l Level4) Init() {
	l.reset = false
	l.cytoBg_2, _, err = ebitenutil.NewImageFromFile(loadFile("CytoBg2.png"))
	if err != nil {
		log.Fatal(err)
	}

	for x := 0; x < 5; x++ {
		l.mrna[x] = newTemplate("RNA.png", newRect(0, 400, 150, 150), transcribe(l.dna[x].codon), x)
	}
	for x := 0; x < 5; x++ {
		l.protein[x] = newTranscript("aminoAcid.png", newRect(50+(150*x), 400, 150, 150), translate(l.mrna[x].codon))
	}

	l.rightChoice = newCodonChoice("codonButton.png", newRect(50, 150, 192, 111), transcribe(dna[0].codon))
	l.wrongChoice1 = newCodonChoice("codonButton.png", newRect(350, 150, 192, 111), randomRNACodon(rightChoice.bases))
	l.wrongChoice2 = newCodonChoice("codonButton.png", newRect(650, 150, 192, 111), randomRNACodon(rightChoice.bases))

	l.ribosome = newRibosome("ribosome.png", newRect(0, 300, 404, 367))

	l.mrna_ptr = 0

	l.rightTrna = newCodonChoice("codonButton.png", newRect(50, 150, 192, 111), translate(mrna[0].codon))
	l.wrongTrna1 = newCodonChoice("codonButton.png", newRect(350, 150, 192, 111), translate(randomRNACodon(rightTrna.bases)))
	l.wrongTrna2 = newCodonChoice("codonButton.png", newRect(650, 150, 192, 111), translate(randomRNACodon(rightTrna.bases)))
}

func (l Level4) Update(g *Game) {
	ebiten.SetWindowTitle("Cell Signaling Pathway - Translation")
	if l.reset {
		l.rightTrna.bases = translate(l.mrna[mrna_ptr].codon)
		l.wrongTrna1.bases = translate(randomRNACodon(l.rightTrna.bases))
		l.wrongTrna2.bases = translate(randomRNACodon(l.rightTrna.bases))
		l.reset = false
	}
	l.mrna[l.mrna_ptr].is_complete = l.rightTrna.update2(l.mrna[l.mrna_ptr].codon)

	if l.mrna[l.mrna_ptr].is_complete {
		nextMRNACodon(g)
	}
}

func (l Level4) Draw(g *Game, screen *ebiten.Image) {
	screen.DrawImage(l.cytoBg_2, nil)
	// for x := 0; x < 5; x++ {
	// 	protein[x].draw(screen)
	// }

	l.mrna[0].draw(screen)
	l.ribosome.draw(screen)

	if l.mrna_ptr != -1 {
		g.codonFont.drawFont(screen, l.mrna[l.mrna_ptr].codon, l.mrna[0].rect.pos.x+500, l.mrna[0].rect.pos.y+200, color.Black)
	}

	l.rightTrna.draw(screen)
	l.wrongTrna1.draw(screen)
	l.wrongTrna2.draw(screen)

	g.codonFont.drawFont(screen, l.rightTrna.bases, l.rightTrna.rect.pos.x+25, l.rightTrna.rect.pos.y+90, color.Black)
	g.codonFont.drawFont(screen, l.wrongTrna1.bases, l.wrongTrna1.rect.pos.x+25, l.wrongTrna1.rect.pos.y+90, color.Black)
	g.codonFont.drawFont(screen, l.wrongTrna2.bases, l.wrongTrna2.rect.pos.x+25, l.wrongTrna2.rect.pos.y+90, color.Black)

	g.defaultFont.drawFont(screen, "FINALLY, BACK TO THE CYTOPLASM! \n Match each codon from your mRNA template \n to its corresponding amino acid to synthesize your protein!!!!", 100, 50, color.Black)

	switch l.mrna_ptr {
	case 1:
		l.protein[0].draw(screen)
		g.codonFont.drawFont(screen, l.protein[0].codon, l.protein[0].rect.pos.x, l.protein[0].rect.pos.y, color.Black)
	case 2:
		l.protein[0].draw(screen)
		l.protein[1].draw(screen)
		g.codonFont.drawFont(screen, l.protein[0].codon, l.protein[0].rect.pos.x, l.protein[0].rect.pos.y, color.Black)
		g.codonFont.drawFont(screen, l.protein[1].codon, l.protein[1].rect.pos.x, l.protein[1].rect.pos.y, color.Black)
	case 3:
		l.protein[0].draw(screen)
		l.protein[1].draw(screen)
		l.protein[2].draw(screen)
		g.codonFont.drawFont(screen, l.protein[0].codon, l.protein[0].rect.pos.x, l.protein[0].rect.pos.y, color.Black)
		g.codonFont.drawFont(screen, l.protein[1].codon, l.protein[1].rect.pos.x, l.protein[1].rect.pos.y, color.Black)
		g.codonFont.drawFont(screen, l.protein[2].codon, l.protein[2].rect.pos.x, l.protein[2].rect.pos.y, color.Black)
	case 4:
		l.protein[0].draw(screen)
		l.protein[1].draw(screen)
		l.protein[2].draw(screen)
		l.protein[3].draw(screen)
		g.codonFont.drawFont(screen, l.protein[0].codon, l.protein[0].rect.pos.x, l.protein[0].rect.pos.y, color.Black)
		g.codonFont.drawFont(screen, l.protein[1].codon, l.protein[1].rect.pos.x, l.protein[1].rect.pos.y, color.Black)
		g.codonFont.drawFont(screen, l.protein[2].codon, l.protein[2].rect.pos.x, l.protein[2].rect.pos.y, color.Black)
		g.codonFont.drawFont(screen, l.protein[3].codon, l.protein[3].rect.pos.x, l.protein[3].rect.pos.y, color.Black)
	default:
		break
	}
}
