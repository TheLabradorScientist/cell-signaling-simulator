package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	currentFrag = 0
	currBase    = 0
	rna         [5]Transcript
	dna         [5]Template
	spots       = [3]int{350, 650, 950}
	baseSpots   = [5]int{100, 250, 400, 550, 700}
)

type TranscriptionLevel struct {
	// TRANSCRIPTION SPRITES
	nucleusBg     		StillImage
	temp_tfa      		TFA
	rnaPolymerase 		RNAPolymerase
	RNA           		[6]Transcript
	DNA           		[6]Template
	DNAbases      		[16]Nucleobase
	origRNAbases  		[16]Nucleobase // Dummy list containing positions and bases, accessed by RNA bases
	RNAbases      		[16]Nucleobase // List that is actually drawn onto screen and updated.
	A_Choice      		BaseChoice
	T_Choice      		BaseChoice
	G_Choice      		BaseChoice
	C_Choice      		BaseChoice
	U_Choice      		BaseChoice
	infoButton        	InfoPage
	otherToMenuButton 	Button
	message           	string

	// Note to self: when updating DNA image, make the sprite like plasma membrane
	// So it can scroll to the left and show different codons, with bases as separate sprites
	// Also maybe try making RNA with theta and scrolling off to a upper-right diagonal

}

var transcriptionStruct *TranscriptionLevel

func newTranscriptionLevel(g *Game) {
	if len(g.transcriptionSprites) == 0 {
		transcriptionStruct = &TranscriptionLevel{
			nucleusBg:     newStillImage("NucleusBg.png", newRect(0, 0, 1250, 750)),
			
			A_Choice:      newBaseChoice("empty.png", newRect(100, 550, 80, 180), "A"),
			T_Choice:      newBaseChoice("empty.png", newRect(250, 550, 80, 180), "T"),
			G_Choice:      newBaseChoice("empty.png", newRect(400, 550, 80, 180), "G"),
			C_Choice:      newBaseChoice("empty.png", newRect(550, 550, 80, 180), "C"),
			U_Choice:      newBaseChoice("empty.png", newRect(700, 550, 80, 180), "U"),

			temp_tfa:      newTFA("inact_TFA.png", "act_TFA.png", newRect(420, -100, 150, 150), "tfa2"),
			rnaPolymerase: newRNAPolymerase("rnaPolym.png", newRect(-400, 100, 340, 265)),

			infoButton:    		infoButton,
			otherToMenuButton: 	otherToMenuButton,

			message: "WELCOME TO THE NUCLEUS! \n" +
				"Drag the complementary nucleobase \n" +
				"to RNA Polymerase to transcribe a \n" +
				"new mRNA molecule!!!",
		}

		copy(transcriptionStruct.DNA[:], dna[:])
		transcriptionStruct.DNA[5] = dna[4]
		copy(transcriptionStruct.RNA[:], rna[:])
		transcriptionStruct.RNA[5] = newTranscript("RNA5.png", transcriptionStruct.RNA[4].rect, transcriptionStruct.RNA[4].codon, true)

		g.transcriptionSprites = []GUI{
			&transcriptionStruct.nucleusBg, &transcriptionStruct.temp_tfa,
			&transcriptionStruct.DNA[0],
			&transcriptionStruct.RNA[1], &transcriptionStruct.RNA[2],
			&transcriptionStruct.RNA[3], &transcriptionStruct.RNA[4],
			&transcriptionStruct.rnaPolymerase,
			&transcriptionStruct.A_Choice,
			&transcriptionStruct.T_Choice,
			&transcriptionStruct.G_Choice,
			&transcriptionStruct.C_Choice,
			&transcriptionStruct.U_Choice,
			&transcriptionStruct.otherToMenuButton, 
			&transcriptionStruct.infoButton,
		}
	}
	g.stateMachine.state = transcriptionStruct
}

