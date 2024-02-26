package main

import (
	"C"
	"fmt"
	"image/color"
	_ "image/png"
	
	"log"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	//"syscall/js"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 1250
	screenHeight = 750
)

var (
	err                error
	scene              string = "Main Menu"
	protoStartBg       *ebiten.Image
	startBg			   Parallax
	startP1			   Parallax
	startP2			   Parallax
	startP3			   Parallax
	startP4			   Parallax
	fixedStart         *ebiten.Image
	infoBg             *ebiten.Image
	plasmaBg           *ebiten.Image
	cytoBg_1           *ebiten.Image
	nucleusBg          *ebiten.Image
	cytoBg_2           *ebiten.Image
	levSelBg           *ebiten.Image
	playbutton         Button
	infoButton         Button
	infoToMenuButton   Button
	levSelButton       Button
	levToPlasmaButton  Button
	levToCyto1Button   Button
	levToNucleusButton Button
	levToCyto2Button   Button
	levToMenuButton    Button
	seedSignal         = rand.Intn(4) + 1
	signal             Signal
	receptorA          Receptor
	receptorB          Receptor
	receptorC          Receptor
	receptorD          Receptor
	template           = [5]string{}
	tk1                Kinase
	tk2                Kinase
	tfa                TFA
	rna                [5]Transcript
	dna                [5]Template
	currentFrag        = 0
	temp_tk1A          Kinase
	temp_tk1B          Kinase
	temp_tk1C          Kinase
	temp_tk1D          Kinase
	temp_tfa           TFA
	rnaPolymerase      RNAPolymerase
	ribosome           Ribosome
	rightChoice        CodonChoice
	wrongChoice1       CodonChoice
	wrongChoice2       CodonChoice
	mrna               [5]Template
	protein            [5]Transcript
	mrna_ptr           int
	rightTrna          CodonChoice
	wrongTrna1         CodonChoice
	wrongTrna2         CodonChoice
	reset              bool
)

