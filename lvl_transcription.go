package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

//var currentFrag   = 0

type TranscriptionLevel struct {
	// NUCLEUS SPRITES
	nucleusBg     StillImage
	rna           [5]Transcript
	dna           [5]Template
	temp_tfa      TFA
	rnaPolymerase RNAPolymerase
	rightChoice   CodonChoice
	wrongChoice1  CodonChoice
	wrongChoice2  CodonChoice
	infoButton        InfoPage
	otherToMenuButton Button
	message           string

	// Note to self: when updating DNA image, make the sprite like plasma membrane
	// So it can scroll to the left and show different codons, with bases as separate sprites
	// Also maybe try making RNA with theta and scrolling off to a upper-right diagonal

}

var transcriptionStruct *TranscriptionLevel

func newTranscriptionLevel(g *Game) {
	for x := 0; x < 5; x++ {
		transcriptionStruct.dna[x] = newTemplate("DNA.png", newRect(-50+200*x, 500, 150, 150), template[x], x)}
	for x := 0; x < 5; x++ {
		transcriptionStruct.rna[x] = newTranscript("RNA.png", newRect(0, 200, 150, 150), transcribe(template[x]))
	}
	if len(nucleusSprites) == 0 {
		transcriptionStruct = &TranscriptionLevel{
			nucleusBg: newStillImage("NucleusBg.png", newRect(0, 0, 1250, 750)),

			temp_tfa: newTFA("act_TFA.png", newRect(400, -100, 150, 150), "tfa2"),
			rnaPolymerase: newRNAPolymerase("rnaPolym.png", newRect(-350, 100, 340, 265)),
			rightChoice: newCodonChoice("codonButton.png", newRect(50, 150, 192, 111), transcribe(dna[0].codon)),
			wrongChoice1: newCodonChoice("codonButton.png", newRect(350, 150, 192, 111), randomRNACodon(rightChoice.bases)),
			wrongChoice2: newCodonChoice("codonButton.png", newRect(650, 150, 192, 111), randomRNACodon(rightChoice.bases)),
		
			message: "WWELCOME TO THE NUCLEUS!\n" +
			"Match each codon on the DNA template to the complementary\n" +
			"RNA codon to transcribe a new mRNA molecule!!!",
		}
		transcriptionStruct.infoButton = infoButton
		transcriptionStruct.otherToMenuButton = otherToMenuButton

		nucleusSprites = []GUI{
			&transcriptionStruct.nucleusBg, //&transcriptionStruct.rna,
			//&transcriptionStruct.dna, 
			&transcriptionStruct.temp_tfa,
			&transcriptionStruct.rnaPolymerase, &transcriptionStruct.rightChoice,
			&transcriptionStruct.wrongChoice1, &transcriptionStruct.wrongChoice2,
			&transcriptionStruct.infoButton, &transcriptionStruct.otherToMenuButton,
		}
	}
	g.stateMachine.state = transcriptionStruct
}

func (t *TranscriptionLevel) Init(g *Game) {
	currentFrag = 0
	reset = false
	state_array = plasmaSprites
}

func (t *TranscriptionLevel) Update(g *Game) {
	t.temp_tfa.activate()
	
	t.rnaPolymerase.update(temp_tfa.rect.pos.y)
	if reset {
		t.rightChoice.bases = transcribe(t.dna[currentFrag].codon)
		t.wrongChoice1.bases = randomRNACodon(t.rightChoice.bases)
		t.wrongChoice2.bases = randomRNACodon(t.rightChoice.bases)
		reset = false
	}
	//fmt.Printf("%t\n", dna[currentFrag].is_complete)
	t.rightChoice.update(t.dna[currentFrag])
	//fmt.Printf("%t\n", dna[currentFrag].is_complete)
	if t.dna[currentFrag].is_complete {
		nextDNACodon(g)
	}
}

func (t *TranscriptionLevel) Draw(g *Game, screen *ebiten.Image) {
	nucleusBg.draw(screen)
	for x := 0; x < 5; x++ {
		t.rna[x].draw(screen)
	}
	t.dna[0].draw(screen)

	defaultFont.drawFont(screen, "WELCOME TO THE NUCLEUS! \n Match each codon on the DNA template to the corresponding RNA \n codon to transcribe a new mRNA molecule!!!", 100, 50, color.White)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 300)
	t.rnaPolymerase.draw(screen)
	temp_tfa.draw(screen)
	//codonFont.drawFont(screen, strings.Join(template[0:5], ""), dna[currentFrag].rect.pos.x+300, dna[currentFrag].rect.pos.y, color.Black)
	if currentFrag != -1 {
		codonFont.drawFont(screen, t.dna[currentFrag].codon, t.dna[0].rect.pos.x+500, t.dna[0].rect.pos.y, color.Black)
	}
	t.rightChoice.draw(screen)
	t.wrongChoice1.draw(screen)
	t.wrongChoice2.draw(screen)
	codonFont.drawFont(screen, t.rightChoice.bases, t.rightChoice.rect.pos.x+25, t.rightChoice.rect.pos.y+90, color.Black)
	codonFont.drawFont(screen, t.wrongChoice1.bases, t.wrongChoice1.rect.pos.x+25, t.wrongChoice1.rect.pos.y+90, color.Black)
	codonFont.drawFont(screen, t.wrongChoice2.bases, t.wrongChoice2.rect.pos.x+25, t.wrongChoice2.rect.pos.y+90, color.Black)
/* 	for x := 0; x < currentFrag; x++ {
		codonFont.drawFont(screen, t.rna[x].codon, t.rna[0].rect.pos.x+500+(150*x), t.rna[0].rect.pos.y+140, color.Black)
	} */
	switch currentFrag {
	case 1:
		codonFont.drawFont(screen, t.rna[0].codon, t.rna[0].rect.pos.x+500, t.rna[0].rect.pos.y+140, color.Black)
	case 2:
		codonFont.drawFont(screen, t.rna[0].codon, t.rna[0].rect.pos.x+500, t.rna[0].rect.pos.y+140, color.Black)
		codonFont.drawFont(screen, t.rna[1].codon, t.rna[0].rect.pos.x+650, t.rna[0].rect.pos.y+140, color.Black)
	case 3:
		codonFont.drawFont(screen, t.rna[0].codon, t.rna[0].rect.pos.x+500, t.rna[0].rect.pos.y+140, color.Black)
		codonFont.drawFont(screen, t.rna[1].codon, t.rna[0].rect.pos.x+650, t.rna[0].rect.pos.y+140, color.Black)
		codonFont.drawFont(screen, t.rna[2].codon, t.rna[0].rect.pos.x+800, t.rna[0].rect.pos.y+140, color.Black)
	case 4:
		codonFont.drawFont(screen, t.rna[0].codon, t.rna[0].rect.pos.x+500, t.rna[0].rect.pos.y+140, color.Black)
		codonFont.drawFont(screen, t.rna[1].codon, t.rna[0].rect.pos.x+650, t.rna[0].rect.pos.y+140, color.Black)
		codonFont.drawFont(screen, t.rna[2].codon, t.rna[0].rect.pos.x+800, t.rna[0].rect.pos.y+140, color.Black)
		codonFont.drawFont(screen, t.rna[3].codon, t.rna[0].rect.pos.x+950, t.rna[0].rect.pos.y+140, color.Black)
	default:
		break
	}
}
