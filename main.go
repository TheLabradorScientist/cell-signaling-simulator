package main

import (
	"image/color"
	_ "image/png"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 960
	screenHeight = 960
)

var (
	err           error
	scene         string = "Main Menu"
	startBg       *ebiten.Image
	plasmaBg      *ebiten.Image
	cytoBg_1      *ebiten.Image
	nucleusBg     *ebiten.Image
	cytoBg_2      *ebiten.Image
	playbutton    Button
	seedSignal    = rand.Intn(4) + 1
	signal        Signal
	receptorA     Receptor
	receptorB     Receptor
	receptorC     Receptor
	receptorD     Receptor
	template      = [5]string{}
	tk1           Kinase
	tk2           Kinase
	tfa           TFA
	rna           [5]RNA
	dna           [5]DNA
	rnaPolymerase *ebiten.Image
	ribosome      Ribosome
	rightChoice   CodonChoice
	wrongChoice1  CodonChoice
	wrongChoice2  CodonChoice
	trna          [5]RNA
	trna_ptr      int
)

type Game struct {
	defaultFont            Font
	switchedMenuToPlasma   bool
	switchedPlasmaToMenu   bool
	switchedPlasmaToCyto1  bool
	switchedCyto1ToNucleus bool
	switchedNucleusToCyto2 bool
}

func MenuToPlasma(g *Game) {
	g.switchedMenuToPlasma = true
	g.switchedPlasmaToMenu = false
	g.switchedPlasmaToCyto1 = false
	g.switchedCyto1ToNucleus = false
	g.switchedNucleusToCyto2 = false
}

func PlasmaToMenu(g *Game) {
	g.switchedMenuToPlasma = false
	g.switchedPlasmaToMenu = true
	g.switchedPlasmaToCyto1 = false
	g.switchedCyto1ToNucleus = false
	g.switchedNucleusToCyto2 = false
}

func PlasmaToCyto1(g *Game) {
	g.switchedMenuToPlasma = false
	g.switchedPlasmaToMenu = false
	g.switchedPlasmaToCyto1 = true
	g.switchedCyto1ToNucleus = false
	g.switchedNucleusToCyto2 = false
}

func Cyto1ToNucleus(g *Game) {
	g.switchedMenuToPlasma = false
	g.switchedPlasmaToMenu = false
	g.switchedPlasmaToCyto1 = false
	g.switchedCyto1ToNucleus = true
	g.switchedNucleusToCyto2 = false
}

func NucleusToCyto2(g *Game) {
	g.switchedMenuToPlasma = false
	g.switchedPlasmaToMenu = false
	g.switchedPlasmaToCyto1 = false
	g.switchedCyto1ToNucleus = false
	g.switchedNucleusToCyto2 = true
}

func init() {
	startBg, _, err = ebitenutil.NewImageFromFile("MenuBg.png")
	if err != nil {
		log.Fatal(err)
	}
	plasmaBg, _, err = ebitenutil.NewImageFromFile("PlasmaBg.png")
	if err != nil {
		log.Fatal(err)
	}
	cytoBg_1, _, err = ebitenutil.NewImageFromFile("CytoBg1.png")
	if err != nil {
		log.Fatal(err)
	}
	nucleusBg, _, err = ebitenutil.NewImageFromFile("NucleusBg.png")
	if err != nil {
		log.Fatal(err)
	}
	cytoBg_2, _, err = ebitenutil.NewImageFromFile("CytoBg2.png")
	if err != nil {
		log.Fatal(err)
	}
	playbutton = newButton("PlayButton.png", newRect(400, 300, 232, 129), MenuToPlasma)

	switch seedSignal {
	case 1:
		signal = newSignal("signalA.png", newRect(100, 100, 75, 75))
		signal.signalType = "signalA"
		// PLACEHOLDER IN CASE WE DO NOT GET TIME TO CODE RANDOM CODONS
		template = [5]string{"TAC", "GTC", "CGG", "ACA", "UGA"}

	case 2:
		signal = newSignal("signalB.png", newRect(100, 100, 75, 75))

		signal.signalType = "signalB"
	case 3:
		signal = newSignal("signalC.png", newRect(100, 100, 75, 75))
		signal.signalType = "signalC"
	case 4:
		signal = newSignal("signalD.png", newRect(100, 100, 75, 75))
		signal.signalType = "signalD"
	}

	receptorA = newReceptor("receptorA.png", newRect(0, 500, 100, 150), "receptorA")
	receptorB = newReceptor("receptorB.png", newRect(250, 500, 100, 150), "receptorB")
	receptorC = newReceptor("receptorC.png", newRect(500, 500, 100, 150), "receptorC")
	receptorD = newReceptor("receptorD.png", newRect(750, 500, 100, 150), "receptorD")

	tk1 = newKinase("TK1.png", newRect(400, 100, 150, 150), "tk1")
	tk2 = newKinase("TK2.png", newRect(100, 175, 150, 150), "tk2")
	tfa = newTFA("TFA.png", newRect(300, 500, 150, 150))

	for x := 0; x < 5; x++ {
		dna[x] = newDNA("DNA.png", newRect(-100, 500, 150, 150), template[x], x)
	}
	for x := 0; x < 5; x++ {
		rna[x] = newRNA("RNA.png", newRect(-100, 200, 150, 150), transcribe(template[x]))
	}

	rnaPolymerase, _, err = ebitenutil.NewImageFromFile("rnaPolym.png")
	if err != nil {
		log.Fatal(err)
	}

	rightChoice = newCodonChoice("codonButton", newRect(0, 0, 258, 144), transcribe(dna[0].codon))
	wrongChoice1 = newCodonChoice("codonButton", newRect(100, 0, 258, 144), randomize())
	wrongChoice2 = newCodonChoice("codonButton", newRect(200, 0, 258, 144), randomize())
	ribosome = newRibosome("PlayButton.png", newRect(0, 0, 160, 160))
	trna_ptr = 0
}

