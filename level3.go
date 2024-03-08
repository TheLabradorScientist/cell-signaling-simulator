//MUST UPDATE ALL FUNCTIONS WITH NEW CODE

package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Level3 struct {
	nucleusBg     *ebiten.Image
	temp_tfa      TFA
	template      [5]string
	rna           [5]Transcript
	dna           [5]Template
	currentFrag   int
	rnaPolymerase RNAPolymerase
	reset         bool
	rightChoice   CodonChoice
	wrongChoice1  CodonChoice
	wrongChoice2  CodonChoice
}

func newLevel3(g *Game) {
//	g.stateMachine.state = Level3{}
}

func (l Level3) Init() {
	l.nucleusBg, _, err = ebitenutil.NewImageFromFile(loadFile("NucleusBg.png"))
	if err != nil {
		log.Fatal(err)
	}

	for x := 0; x < 5; x++ {
		l.dna[x] = newTemplate("DNA.png", newRect(-50+200*x, 500, 150, 150), l.template[x], x)
	}
	for x := 0; x < 5; x++ {
		l.rna[x] = newTranscript("RNA.png", newRect(0, 200, 150, 150), transcribe(l.template[x]))
	}

	temp_tfa = newTFA("act_TFA.png", newRect(400, -100, 150, 150), "tfa2")

	l.currentFrag = 0
	l.reset = false
}

func (l Level3) Update(g *Game) {
	ebiten.SetWindowTitle("Cell Signaling Pathway - Transcription")
	temp_tfa.activate()
	temp_tfa.update()
	l.rnaPolymerase.update(temp_tfa.rect.pos.y)
	if l.reset {
		l.rightChoice.bases = transcribe(l.dna[l.currentFrag].codon)
		l.wrongChoice1.bases = randomRNACodon(l.rightChoice.bases)
		l.wrongChoice2.bases = randomRNACodon(l.rightChoice.bases)
		l.reset = false
	}
	//fmt.Printf("%t\n", dna[currentFrag].is_complete)
	l.dna[l.currentFrag].is_complete = l.rightChoice.update1(l.dna[l.currentFrag].codon)
	//fmt.Printf("%t\n", dna[currentFrag].is_complete)
	if l.dna[l.currentFrag].is_complete {
		nextDNACodon(g)
	}
}

func (l Level3) Draw(g *Game, screen *ebiten.Image) {
	screen.DrawImage(l.nucleusBg, nil)
	for x := 0; x < 5; x++ {
		l.rna[x].draw(screen)
	}
	l.dna[0].draw(screen)

	defaultFont.drawFont(screen, "WELCOME TO THE NUCLEUS! \n Match each codon on the DNA template to the corresponding RNA \n codon to transcribe a new mRNA molecule!!!", 100, 50, color.White)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 300)
	l.rnaPolymerase.draw(screen)
	temp_tfa.draw(screen)
	//codonFont.drawFont(screen, strings.Join(template[0:5], ""), dna[currentFrag].rect.pos.x+300, dna[currentFrag].rect.pos.y, color.Black)
	if l.currentFrag != -1 {
		codonFont.drawFont(screen, l.dna[l.currentFrag].codon, l.dna[0].rect.pos.x+500, l.dna[0].rect.pos.y, color.Black)
	}
	l.rightChoice.draw(screen)
	l.wrongChoice1.draw(screen)
	l.wrongChoice2.draw(screen)
	codonFont.drawFont(screen, l.rightChoice.bases, l.rightChoice.rect.pos.x+25, l.rightChoice.rect.pos.y+90, color.Black)
	codonFont.drawFont(screen, l.wrongChoice1.bases, l.wrongChoice1.rect.pos.x+25, l.wrongChoice1.rect.pos.y+90, color.Black)
	codonFont.drawFont(screen, l.wrongChoice2.bases, l.wrongChoice2.rect.pos.x+25, l.wrongChoice2.rect.pos.y+90, color.Black)
	switch l.currentFrag {
	case 1:
		codonFont.drawFont(screen, l.rna[0].codon, l.rna[0].rect.pos.x+500, l.rna[0].rect.pos.y+140, color.Black)
	case 2:
		codonFont.drawFont(screen, l.rna[0].codon, l.rna[0].rect.pos.x+500, l.rna[0].rect.pos.y+140, color.Black)
		codonFont.drawFont(screen, l.rna[1].codon, l.rna[0].rect.pos.x+650, l.rna[0].rect.pos.y+140, color.Black)
	case 3:
		codonFont.drawFont(screen, l.rna[0].codon, l.rna[0].rect.pos.x+500, l.rna[0].rect.pos.y+140, color.Black)
		codonFont.drawFont(screen, l.rna[1].codon, l.rna[0].rect.pos.x+650, l.rna[0].rect.pos.y+140, color.Black)
		codonFont.drawFont(screen, l.rna[2].codon, l.rna[0].rect.pos.x+800, l.rna[0].rect.pos.y+140, color.Black)
	case 4:
		codonFont.drawFont(screen, l.rna[0].codon, l.rna[0].rect.pos.x+500, l.rna[0].rect.pos.y+140, color.Black)
		codonFont.drawFont(screen, l.rna[1].codon, l.rna[0].rect.pos.x+650, l.rna[0].rect.pos.y+140, color.Black)
		codonFont.drawFont(screen, l.rna[2].codon, l.rna[0].rect.pos.x+800, l.rna[0].rect.pos.y+140, color.Black)
		codonFont.drawFont(screen, l.rna[3].codon, l.rna[0].rect.pos.x+950, l.rna[0].rect.pos.y+140, color.Black)
	default:
		break
	}
}
