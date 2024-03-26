package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	currentFrag = 0
	rna         [5]Transcript
	dna         [5]Template
	spots       = [3]int{350, 650, 950}
	DNAbases 	[15]Nucleobase
	RNAbases	[15]Nucleobase
)

type TranscriptionLevel struct {
	// TRANSCRIPTION SPRITES
	nucleusBg         StillImage
	temp_tfa          TFA
	rnaPolymerase     RNAPolymerase
	rightChoice       CodonChoice
	wrongChoice1      CodonChoice
	wrongChoice2      CodonChoice
	infoButton        InfoPage
	otherToMenuButton Button
	message           string

	// Note to self: when updating DNA image, make the sprite like plasma membrane
	// So it can scroll to the left and show different codons, with bases as separate sprites
	// Also maybe try making RNA with theta and scrolling off to a upper-right diagonal

}

var transcriptionStruct *TranscriptionLevel

func newTranscriptionLevel(g *Game) {
	if len(g.transcriptionSprites) == 0 {
		transcriptionStruct = &TranscriptionLevel{
			nucleusBg: newStillImage("NucleusBg.png", newRect(0, 0, 1250, 750)),

			temp_tfa:      newTFA("inact_TFA.png", "act_TFA.png", newRect(450, -100, 150, 150), "tfa2"),
			rnaPolymerase: newRNAPolymerase("rnaPolym.png", newRect(-400, 100, 340, 265)),

			message: "WELCOME TO THE NUCLEUS! \n" +
				"Drag the complementary RNA codon \n" +
				"to RNA Polymerase to transcribe \n" +
				"a new mRNA molecule!!!",
		}

		transcriptionStruct.rightChoice = newCodonChoice("codonButton.png", newRect(100, 200, 192, 111), transcribe(dna[0].codon))
		transcriptionStruct.wrongChoice1 = newCodonChoice("codonButton.png", newRect(400, 200, 192, 111), randomRNACodon(transcriptionStruct.rightChoice.bases))
		transcriptionStruct.wrongChoice2 = newCodonChoice("codonButton.png", newRect(700, 200, 192, 111), randomRNACodon(transcriptionStruct.rightChoice.bases))
		transcriptionStruct.infoButton = infoButton
		transcriptionStruct.otherToMenuButton = otherToMenuButton

		g.transcriptionSprites = []GUI{
			&transcriptionStruct.nucleusBg, &transcriptionStruct.temp_tfa,
			&transcriptionStruct.rnaPolymerase, &transcriptionStruct.rightChoice,
			&transcriptionStruct.wrongChoice1, &transcriptionStruct.wrongChoice2,
			&transcriptionStruct.otherToMenuButton, &transcriptionStruct.infoButton,
		}
	}
	g.stateMachine.state = transcriptionStruct
}

func (t *TranscriptionLevel) Init(g *Game) {
	currentFrag = 0
	for x := 0; x < len(RNAbases); x++ {
		base := string(rna[x/3].codon[x%3])
		posX := 125+rna[4].rect.pos.x + (50*x)
		posY := rna[4].rect.pos.y+275 - (15*x)
		RNAbases[x] = newNucleobase(base, newRect(posX, posY, 65, 150), x, false)
	}
	t.ResetChoices()
	g.state_array = g.transcriptionSprites
	t.temp_tfa.activate()
}

func (t *TranscriptionLevel) ResetChoices() {
	curr := &dna[currentFrag]
	rand.Shuffle(len(spots), func(i, j int) { spots[i], spots[j] = spots[j], spots[i] })
	t.rightChoice.reset(0, 600, transcribe(curr.codon))
	t.wrongChoice1.reset(1, 600, randomRNACodon(t.rightChoice.bases))
	t.wrongChoice2.reset(2, 600, randomRNACodon(t.rightChoice.bases))
	reset = false
}

func (t *TranscriptionLevel) Update(g *Game) {
	t.otherToMenuButton.update(g)
	t.infoButton.update()
	t.temp_tfa.update()
	t.rnaPolymerase.update(t.temp_tfa.rect.pos.y)

	curr := &dna[currentFrag]

	if reset {
		t.ResetChoices()
	}

	//fmt.Printf("%t\n", dna[currentFrag].is_complete)
	t.rightChoice.update(curr)
	t.wrongChoice1.update(curr)
	t.wrongChoice2.update(curr)

	if curr.is_complete {
		nextDNACodon(g)
	}

	for i, base := range RNAbases {
		base.update()
		RNAbases[i] = base
	}

}

func (t *TranscriptionLevel) Draw(g *Game, screen *ebiten.Image) {
	t.nucleusBg.draw(screen)

	rna[currentFrag].draw(screen)
	t.rnaPolymerase.draw(screen)
	dna[0].draw(screen)
	t.temp_tfa.draw(screen)
	//codonFont.drawFont(screen, strings.Join(template[0:5], ""), dna[currentFrag].rect.pos.x+300, dna[currentFrag].rect.pos.y, color.Black)

	t.rightChoice.draw(screen)
	t.wrongChoice1.draw(screen)
	t.wrongChoice2.draw(screen)

	for z := 0; z < currentFrag*3; z++ {
		RNAbases[z].draw(screen)
	}

	if currentFrag != -1 {
		codonFont.drawFont(screen, dna[currentFrag].codon, dna[0].rect.pos.x+500, dna[0].rect.pos.y+50, color.Black)
	}

	defaultFont.drawFont(screen, t.message, 75, 50, color.Black)

	t.otherToMenuButton.draw(screen)

	t.infoButton.draw(screen)
}
