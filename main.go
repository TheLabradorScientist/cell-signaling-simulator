package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 1250
	screenHeight = 750
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
	rna           [5]Transcript
	dna           [5]Template
	currentFrag   = 0
	temp_tk1A      Kinase
	temp_tk1B      Kinase
	temp_tk1C      Kinase
	temp_tk1D      Kinase
	rnaPolymerase *ebiten.Image
	ribosome      Ribosome
	rightChoice   CodonChoice
	wrongChoice1  CodonChoice
	wrongChoice2  CodonChoice
	mrna          [5]Template
	protein       [5]Transcript
	mrna_ptr      int
	rightTrna     CodonChoice
	wrongTrna1    CodonChoice
	wrongTrna2    CodonChoice
	reset         bool
)

type Game struct {
	defaultFont            Font
	codonFont              Font
	switchedMenuToPlasma   bool
	switchedPlasmaToMenu   bool
	switchedPlasmaToCyto1  bool
	switchedCyto1ToNucleus bool
	switchedNucleusToCyto2 bool
}


func loadFile(image string) string {
	// Construct path to file
	imagepath := filepath.Join("Assets", "Images", image)
	// Open file
	file, err := os.Open(imagepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "Error"
	}

	defer file.Close()

	return file.Name()
}

func init() {
	startBg, _, err = ebitenutil.NewImageFromFile(loadFile("MenuBg.png"))
	if err != nil {
		log.Fatal(err)
	}
	plasmaBg, _, err = ebitenutil.NewImageFromFile(loadFile("PlasmaBg.png"))
	if err != nil {
		log.Fatal(err)
	}
	cytoBg_1, _, err = ebitenutil.NewImageFromFile(loadFile("CytoBg1.png"))
	if err != nil {
		log.Fatal(err)
	}
	nucleusBg, _, err = ebitenutil.NewImageFromFile(loadFile("NucleusBg.png"))
	if err != nil {
		log.Fatal(err)
	}
	cytoBg_2, _, err = ebitenutil.NewImageFromFile(loadFile("CytoBg2.png"))
	if err != nil {
		log.Fatal(err)
	}
	playbutton = newButton("PlayButton.png", newRect(500, 500, 232, 129), MenuToPlasma)

	switch seedSignal {
	case 1:
		signal = newSignal("signalA.png", newRect(500, 100, 75, 75))
		signal.signalType = "signalA"
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ACT"}
	case 2:
		signal = newSignal("signalB.png", newRect(500, 100, 75, 75))
		signal.signalType = "signalB"
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATT"}
	case 3:
		signal = newSignal("signalC.png", newRect(500, 100, 75, 75))
		signal.signalType = "signalC"
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATC"}
	case 4:
		signal = newSignal("signalD.png", newRect(500, 100, 75, 75))
		signal.signalType = "signalD"
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATT"}
		// PLACEHOLDER IN CASE WE DO NOT GET TIME TO CODE RANDOM CODONS
		//template = [5]string{"TAC", "GTC", "CGG", "ACA", "ACT"}
	}

	receptorA = newReceptor("inact_receptorA.png", newRect(50, 400, 100, 150), "receptorA")
	receptorB = newReceptor("inact_receptorB.png", newRect(350, 400, 100, 150), "receptorB")
	receptorC = newReceptor("inact_receptorC.png", newRect(650, 400, 100, 150), "receptorC")
	receptorD = newReceptor("inact_receptorD.png", newRect(950, 400, 100, 150), "receptorD")

	temp_tk1A = newKinase("inact_TK1.png", newRect(50, 600, 150, 150), "temp_tk1")
	temp_tk1B = newKinase("inact_TK1.png", newRect(350, 600, 150, 150), "temp_tk1")
	temp_tk1C = newKinase("inact_TK1.png", newRect(650, 600, 150, 150), "temp_tk1")
	temp_tk1D = newKinase("inact_TK1.png", newRect(950, 600, 150, 150), "temp_tk1")

	tk1 = newKinase("inact_TK1.png", newRect(500, -100, 150, 150), "tk1")
	tk2 = newKinase("inact_TK2.png", newRect(250, 175, 150, 150), "tk2")
	tfa = newTFA("inact_TFA.png", newRect(700, 500, 150, 150))

	for x := 0; x < 5; x++ {
		dna[x] = newTemplate("DNA.png", newRect(-100+200*x, 500, 150, 150), template[x], x)
	}
	for x := 0; x < 5; x++ {
		rna[x] = newTranscript("RNA.png", newRect(0, 200, 150, 150), transcribe(template[x]))
	}

	if err != nil {
		log.Fatal(err)
	}

	rnaPolymerase, _, err = ebitenutil.NewImageFromFile(loadFile("rnaPolym.png"))
	if err != nil {
		log.Fatal(err)
	}

	for x := 0; x < 5; x++ {
		mrna[x] = newTemplate("RNA.png", newRect(0, 400, 150, 150), transcribe(dna[x].codon), x)
	}
	for x := 0; x < 5; x++ {
		protein[x] = newTranscript("RNA.png", newRect(-100, 400, 150, 150), translate(mrna[x].codon))
	}

	reset = false

	rightChoice = newCodonChoice("codonButton.png", newRect(50, 150, 192, 111), transcribe(dna[0].codon))
	wrongChoice1 = newCodonChoice("codonButton.png", newRect(350, 150, 192, 111), randomRNACodon(rightChoice.bases))
	wrongChoice2 = newCodonChoice("codonButton.png", newRect(650, 150, 192, 111), randomRNACodon(rightChoice.bases))

	ribosome = newRibosome("ribosome.png", newRect(0, 300, 404, 367))

	mrna_ptr = 0

	rightTrna = newCodonChoice("codonButton.png", newRect(50, 150, 192, 111), translate(mrna[0].codon))
	wrongTrna1 = newCodonChoice("codonButton.png", newRect(350, 150, 192, 111), translate(randomRNACodon(rightTrna.bases)))
	wrongTrna2 = newCodonChoice("codonButton.png", newRect(650, 150, 192, 111), translate(randomRNACodon(rightTrna.bases)))

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
		temp_tk1A.update(temp_tk1B.rect)
		temp_tk1B.update(temp_tk1C.rect)
		temp_tk1C.update(temp_tk1D.rect)
		temp_tk1D.update(temp_tk1A.rect)
		if receptorA.is_touching_signal {
			if matchSR(signal.signalType, receptorA.receptorType) {
				receptorA.animate("act_receptorA.png")
				temp_tk1A.activate()
			}
		}
		if receptorB.is_touching_signal {
			if matchSR(signal.signalType, receptorB.receptorType) {
				receptorB.animate("act_receptorB.png")
				temp_tk1B.activate()
			}
		}
		if receptorC.is_touching_signal {
			if matchSR(signal.signalType, receptorC.receptorType) {
				receptorC.animate("act_receptorC.png")
				temp_tk1C.activate()
			}
		}
		if receptorD.is_touching_signal {
			if matchSR(signal.signalType, receptorD.receptorType) {
				receptorD.animate("act_receptorD.png")
				temp_tk1D.activate()
			}
		}
		if (temp_tk1A.rect.pos.y >= 750 || temp_tk1B.rect.pos.y >= 750 || temp_tk1C.rect.pos.y >= 750 || temp_tk1D.rect.pos.y >= 750) {
			PlasmaToCyto1(g)
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
		if reset {
			rightChoice.bases = transcribe(dna[currentFrag].codon)
			wrongChoice1.bases = randomRNACodon(rightChoice.bases)
			wrongChoice2.bases = randomRNACodon(rightChoice.bases)
			reset = false
		}
		//fmt.Printf("%t\n", dna[currentFrag].is_complete)
		dna[currentFrag].is_complete = rightChoice.update1(g, dna[currentFrag].codon)
		//fmt.Printf("%t\n", dna[currentFrag].is_complete)
		if dna[currentFrag].is_complete {
			nextDNACodon(g)
		}

	case "Translation":
		ebiten.SetWindowTitle("Cell Signaling Synthesis - Translation")
		if reset {
			rightTrna.bases = translate(mrna[mrna_ptr].codon)
			wrongTrna1.bases = translate(randomRNACodon(rightTrna.bases))
			wrongTrna2.bases = translate(randomRNACodon(rightTrna.bases))
			reset = false
		}
		mrna[mrna_ptr].is_complete = rightTrna.update2(g, mrna[mrna_ptr].codon)

		if mrna[mrna_ptr].is_complete {
			nextMRNACodon(g)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch scene {
	case "Main Menu":
		screen.DrawImage(startBg, nil)
		playbutton.draw(screen)
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
		temp_tk1A.draw(screen)
		temp_tk1B.draw(screen)
		temp_tk1C.draw(screen)
		temp_tk1D.draw(screen)
		g.defaultFont.drawFont(screen, "WELCOME TO THE PLASMA MEMBRANE! \n Drag the signal to the matching receptor \n to enter the cell!", 100, 50, color.White)
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
		g.defaultFont.drawFont(screen, "WELCOME TO THE CYTOPLASM! \n Click when each kinase overlaps to follow \n the phosphorylation cascade!!", 100, 50, color.Black)
		if g.switchedCyto1ToNucleus {
			scene = "Transcription"
		}
	case "Transcription":
		screen.DrawImage(nucleusBg, nil)
		for x := 0; x < 5; x++ {
			rna[x].draw(screen)
		}
		dna[0].draw(screen)

		g.defaultFont.drawFont(screen, "WELCOME TO THE NUCLEUS! \n Match each codon on the DNA template to the corresponding RNA \n codon to transcribe a new mRNA molecule!!!", 100, 50, color.White)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, 300)
		screen.DrawImage(rnaPolymerase, op)
		//g.codonFont.drawFont(screen, strings.Join(template[0:5], ""), dna[currentFrag].rect.pos.x+300, dna[currentFrag].rect.pos.y, color.Black)
		if currentFrag != -1 {
			g.codonFont.drawFont(screen, dna[currentFrag].codon, dna[0].rect.pos.x+500, dna[0].rect.pos.y, color.Black)
		}
		rightChoice.draw(screen)
		wrongChoice1.draw(screen)
		wrongChoice2.draw(screen)
		g.codonFont.drawFont(screen, rightChoice.bases, rightChoice.rect.pos.x+25, rightChoice.rect.pos.y+90, color.Black)
		g.codonFont.drawFont(screen, wrongChoice1.bases, wrongChoice1.rect.pos.x+25, wrongChoice1.rect.pos.y+90, color.Black)
		g.codonFont.drawFont(screen, wrongChoice2.bases, wrongChoice2.rect.pos.x+25, wrongChoice2.rect.pos.y+90, color.Black)
		switch currentFrag {
		case 1:
			g.codonFont.drawFont(screen, rna[0].codon, rna[0].rect.pos.x+500, rna[0].rect.pos.y+140, color.Black)
		case 2:
			g.codonFont.drawFont(screen, rna[0].codon, rna[0].rect.pos.x+500, rna[0].rect.pos.y+140, color.Black)
			g.codonFont.drawFont(screen, rna[1].codon, rna[0].rect.pos.x+650, rna[0].rect.pos.y+140, color.Black)
		case 3:
			g.codonFont.drawFont(screen, rna[0].codon, rna[0].rect.pos.x+500, rna[0].rect.pos.y+140, color.Black)
			g.codonFont.drawFont(screen, rna[1].codon, rna[0].rect.pos.x+650, rna[0].rect.pos.y+140, color.Black)
			g.codonFont.drawFont(screen, rna[2].codon, rna[0].rect.pos.x+800, rna[0].rect.pos.y+140, color.Black)
		case 4:
			g.codonFont.drawFont(screen, rna[0].codon, rna[0].rect.pos.x+500, rna[0].rect.pos.y+140, color.Black)
			g.codonFont.drawFont(screen, rna[1].codon, rna[0].rect.pos.x+650, rna[0].rect.pos.y+140, color.Black)
			g.codonFont.drawFont(screen, rna[2].codon, rna[0].rect.pos.x+800, rna[0].rect.pos.y+140, color.Black)
			g.codonFont.drawFont(screen, rna[3].codon, rna[0].rect.pos.x+950, rna[0].rect.pos.y+140, color.Black)
		default:
			break
		}

		if g.switchedNucleusToCyto2 {
			scene = "Translation"
		}
	case "Translation":
		screen.DrawImage(cytoBg_2, nil)
		for x := 0; x < 5; x++ {
			protein[x].draw(screen)
		}

		mrna[0].draw(screen)
		ribosome.draw(screen)

		if mrna_ptr != -1 {
			g.codonFont.drawFont(screen, mrna[mrna_ptr].codon, mrna[0].rect.pos.x+500, mrna[0].rect.pos.y+200, color.Black)
		}

		rightTrna.draw(screen)
		wrongTrna1.draw(screen)
		wrongTrna2.draw(screen)

		g.codonFont.drawFont(screen, rightTrna.bases, rightTrna.rect.pos.x+25, rightTrna.rect.pos.y+90, color.Black)
		g.codonFont.drawFont(screen, wrongTrna1.bases, wrongTrna1.rect.pos.x+25, wrongTrna1.rect.pos.y+90, color.Black)
		g.codonFont.drawFont(screen, wrongTrna2.bases, wrongTrna2.rect.pos.x+25, wrongTrna2.rect.pos.y+90, color.Black)

		g.defaultFont.drawFont(screen, "FINALLY, BACK TO THE CYTOPLASM! \n Match each codon from your mRNA template \n to its corresponding amino acid to synthesize your protein!!!!", 100, 50, color.Black)
		switch mrna_ptr {
		case 1:
			g.codonFont.drawFont(screen, protein[0].codon, protein[0].rect.pos.x+150, protein[0].rect.pos.y, color.Black)
		case 2:
			g.codonFont.drawFont(screen, protein[0].codon, protein[0].rect.pos.x+150, protein[0].rect.pos.y, color.Black)
			g.codonFont.drawFont(screen, protein[1].codon, protein[0].rect.pos.x+300, protein[0].rect.pos.y, color.Black)
		case 3:
			g.codonFont.drawFont(screen, protein[0].codon, protein[0].rect.pos.x+150, protein[0].rect.pos.y, color.Black)
			g.codonFont.drawFont(screen, protein[1].codon, protein[0].rect.pos.x+300, protein[0].rect.pos.y, color.Black)
			g.codonFont.drawFont(screen, protein[2].codon, protein[0].rect.pos.x+450, protein[0].rect.pos.y, color.Black)
		case 4:
			g.codonFont.drawFont(screen, protein[0].codon, protein[0].rect.pos.x+150, protein[0].rect.pos.y, color.Black)
			g.codonFont.drawFont(screen, protein[1].codon, protein[0].rect.pos.x+300, protein[0].rect.pos.y, color.Black)
			g.codonFont.drawFont(screen, protein[2].codon, protein[0].rect.pos.x+450, protein[0].rect.pos.y, color.Black)
			g.codonFont.drawFont(screen, protein[3].codon, protein[0].rect.pos.x+600, protein[0].rect.pos.y, color.Black)
		default:
			break
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &Game{}
	game.defaultFont = newFont("CourierPrime-Regular.ttf", 32)
	game.codonFont = newFont("Honk-Regular.ttf", 90)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