func (g *Game) Update() error {
	switch scene {
	case "Main Menu":
		ebiten.SetWindowTitle("Cell Signaling Synthesis - Main Menu")
		ebiten.SetWindowSize(screenWidth, screenHeight)
		playbutton.on_click(g)
	case "Signal Reception":
		ebiten.SetWindowTitle("Cell Signaling Synthesis - Signal Reception")
		ebiten.SetWindowSize(screenWidth, screenHeight)
		signal.on_click(g)
		receptorA.update()
		receptorB.update()
		receptorC.update()
		receptorD.update()
		if receptorA.is_touching_signal {
			if matchSR(signal.signalType, receptorA.receptorType) {
				PlasmaToCyto1(g)
			}
		}
		if receptorB.is_touching_signal {
			if matchSR(signal.signalType, receptorB.receptorType) {
				PlasmaToCyto1(g)
			}
		}
		if receptorC.is_touching_signal {
			if matchSR(signal.signalType, receptorC.receptorType) {
				PlasmaToCyto1(g)
			}
		}
		if receptorD.is_touching_signal {
			if matchSR(signal.signalType, receptorD.receptorType) {
				PlasmaToCyto1(g)
			}
		}
	case "Signal Transduction":
		ebiten.SetWindowTitle("Cell Signaling Synthesis - Signal Transduction")
		ebiten.SetWindowSize(screenWidth, screenHeight)
		tk1.activate()
		tk1.update(tk2.rect)
		tk2.update(tfa.rect)
		tfa.update()
		if tk1.is_clicked_on {
			tk2.activate()
			tk1.is_clicked_on = false
		}
		if tk2.is_clicked_on {
			tfa.activate()
			tk2.is_clicked_on = false
		}
		if tfa.rect.pos.y > 750 {
			Cyto1ToNucleus(g)
		}
	case "Transcription":
		ebiten.SetWindowTitle("Cell Signaling Synthesis - Transcription")
	case "Translation":
		ebiten.SetWindowTitle("Cell Signaling Synthesis - Translation")
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch scene {
	case "Main Menu":
		screen.DrawImage(startBg, nil)
		playbutton.draw(screen)
		//ebitenutil.DebugPrint(screen, transcribe("CAT"))
		//ebitenutil.DebugPrint(screen, translate("UAG"))
		if g.switchedMenuToPlasma {
			scene = "Signal Reception"
		}
	case "Signal Reception":
		screen.DrawImage(plasmaBg, nil)
		receptorA.draw(screen)
		receptorB.draw(screen)
		receptorC.draw(screen)
		receptorD.draw(screen)
		signal.draw(screen)
		g.defaultFont.drawFont(screen, "WELCOME TO THE PLASMA MEMBRANE! \n Drag the signal to the matching receptor to enter the cell!", 100, 100, color.White)
		if signal.is_dragged {
			signal.draw(screen)
		}
		if g.switchedPlasmaToMenu {
			scene = "Main Menu"
		}
		if g.switchedPlasmaToCyto1 {
			scene = "Signal Transduction"
		}
	case "Signal Transduction":
		screen.DrawImage(cytoBg_1, nil)
		tk1.draw(screen)
		tk2.draw(screen)
		tfa.draw(screen)
		g.defaultFont.drawFont(screen, "WELCOME TO THE CYTOPLASM! \n Click when each kinase overlaps to follow the phosphorylation cascade!!", 100, 100, color.Black)
		if g.switchedCyto1ToNucleus {
			scene = "Transcription"
		}
	case "Transcription":
		screen.DrawImage(nucleusBg, nil)
		for x := 0; x < 5; x++ {
			rna[x].draw(screen)
			dna[x].draw(screen)
		}
		g.defaultFont.drawFont(screen, "WELCOME TO THE NUCLEUS! \n Match each codon on the DNA template to the corresponding\n RNA codon to transcribe a new mRNA molecule!!!", 100, 100, color.White)
		if g.switchedNucleusToCyto2 {
			scene = "Transcription"
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, 300)
		screen.DrawImage(rnaPolymerase, op)

		rightChoice.draw(screen)
		wrongChoice1.draw(screen)
		wrongChoice2.draw(screen)

	case "Translation":
		screen.DrawImage(cytoBg_2, nil)

		ribosome.draw(screen)
		trna[trna_ptr].draw(screen)
		g.defaultFont.drawFont(screen, "FINALLY, BACK TO THE CYTOPLASM! \n Match each codon from your mRNA template to its corresponding amino acid to synthesize your protein!!!!", 100, 100, color.Black)

	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &Game{}
	game.defaultFont = newFont("ThaleahFat.ttf", 25)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