func (t *TranscriptionLevel) Init(g *Game) {
	currentFrag = 0
	currBase = 0
	t.origRNAbases[0] = newNucleobase("N/A", newRect(0, 0, 65, 150), 0, false)
	for x := 0; x < len(t.origRNAbases)-1; x++ {
		base := string(t.RNA[x/3].codon[x%3])
		posX := 125 + t.RNA[currentFrag].rect.pos.x + (50 * x)
		posY := 220 + t.RNA[currentFrag].rect.pos.y + 220 - (15 * x)
		t.origRNAbases[x+1] = newNucleobase(base, newRect(posX, posY, 65, 150), x, false)
	}
	t.RNAbases = t.origRNAbases
	for x := 0; x < len(t.DNAbases); x++ {
		base := string(t.DNA[x/3].codon[x%3])
		posX := t.DNA[2].rect.pos.x + (50 * x) - 100
		posY := t.DNA[2].rect.pos.y
		if (x < 15) {
			t.DNAbases[x] = newNucleobase(base, newRect(posX, posY, 65, 150), x, true)
		} else {
			t.DNAbases[x] = newNucleobase("Term-\ninator", newRect(posX, posY, 65, 150), x, true)
		}
	}
	t.ResetChoices()
	g.state_array = g.transcriptionSprites
	t.temp_tfa.activate()
}

func (t *TranscriptionLevel) ResetChoices() {
	rand.Shuffle(len(spots), func(i, j int) { spots[i], spots[j] = spots[j], spots[i] })
	rand.Shuffle(len(baseSpots), func(i, j int) { baseSpots[i], baseSpots[j] = baseSpots[j], baseSpots[i] })

	t.A_Choice.reset(0, 550)
	t.T_Choice.reset(1, 550)
	t.G_Choice.reset(2, 550)
	t.C_Choice.reset(3, 550)
	t.U_Choice.reset(4, 550)

	for x := 0; x < (currBase); x++ {
		temp := currBase - x
		base := t.RNAbases[x]
		base.rect.pos.y = t.origRNAbases[temp].rect.pos.y
		t.RNAbases[x] = base
	}
	reset = false
}

func (t *TranscriptionLevel) Update(g *Game) {
	t.otherToMenuButton.update(g)
	t.RNA[currentFrag].update()
	if currentFrag < 5 {
		t.RNA[currentFrag+1].update()
	}
	t.infoButton.update()
	t.temp_tfa.update()
	t.rnaPolymerase.update(g)

	t.A_Choice.update(&t.DNAbases, currBase)
	t.T_Choice.update(&t.DNAbases, currBase)
	t.G_Choice.update(&t.DNAbases, currBase)
	t.C_Choice.update(&t.DNAbases, currBase)
	t.U_Choice.update(&t.DNAbases, currBase)

	for i, base := range t.RNAbases {
		base.update()
		t.RNAbases[i] = base
	}

	if t.DNAbases[currBase].isComplete {
		if (currBase+1)%3 == 0 {
			t.DNA[currentFrag].is_complete = true
		}
		currBase++
		t.ResetChoices()
	}

	// Checks if current DNA codon is complete
	if t.DNA[currentFrag].is_complete && !t.rnaPolymerase.next {
		nextDNACodon()
	}
	if t.RNA[5].rect.pos.y <= -600 {
		ToCyto2(g)
		reset = false
	}
}

func (t *TranscriptionLevel) Draw(g *Game, screen *ebiten.Image) {
	t.nucleusBg.draw(screen)

	t.RNA[currentFrag].draw(screen)
	t.rnaPolymerase.draw(screen)
	t.DNA[0].draw(screen)
	t.temp_tfa.draw(screen)

	t.A_Choice.draw(screen)
	t.T_Choice.draw(screen)
	t.G_Choice.draw(screen)
	t.C_Choice.draw(screen)
	t.U_Choice.draw(screen)

	for y := 0; y <= currBase; y++ {
		t.RNAbases[y].draw(screen)
	}

	if currentFrag != -1 {
		for _, base := range t.DNAbases {
			if (base.index != currBase) {
				base.opColor.SetA(0.25)
			} else {base.opColor.SetA(1)}
			base.draw(screen)
		}
	}

	defaultFont.drawFont(screen, t.message, 75, 50, color.Black)

	t.otherToMenuButton.draw(screen)

	t.infoButton.draw(screen)
}
