package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	currentFrag   = 0
	rna		[5]Transcript
	dna		[5]Template
)

type TranscriptionLevel struct {
	// NUCLEUS SPRITES
	nucleusBg     StillImage
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
	if len(nucleusSprites) == 0 {
		transcriptionStruct = &TranscriptionLevel{
			nucleusBg: newStillImage("NucleusBg.png", newRect(0, 0, 1250, 750)),

			temp_tfa: newTFA("act_TFA.png", newRect(400, -100, 150, 150), "tfa2"),
			rnaPolymerase: newRNAPolymerase("rnaPolym.png", newRect(-350, 100, 340, 265)),
		
			message: "WELCOME TO THE NUCLEUS!\n" +
			"Match each codon on the DNA template to the complementary\n" +
			"RNA codon to transcribe a new mRNA molecule!!!",
		}
		for x := 0; x < 5; x++ {
			dna[x] = newTemplate("DNA.png", newRect(-50+200*x, 500, 150, 150), template[x], x)}
		for x := 0; x < 5; x++ {
			rna[x] = newTranscript("RNA.png", newRect(0, 200, 150, 150), transcribe(template[x]))
		}
		transcriptionStruct.rightChoice = newCodonChoice("codonButton.png", newRect(50, 150, 192, 111), transcribe(dna[0].codon))
		transcriptionStruct.wrongChoice1 = newCodonChoice("codonButton.png", newRect(350, 150, 192, 111), randomRNACodon(transcriptionStruct.rightChoice.bases))
		transcriptionStruct.wrongChoice2 = newCodonChoice("codonButton.png", newRect(650, 150, 192, 111), randomRNACodon(transcriptionStruct.rightChoice.bases))
		transcriptionStruct.infoButton = infoButton
		transcriptionStruct.otherToMenuButton = otherToMenuButton

		nucleusSprites = []GUI{
			&transcriptionStruct.nucleusBg, &transcriptionStruct.temp_tfa,
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
	state_array = nucleusSprites
}

func (t *TranscriptionLevel) Update(g *Game) {
		t.otherToMenuButton.update(g)
		t.infoButton.update()
		t.temp_tfa.activate()
		t.temp_tfa.update()
		t.rnaPolymerase.update(t.temp_tfa.rect.pos.y)

		curr := &dna[currentFrag]

		if reset {
			t.rightChoice.bases = transcribe(curr.codon)
			t.wrongChoice1.bases = randomRNACodon(t.rightChoice.bases)
			t.wrongChoice2.bases = randomRNACodon(t.rightChoice.bases)
			reset = false
		}

		//fmt.Printf("%t\n", dna[currentFrag].is_complete)
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			t.rightChoice.update(curr)
		}
		
		if curr.is_complete {
			nextDNACodon(g)
		}

}

func (t *TranscriptionLevel) Draw(g *Game, screen *ebiten.Image) {
	t.nucleusBg.draw(screen)
	for x := 0; x < 5; x++ {
		rna[x].draw(screen)
	} 

	dna[0].draw(screen)

	t.rnaPolymerase.draw(screen)
	t.temp_tfa.draw(screen)
	//codonFont.drawFont(screen, strings.Join(template[0:5], ""), dna[currentFrag].rect.pos.x+300, dna[currentFrag].rect.pos.y, color.Black)

	t.rightChoice.draw(screen)
	t.wrongChoice1.draw(screen)
	t.wrongChoice2.draw(screen)

	t.infoButton.draw(screen)
	t.otherToMenuButton.draw(screen)

	for x := 0; x < currentFrag; x++ {
		codonFont.drawFont(screen, rna[x].codon, rna[0].rect.pos.x+500+(150*x), rna[0].rect.pos.y+140, color.Black)
	}

	if currentFrag != -1 {
		codonFont.drawFont(screen, dna[currentFrag].codon, dna[0].rect.pos.x+500, dna[0].rect.pos.y, color.Black)
	}

	defaultFont.drawFont(screen, t.message, 100, 50, color.White)

}