type Game struct {
	defaultFont           Font
	codonFont             Font
	switchedToPlasma      bool
	switchedToMenu        bool
	switchedToCyto1       bool
	switchedToNucleus     bool
	switchedToCyto2       bool
	switchedToInfo        bool
	switchedToLevelSelect bool
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
	protoStartBg, _, err = ebitenutil.NewImageFromFile(loadFile("MenuBg.png"))
	startBg = newParallax("StartBg.png", newRect(0, 0, 1250, 750), 5)
	startP1 = newParallax("parallax-Start2.png", newRect(0, 0, 1250, 750), 4)
	startP2 = newParallax("parallax-Start3.png", newRect(0, 0, 1250, 750), 3)
	startP3 = newParallax("parallax-Start4.png", newRect(0, 0, 1250, 750), 2)
	startP4 = newParallax("parallax-Start5.png", newRect(0, 0, 1250, 750), 1)
	fixedStart, _, err = ebitenutil.NewImageFromFile(loadFile("fixed-Start.png"))
	if err != nil {
		log.Fatal(err)
	}
	infoBg, _, err = ebitenutil.NewImageFromFile(loadFile("InfoBg.png"))
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
	levSelBg, _, err = ebitenutil.NewImageFromFile(loadFile("levSelBg.png"))
	if err != nil {
		log.Fatal(err)
	}

	playbutton = newButton("PlayButton.png", newRect(750, 100, 242, 138), ToPlasma)
	infoButton = newButton("infoButton.png", newRect(770, 260, 242, 138), ToInfo)
	levSelButton = newButton("levSelButton.png", newRect(700, 450, 232, 140), ToLevelSelect)
	infoToMenuButton = newButton("menuButton.png", newRect(350, 450, 242, 138), ToMenu)
	levToPlasmaButton = newButton("levToPlasmaBtn.png", newRect(520, 110, 232, 129), ToPlasma)
	levToCyto1Button = newButton("levToCyto1Btn.png", newRect(820, 110, 232, 129), ToCyto1)
	levToNucleusButton = newButton("levToNucleusBtn.png", newRect(520, 285, 232, 129), ToNucleus)
	levToCyto2Button = newButton("levToCyto2Btn.png", newRect(820, 285, 232, 129), ToCyto2)
	levToMenuButton = newButton("menuButton.png", newRect(250, 190, 232, 129), ToMenu)

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

	tk1 = newKinase("act_TK1.png", newRect(500, -100, 150, 150), "tk1")
	tk2 = newKinase("inact_TK2.png", newRect(250, 175, 150, 150), "tk2")
	tfa = newTFA("inact_TFA.png", newRect(700, 500, 150, 150), "tfa1")

	for x := 0; x < 5; x++ {
		dna[x] = newTemplate("DNA.png", newRect(-50+200*x, 500, 150, 150), template[x], x)
	}
	for x := 0; x < 5; x++ {
		rna[x] = newTranscript("RNA.png", newRect(0, 200, 150, 150), transcribe(template[x]))
	}

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	for x := 0; x < 5; x++ {
		mrna[x] = newTemplate("RNA.png", newRect(0, 400, 150, 150), transcribe(dna[x].codon), x)
	}
	for x := 0; x < 5; x++ {
		protein[x] = newTranscript("aminoAcid.png", newRect(50+(150*x), 400, 150, 150), translate(mrna[x].codon))
	}

	rnaPolymerase = newRNAPolymerase("rnaPolym.png", newRect(-350, 100, 340, 265))

	temp_tfa = newTFA("act_TFA.png", newRect(400, -100, 150, 150), "tfa2")

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
		ebiten.SetWindowTitle("Cell Signaling Pathway - Main Menu")
		ebiten.SetWindowSize(screenWidth, screenHeight)
		startBg.update(g)
		startP1.update(g)
		startP2.update(g)
		startP3.update(g)
		startP4.update(g)
		infoButton.on_click(g)
		playbutton.on_click(g)
		levSelButton.on_click(g)
	case "Information":
		ebiten.SetWindowTitle("Cell Signaling Pathway - About")
		ebiten.SetWindowSize(screenWidth, screenHeight)
		infoToMenuButton.on_click(g)
	case "Level Selection":
		ebiten.SetWindowTitle("Cell Signaling Pathway - Level Selection")
		ebiten.SetWindowSize(screenWidth, screenHeight)
		levToPlasmaButton.on_click(g)
		levToCyto1Button.on_click(g)
		levToNucleusButton.on_click(g)
		levToCyto2Button.on_click(g)
		levToMenuButton.on_click(g)
	case "Signal Reception":
		ebiten.SetWindowTitle("Cell Signaling Pathway - Signal Reception")
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
		if temp_tk1A.rect.pos.y >= 750 || temp_tk1B.rect.pos.y >= 750 || temp_tk1C.rect.pos.y >= 750 || temp_tk1D.rect.pos.y >= 750 {
			ToCyto1(g)
		}
	case "Signal Transduction":
		ebiten.SetWindowTitle("Cell Signaling Pathway - Signal Transduction")
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
			ToNucleus(g)
		}
	case "Transcription":
		ebiten.SetWindowTitle("Cell Signaling Pathway - Transcription")
		temp_tfa.activate()
		temp_tfa.update()
		rnaPolymerase.update(temp_tfa.rect.pos.y)
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
		ebiten.SetWindowTitle("Cell Signaling Pathway - Translation")
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
		screen.DrawImage(protoStartBg, nil)
		startBg.draw(screen)
		startP1.draw(screen)
		startP2.draw(screen)
		startP3.draw(screen)
		startP4.draw(screen)
		screen.DrawImage(fixedStart, nil)
		playbutton.draw(screen)
		infoButton.draw(screen)
		levSelButton.draw(screen)
		if g.switchedToPlasma {
			scene = "Signal Reception"
		}
		if g.switchedToLevelSelect {
			scene = "Level Selection"
		}
		if g.switchedToInfo {
			scene = "Information"
		}
	case "Information":
		screen.DrawImage(infoBg, nil)
		infoToMenuButton.draw(screen)
		m := "WELCOME TO THE CELL\nSIGNALING PATHWAY\nSIMULATOR!\n"
		m += "This simulator will\nguide you through the\ncomplete cell signaling\n"
		m += "pathway from reception\nthrough translation!\nClick the play "
		m += "button\nor select a level\nto begin."
		Red := color.RGBA{50, 5, 5, 250}
		g.defaultFont.drawFont(screen, m, 775, 275, color.RGBA(Red))

		if g.switchedToMenu {
			scene = "Main Menu"
		}
	case "Level Selection":
		screen.DrawImage(levSelBg, nil)
		levToPlasmaButton.draw(screen)
		levToCyto1Button.draw(screen)
		levToNucleusButton.draw(screen)
		levToCyto2Button.draw(screen)
		levToMenuButton.draw(screen)
		if g.switchedToMenu {
			scene = "Main Menu"
		}
		if g.switchedToPlasma {
			scene = "Signal Reception"
		}
		if g.switchedToCyto1 {
			scene = "Signal Transduction"
		}
		if g.switchedToNucleus {
			scene = "Transcription"
		}
		if g.switchedToCyto2 {
			scene = "Translation"
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
		if g.switchedToMenu {
			scene = "Main Menu"
		}
		if g.switchedToCyto1 {
			scene = "Signal Transduction"
		}
	case "Signal Transduction":
		screen.DrawImage(cytoBg_1, nil)
		tk1.draw(screen)
		tk2.draw(screen)
		tfa.draw(screen)
		g.defaultFont.drawFont(screen, "WELCOME TO THE CYTOPLASM! \n Click when each kinase overlaps to follow \n the phosphorylation cascade!!", 100, 50, color.Black)
		if g.switchedToNucleus {
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
		rnaPolymerase.draw(screen)
		temp_tfa.draw(screen)
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

		if g.switchedToCyto2 {
			scene = "Translation"
		}
	case "Translation":
		screen.DrawImage(cytoBg_2, nil)
		// for x := 0; x < 5; x++ {
		// 	protein[x].draw(screen)
		// }

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

		if g.switchedToMenu {
			scene = "Main Menu"
		}

		switch mrna_ptr {
		case 1:
			protein[0].draw(screen)
			g.codonFont.drawFont(screen, protein[0].codon, protein[0].rect.pos.x, protein[0].rect.pos.y, color.Black)
		case 2:
			protein[0].draw(screen)
			protein[1].draw(screen)
			g.codonFont.drawFont(screen, protein[0].codon, protein[0].rect.pos.x, protein[0].rect.pos.y, color.Black)
			g.codonFont.drawFont(screen, protein[1].codon, protein[1].rect.pos.x, protein[1].rect.pos.y, color.Black)
		case 3:
			protein[0].draw(screen)
			protein[1].draw(screen)
			protein[2].draw(screen)
			g.codonFont.drawFont(screen, protein[0].codon, protein[0].rect.pos.x, protein[0].rect.pos.y, color.Black)
			g.codonFont.drawFont(screen, protein[1].codon, protein[1].rect.pos.x, protein[1].rect.pos.y, color.Black)
			g.codonFont.drawFont(screen, protein[2].codon, protein[2].rect.pos.x, protein[2].rect.pos.y, color.Black)
		case 4:
			protein[0].draw(screen)
			protein[1].draw(screen)
			protein[2].draw(screen)
			protein[3].draw(screen)
			g.codonFont.drawFont(screen, protein[0].codon, protein[0].rect.pos.x, protein[0].rect.pos.y, color.Black)
			g.codonFont.drawFont(screen, protein[1].codon, protein[1].rect.pos.x, protein[1].rect.pos.y, color.Black)
			g.codonFont.drawFont(screen, protein[2].codon, protein[2].rect.pos.x, protein[2].rect.pos.y, color.Black)
			g.codonFont.drawFont(screen, protein[3].codon, protein[3].rect.pos.x, protein[3].rect.pos.y, color.Black)
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
