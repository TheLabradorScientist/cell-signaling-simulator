package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	mrna_ptr = 0
	mrna     [5]Template
	protein  [5]Transcript
)

type TranslationLevel struct {
	// CYTO 2 SPRITES
	cytoBg_2          StillImage
	ribosome          Ribosome
	rightTrna         tRNA
	wrongTrna1        tRNA
	wrongTrna2        tRNA
	infoButton        InfoPage
	otherToMenuButton Button
	message           string
}

var translationStruct *TranslationLevel

func newTranslationLevel(g *Game) {
	if len(g.translationSprites) == 0 {
		translationStruct = &TranslationLevel{
			cytoBg_2: newStillImage("CytoBg2.png", newRect(0, 0, 1250, 750)),

			ribosome: newRibosome("ribosome.png", newRect(40, 300, 404, 367)),

			message: 
				"FINALLY, BACK TO THE CYTOPLASM! \n" +
				"Match each codon from your mRNA template \n" +
				"to its corresponding amino acid to \n" +
				"synthesize your protein!!!!",
		}

		translationStruct.rightTrna = newTRNA("codonButton.png", newRect(50, 150, 192, 111), translate(mrna[0].codon))
		translationStruct.wrongTrna1 = newTRNA("codonButton.png", newRect(350, 150, 192, 111), translate(randomRNACodon(translationStruct.rightTrna.bases)))
		translationStruct.wrongTrna2 = newTRNA("codonButton.png", newRect(650, 150, 192, 111), translate(randomRNACodon(translationStruct.rightTrna.bases)))
		translationStruct.infoButton = infoButton
		translationStruct.otherToMenuButton = otherToMenuButton

		g.translationSprites = []GUI{
			&translationStruct.cytoBg_2, &translationStruct.ribosome, 
			&translationStruct.rightTrna, &translationStruct.wrongTrna1,
			&translationStruct.wrongTrna2, &translationStruct.otherToMenuButton,
			&translationStruct.infoButton, 
		}
	}
	g.stateMachine.state = translationStruct
}

func (t *TranslationLevel) Init(g *Game) {
	mrna_ptr = 0
	reset = false
	t.ResetChoices()
	g.state_array = g.translationSprites
}

func (t *TranslationLevel) ResetChoices() {
	curr := &mrna[mrna_ptr]
	rand.Shuffle(len(spots), func(i, j int) {spots[i], spots[j] = spots[j], spots[i]})
	t.rightTrna.reset(0, translate(curr.codon))
	t.wrongTrna1.reset(1, translate(randomRNACodon(t.rightTrna.bases)))
	t.wrongTrna2.reset(2, translate(randomRNACodon(t.rightTrna.bases)))
	reset = false
}

func (t *TranslationLevel) Update(g *Game) {
	t.otherToMenuButton.update(g)
	t.infoButton.update()

	curr := &mrna[mrna_ptr]

	if reset {t.ResetChoices()}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		t.rightTrna.update(curr)
	}

	t.ribosome.update(g)
}

func (t *TranslationLevel) Draw(g *Game, screen *ebiten.Image) {
	t.cytoBg_2.draw(screen)

	mrna[0].draw(screen)

	t.rightTrna.draw(screen)
	t.wrongTrna1.draw(screen)
	t.wrongTrna2.draw(screen)

	t.infoButton.draw(screen)
	t.otherToMenuButton.draw(screen)

	// Draw amino acids before ribosome moves without drawing amino acid for STOP.
	for x := 0; x <= mrna_ptr; x++ {
		if x < 4 {
			protein[x].draw(screen)
			codonFont.drawFont(screen, protein[x].codon, protein[x].rect.pos.x, protein[x].rect.pos.y, color.Black)
		}
	}

	if mrna_ptr != -1 {
		codonFont.drawFont(screen, mrna[mrna_ptr].codon, mrna[0].rect.pos.x+500, mrna[0].rect.pos.y+200, color.Black)
	}

	t.ribosome.draw(screen)

	defaultFont.drawFont(screen, t.message, 100, 50, color.Black)
}
