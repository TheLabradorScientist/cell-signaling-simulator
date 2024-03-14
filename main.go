package main

import (
	"bytes"
	"fmt"
	"image/color"
	_ "image/png"

	"log"
	"math/rand"
	"os"
	"path/filepath"

	//"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	//"github.com/hajimehoshi/ebiten/ebitenutil"
	//"github.com/hajimehoshi/ebiten/v2/audio"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	baseScreenWidth, baseScreenHeight = 1250, 750
)

var (
	screenWidth, screenHeight = 1250, 750
	maxWidth, maxHeight       = ebiten.ScreenSizeInFullscreen()
	widthRatio                = float64(maxWidth / baseScreenWidth)
	heightRatio               = float64(maxHeight / baseScreenHeight)
)

var (
	audioContext *audio.Context
)

var (

	// GENERAL SPRITES
	err               error
	scene             string = "Main Menu"
	otherToMenuButton Button
	audioPlayer       *audio.Player
	template          = [5]string{}
	reset             bool
	infoButton        InfoPage
	info              string
	state_array       []GUI
	defaultFont       Font
	codonFont         Font

	// MENU SPRITES
	/* 	protoStartBg StillImage
	   	startBg      Parallax
	   	startP1      Parallax
	   	startP2      Parallax
	   	startP3      Parallax
	   	startP4      Parallax
	   	fixedStart   StillImage
	   	playbutton   Button
	   	aboutButton  Button
	   	levSelButton Button
	   	volButton    Button */

	// ABOUT SPRITES
	//aboutBg           StillImage
	//aboutToMenuButton Button

	// LEVEL SELECT SPRITES
	levSelBg           StillImage
	levToPlasmaButton  Button
	levToCyto1Button   Button
	levToNucleusButton Button
	levToCyto2Button   Button
	levToMenuButton    Button

	// PLASMA SPRITES
	protoPlasmaBg  StillImage
	plasmaBg       Parallax
	plasmaMembrane Parallax
	//plasmaMembrane2 	Parallax
	// ^ Note: function to receptors that sets x and y to plasma membrane coord, -+ respective pos
	seedSignal = rand.Intn(4) + 1
	signal     Signal
	receptorA  Receptor
	receptorB  Receptor
	receptorC  Receptor
	receptorD  Receptor
	temp_tk1A  Kinase
	temp_tk1B  Kinase
	temp_tk1C  Kinase
	temp_tk1D  Kinase

	// CYTO 1 SPRITES
	protoCytoBg_1 StillImage
	cytoBg_1      Parallax
	cytoNuc_1     Parallax
	tk1           Kinase
	tk2           Kinase
	tfa           TFA

	// NUCLEUS SPRITES
	nucleusBg     StillImage
	rna           [5]Transcript
	dna           [5]Template
	currentFrag   = 0
	temp_tfa      TFA
	rnaPolymerase RNAPolymerase
	rightChoice   CodonChoice
	wrongChoice1  CodonChoice
	wrongChoice2  CodonChoice

	// Note to self: when updating DNA image, make the sprite like plasma membrane
	// So it can scroll to the left and show different codons, with bases as separate sprites
	// Also maybe try making RNA with theta and scrolling off to a upper-right diagonal

	// CYTO 2 SPRITES
	cytoBg_2   StillImage
	ribosome   Ribosome
	mrna       [5]Template
	protein    [5]Transcript
	mrna_ptr   int
	rightTrna  CodonChoice
	wrongTrna1 CodonChoice
	wrongTrna2 CodonChoice
)

var menuSprites []GUI
var aboutSprites []GUI
var levSelSprites []GUI
var plasmaSprites []GUI

//const sampleRate = 48000

type Game struct {
	switchedToPlasma      bool
	switchedToMenu        bool
	switchedToCyto1       bool
	switchedToNucleus     bool
	switchedToCyto2       bool
	switchedToAbout       bool
	switchedToLevelSelect bool
	stateMachine          *StateMachine
	//musicPlayer           *Player
	//musicPlayerCh         chan *Player
	//errCh                 chan error
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

func loadFont(font string) string {
	// Construct path to file
	fontpath := filepath.Join("Assets", "Fonts", font)
	// Open file
	file, err := os.Open(fontpath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "Error"
	}

	defer file.Close()

	return file.Name()
}

func loadMusic(music string) string {
	// Construct path to file
	musicpath := filepath.Join("Assets", "Music", music)
	// Open file
	file, err := os.Open(musicpath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "Error"
	}

	defer file.Close()

	return file.Name()
}

func (g *Game) init() {
	ebiten.SetWindowPosition(100, 0)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)
	defaultFont = newFont(loadFont("CourierPrime-Regular.ttf"), 32)
	codonFont = newFont(loadFont("BlackOpsOne-Regular.ttf"), 60)
	//maxWidth, maxHeight       = g.Layout(maxWidth, maxHeight)
	// Initialize audio context
	audioContext = audio.NewContext(44100)

	mp3Bytes, err := os.ReadFile(loadMusic("Signaling_of_the_Cell_MenuScreen.mp3"))
	if err != nil {
		log.Fatal(err)
	}
	mp3Stream, err := mp3.DecodeWithoutResampling(bytes.NewReader(mp3Bytes))
	if err != nil {
		log.Fatal(err)
	}
	audioPlayer, err = audioContext.NewPlayer(mp3Stream)
	if err != nil {
		log.Fatal(err)
	}

	var s_map = SceneConstructorMap{
		"menu": newMainMenu, "about": newAbout,
	}

	g.stateMachine = newStateMachine(s_map)
	g.stateMachine.changeState(g, "menu")

	infoButton = newInfoPage("infoButton.png", "infoPage.png", newRect(850, 0, 165, 165), "btn")
	otherToMenuButton = newButton("menuButton.png", newRect(1000, 0, 300, 200), ToMenu)

	/* 	protoStartBg = newStillImage("MenuBg.png", newRect(0, 0, 1250, 750))
	   	startBg = newParallax("StartBg.png", newRect(0, 0, 1250, 750), 5)
	   	startP1 = newParallax("parallax-Start2.png", newRect(0, 0, 1250, 750), 4)
	   	startP2 = newParallax("parallax-Start3.png", newRect(0, 0, 1250, 750), 3)
	   	startP3 = newParallax("parallax-Start4.png", newRect(0, 0, 1250, 750), 2)
	   	startP4 = newParallax("parallax-Start5.png", newRect(0, 0, 1250, 750), 1)
	   	volButton = newButton("volButtonOn.png", newRect(100, 100, 165, 165), SwitchVol)
	   	fixedStart = newStillImage("fixed-Start.png", newRect(0, 0, 1250, 750))
	   	playbutton = newButton("PlayButton.png", newRect(750, 100, 300, 200), ToPlasma)
	   	aboutButton = newButton("aboutButton.png", newRect(770, 260, 300, 200), ToAbout)
	   	levSelButton = newButton("levSelButton.png", newRect(700, 450, 300, 200), ToLevelSelect) */

	//aboutBg = newStillImage("AboutBg.png", newRect(0, 0, 1250, 750))

	levSelBg = newStillImage("levSelBg.png", newRect(0, 0, 1250, 750))

	protoPlasmaBg = newStillImage("PlasmaBg.png", newRect(0, 0, 1250, 750))
	plasmaBg = newParallax("ParallaxPlasma.png", newRect(100, 100, 1250, 750), 4)
	plasmaMembrane = newParallax("plasmaMembrane.png", newRect(100, 300, 1250, 750), 2)
	//plasmaMembrane2 = newParallax("plasmaMembrane.png", newRect(0, 200, 1250, 750), 2)

	protoCytoBg_1 = newStillImage("CytoBg1.png", newRect(0, 0, 1250, 750))
	cytoBg_1 = newParallax("ParallaxCyto1.png", newRect(100, 100, 1250, 750), 4)
	cytoNuc_1 = newParallax("ParallaxCyto1.5.png", newRect(100, 100, 1250, 750), 3)

	nucleusBg = newStillImage("NucleusBg.png", newRect(0, 0, 1250, 750))
	cytoBg_2 = newStillImage("CytoBg2.png", newRect(0, 0, 1250, 750))

	//aboutToMenuButton = newButton("menuButton.png", newRect(350, 450, 300, 200), ToMenu)
	levToPlasmaButton = newButton("levToPlasmaBtn.png", newRect(520, 110, 300, 180), ToPlasma)
	levToCyto1Button = newButton("levToCyto1Btn.png", newRect(820, 110, 300, 180), ToCyto1)
	levToNucleusButton = newButton("levToNucleusBtn.png", newRect(520, 285, 300, 180), ToNucleus)
	levToCyto2Button = newButton("levToCyto2Btn.png", newRect(820, 285, 300, 180), ToCyto2)
	levToMenuButton = newButton("menuButton.png", newRect(250, 190, 300, 200), ToMenu)

	switch seedSignal {
	case 1:
		signal = newSignal("signalA.png", newRect(500, 100, 100, 100))
		signal.signalType = "signalA"
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ACT"}
	case 2:
		signal = newSignal("signalB.png", newRect(500, 100, 100, 100))
		signal.signalType = "signalB"
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATT"}
	case 3:
		signal = newSignal("signalC.png", newRect(500, 100, 100, 100))
		signal.signalType = "signalC"
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATC"}
	case 4:
		signal = newSignal("signalD.png", newRect(500, 100, 100, 100))
		signal.signalType = "signalD"
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATT"}
		// PLACEHOLDER IN CASE WE DO NOT GET TIME TO CODE RANDOM CODONS
		//template = [5]string{"TAC", "GTC", "CGG", "ACA", "ACT"}
	}

	receptorA = newReceptor("inact_receptorA.png", newRect(50, 400, 100, 100), "receptorA")
	receptorB = newReceptor("inact_receptorB.png", newRect(350, 400, 100, 100), "receptorB")
	receptorC = newReceptor("inact_receptorC.png", newRect(650, 400, 100, 100), "receptorC")
	receptorD = newReceptor("inact_receptorD.png", newRect(950, 400, 100, 100), "receptorD")

	temp_tk1A = newKinase("inact_TK1.png", newRect(50, 600, 150, 150), "temp_tk1A")
	temp_tk1B = newKinase("inact_TK1.png", newRect(350, 600, 150, 150), "temp_tk1B")
	temp_tk1C = newKinase("inact_TK1.png", newRect(650, 600, 150, 150), "temp_tk1C")
	temp_tk1D = newKinase("inact_TK1.png", newRect(950, 600, 150, 150), "temp_tk1D")

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

	ribosome = newRibosome("ribosome.png", newRect(40, 300, 404, 367))

	mrna_ptr = 0

	rightTrna = newCodonChoice("codonButton.png", newRect(50, 150, 192, 111), translate(mrna[0].codon))
	wrongTrna1 = newCodonChoice("codonButton.png", newRect(350, 150, 192, 111), translate(randomRNACodon(rightTrna.bases)))
	wrongTrna2 = newCodonChoice("codonButton.png", newRect(650, 150, 192, 111), translate(randomRNACodon(rightTrna.bases)))

	//menuSprites = []GUI{protoStartBg, &startBg, &startP1, &startP2, &startP3, &startP4, fixedStart, playbutton, aboutButton, levSelButton, volButton}
	//aboutSprites = []GUI{&aboutBg, &aboutToMenuButton}
	levSelSprites = []GUI{&levSelBg, &levToPlasmaButton, &levToCyto1Button, &levToNucleusButton, &levToCyto2Button, &levToMenuButton}
	plasmaSprites = []GUI{&protoPlasmaBg, &plasmaBg, &plasmaMembrane, &signal, &receptorA, &receptorB, &receptorC, &receptorD, &temp_tk1A, &temp_tk1B, &temp_tk1C, &temp_tk1D}

}

func (g *Game) Update() error {

	//g.stateMachine.update(g)

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		ebiten.SetFullscreen(false)
	}

	switch scene {
	case "Main Menu":
		g.stateMachine.update(g)
		//g.musicPlayer.Play()
		//ebiten.SetWindowTitle("Cell Signaling Pathway - Main Menu")
		/*for _, element := range menuSprites {
			//e := element.(type)
			switch element.(type) {
			case (Button):
				element.update(g)
			default:
				element.update()
			}
		} */
		/* 		startBg.update()
		   		startP1.update()
		   		startP2.update()
		   		startP3.update()
		   		startP4.update()
		   		aboutButton.update(g)
		   		playbutton.update(g)
		   		levSelButton.update(g)
		   		volButton.update(g) */
	case "About":
		g.stateMachine.update(g)
	case "Level Selection":
		ebiten.SetWindowTitle("Cell Signaling Pathway - Level Selection")
		levToPlasmaButton.update(g)
		levToCyto1Button.update(g)
		levToNucleusButton.update(g)
		levToCyto2Button.update(g)
		levToMenuButton.update(g)
	case "Signal Reception":
		ebiten.SetWindowTitle("Cell Signaling Pathway - Signal Reception")
		plasmaBg.update()
		plasmaMembrane.update()
		//plasmaMembrane2.update()
		signal.update()
		receptorA.update()
		receptorB.update()
		receptorC.update()
		receptorD.update()
		temp_tk1A.update(temp_tk1B.rect)
		temp_tk1B.update(temp_tk1C.rect)
		temp_tk1C.update(temp_tk1D.rect)
		temp_tk1D.update(temp_tk1A.rect)
		otherToMenuButton.update(g)
		infoButton.update()
		info = updateInfo()
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
		if temp_tk1A.rect.pos.y >= screenHeight || temp_tk1B.rect.pos.y >= screenHeight || temp_tk1C.rect.pos.y >= screenHeight || temp_tk1D.rect.pos.y >= screenHeight {
			ToCyto1(g)
		}

	case "Signal Transduction":
		ebiten.SetWindowTitle("Cell Signaling Pathway - Signal Transduction")
		otherToMenuButton.update(g)
		cytoBg_1.update()
		cytoNuc_1.update()
		tk1.activate()
		tk1.update(tk2.rect)
		tk2.update(tfa.rect)
		tfa.update()
		infoButton.update()
		info = updateInfo()

		if tk1.is_clicked_on {
			tk2.activate()
			tk1.is_clicked_on = false
		}
		if tk2.is_clicked_on {
			tfa.activate()
			tk2.is_clicked_on = false
		}
		if tfa.rect.pos.y > screenHeight {
			ToNucleus(g)
		}

	case "Transcription":
		ebiten.SetWindowTitle("Cell Signaling Pathway - Transcription")
		otherToMenuButton.update(g)
		temp_tfa.activate()
		temp_tfa.update()
		rnaPolymerase.update(temp_tfa.rect.pos.y)
		infoButton.update()
		info = updateInfo()

		if reset {
			rightChoice.bases = transcribe(dna[currentFrag].codon)
			wrongChoice1.bases = randomRNACodon(rightChoice.bases)
			wrongChoice2.bases = randomRNACodon(rightChoice.bases)
			reset = false
		}
		//fmt.Printf("%t\n", dna[currentFrag].is_complete)
		dna[currentFrag].is_complete = rightChoice.update1(dna[currentFrag].codon)
		//fmt.Printf("%t\n", dna[currentFrag].is_complete)
		if dna[currentFrag].is_complete {
			nextDNACodon(g)
		}

	case "Translation":
		ebiten.SetWindowTitle("Cell Signaling Pathway - Translation")
		otherToMenuButton.update(g)
		infoButton.update()
		info = updateInfo()

		if reset {
			rightTrna.bases = translate(mrna[mrna_ptr].codon)
			wrongTrna1.bases = translate(randomRNACodon(rightTrna.bases))
			wrongTrna2.bases = translate(randomRNACodon(rightTrna.bases))
			reset = false
		}
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			mrna[mrna_ptr].is_complete = rightTrna.update2(mrna[mrna_ptr].codon)
		}

		if ribosome.update_movement() {
			nextMRNACodon(g)
		} else {
			ribosome.update_movement()
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	if ebiten.IsFullscreen() {
		// Use this if statement to set sizes of graphics to fullscreen scale, else normal scale.
		if screenWidth != maxWidth && screenHeight != maxHeight {
			screenWidth, screenHeight = maxWidth, maxHeight
			g.stateMachine.Scale(screen)
			defaultFont = newFont(loadFont("CourierPrime-Regular.ttf"), 32*int(heightRatio))
			codonFont = newFont(loadFont("BlackOpsOne-Regular.ttf"), 60*int(heightRatio))
		}
	} else {
		if screenWidth != baseScreenWidth && screenHeight != baseScreenHeight {
			screenWidth, screenHeight = baseScreenWidth, baseScreenHeight
			g.stateMachine.Scale(screen)
			defaultFont = newFont(loadFont("CourierPrime-Regular.ttf"), 32)
			codonFont = newFont(loadFont("BlackOpsOne-Regular.ttf"), 60)
		}
	}

	//g.stateMachine.draw(g, screen)
	switch scene {
	case "Main Menu":
		g.stateMachine.draw(g, screen)
		//for _, element := range menuSprites {
		//	element.draw(screen)
		//}
		/* 		protoStartBg.draw(screen)
		   		startBg.draw(screen)
		   		startP1.draw(screen)
		   		startP2.draw(screen)
		   		startP3.draw(screen)
		   		startP4.draw(screen)
		   		fixedStart.draw(screen)
		   		playbutton.draw(screen)
		   		aboutButton.draw(screen)
		   		levSelButton.draw(screen)
		   		volButton.draw(screen) */

		if g.switchedToPlasma {
			scene = "Signal Reception"
		}
		if g.switchedToLevelSelect {
			scene = "Level Selection"
		}
		if g.switchedToAbout {
			scene = "About"
		}
	case "About":
		g.stateMachine.draw(g, screen)
		if g.switchedToMenu {
			scene = "Main Menu"
		}

	case "Level Selection":
		levSelBg.draw(screen)
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
		protoPlasmaBg.draw(screen)
		plasmaBg.draw(screen)
		plasmaMembrane.draw(screen)
		receptorA.draw(screen)
		receptorB.draw(screen)
		receptorC.draw(screen)
		receptorD.draw(screen)
		//plasmaMembrane2.draw(screen)
		signal.draw(screen)
		temp_tk1A.draw(screen)
		temp_tk1B.draw(screen)
		temp_tk1C.draw(screen)
		temp_tk1D.draw(screen)
		otherToMenuButton.draw(screen)
		m := "WELCOME TO THE PLASMA MEMBRANE!"
		m += "\nDrag the signal to the matching receptor\nto enter the cell!"
		Pink := color.RGBA{220, 100, 100, 50}
		defaultFont.drawFont(screen, m, 100, 50, color.RGBA(Pink))
		if signal.is_dragged {
			signal.draw(screen)
		}
		infoButton.draw(screen)
		if g.switchedToMenu {
			scene = "Main Menu"
		}
		if g.switchedToCyto1 {
			scene = "Signal Transduction"
		}
	case "Signal Transduction":
		protoCytoBg_1.draw(screen)
		cytoBg_1.draw(screen)
		cytoNuc_1.draw(screen)
		tk1.draw(screen)
		tk2.draw(screen)
		tfa.draw(screen)
		defaultFont.drawFont(screen, "WELCOME TO THE CYTOPLASM! \n Click when each kinase overlaps to follow \n the phosphorylation cascade!!", 100, 50, color.Black)
		infoButton.draw(screen)
		otherToMenuButton.draw(screen)
		if g.switchedToNucleus {
			scene = "Transcription"
		}
		if g.switchedToMenu {
			scene = "Main Menu"
		}
	case "Transcription":
		nucleusBg.draw(screen)
		for x := 0; x < 5; x++ {
			rna[x].draw(screen)
		}
		dna[0].draw(screen)

		defaultFont.drawFont(screen, "WELCOME TO THE NUCLEUS! \n Match each codon on the DNA template to the corresponding RNA \n codon to transcribe a new mRNA molecule!!!", 100, 50, color.White)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, 300)
		rnaPolymerase.draw(screen)
		temp_tfa.draw(screen)
		//codonFont.drawFont(screen, strings.Join(template[0:5], ""), dna[currentFrag].rect.pos.x+300, dna[currentFrag].rect.pos.y, color.Black)
		if currentFrag != -1 {
			codonFont.drawFont(screen, dna[currentFrag].codon, dna[0].rect.pos.x+500, dna[0].rect.pos.y, color.Black)
		}
		rightChoice.draw(screen)
		wrongChoice1.draw(screen)
		wrongChoice2.draw(screen)
		codonFont.drawFont(screen, rightChoice.bases, rightChoice.rect.pos.x+25, rightChoice.rect.pos.y+90, color.Black)
		codonFont.drawFont(screen, wrongChoice1.bases, wrongChoice1.rect.pos.x+25, wrongChoice1.rect.pos.y+90, color.Black)
		codonFont.drawFont(screen, wrongChoice2.bases, wrongChoice2.rect.pos.x+25, wrongChoice2.rect.pos.y+90, color.Black)
		switch currentFrag {
		case 1:
			codonFont.drawFont(screen, rna[0].codon, rna[0].rect.pos.x+500, rna[0].rect.pos.y+140, color.Black)
		case 2:
			codonFont.drawFont(screen, rna[0].codon, rna[0].rect.pos.x+500, rna[0].rect.pos.y+140, color.Black)
			codonFont.drawFont(screen, rna[1].codon, rna[0].rect.pos.x+650, rna[0].rect.pos.y+140, color.Black)
		case 3:
			codonFont.drawFont(screen, rna[0].codon, rna[0].rect.pos.x+500, rna[0].rect.pos.y+140, color.Black)
			codonFont.drawFont(screen, rna[1].codon, rna[0].rect.pos.x+650, rna[0].rect.pos.y+140, color.Black)
			codonFont.drawFont(screen, rna[2].codon, rna[0].rect.pos.x+800, rna[0].rect.pos.y+140, color.Black)
		case 4:
			codonFont.drawFont(screen, rna[0].codon, rna[0].rect.pos.x+500, rna[0].rect.pos.y+140, color.Black)
			codonFont.drawFont(screen, rna[1].codon, rna[0].rect.pos.x+650, rna[0].rect.pos.y+140, color.Black)
			codonFont.drawFont(screen, rna[2].codon, rna[0].rect.pos.x+800, rna[0].rect.pos.y+140, color.Black)
			codonFont.drawFont(screen, rna[3].codon, rna[0].rect.pos.x+950, rna[0].rect.pos.y+140, color.Black)
		default:
			break
		}
		infoButton.draw(screen)
		otherToMenuButton.draw(screen)
		if g.switchedToCyto2 {
			scene = "Translation"
		}
		if g.switchedToMenu {
			scene = "Main Menu"
		}
	case "Translation":
		cytoBg_2.draw(screen)
		// for x := 0; x < 5; x++ {
		// 	protein[x].draw(screen)
		// }

		mrna[0].draw(screen)

		if mrna_ptr != -1 {
			codonFont.drawFont(screen, mrna[mrna_ptr].codon, mrna[0].rect.pos.x+500, mrna[0].rect.pos.y+200, color.Black)
		}

		rightTrna.draw(screen)
		wrongTrna1.draw(screen)
		wrongTrna2.draw(screen)

		codonFont.drawFont(screen, rightTrna.bases, rightTrna.rect.pos.x+25, rightTrna.rect.pos.y+90, color.Black)
		codonFont.drawFont(screen, wrongTrna1.bases, wrongTrna1.rect.pos.x+25, wrongTrna1.rect.pos.y+90, color.Black)
		codonFont.drawFont(screen, wrongTrna2.bases, wrongTrna2.rect.pos.x+25, wrongTrna2.rect.pos.y+90, color.Black)

		defaultFont.drawFont(screen, "FINALLY, BACK TO THE CYTOPLASM! \n Match each codon from your mRNA template \n to its corresponding amino acid to synthesize your protein!!!!", 100, 50, color.Black)

		if g.switchedToMenu {
			scene = "Main Menu"
		}

		switch mrna_ptr {
		case 0:
			protein[0].draw(screen)
			codonFont.drawFont(screen, protein[0].codon, protein[0].rect.pos.x, protein[0].rect.pos.y, color.Black)
		case 1:
			protein[0].draw(screen)
			protein[1].draw(screen)
			codonFont.drawFont(screen, protein[0].codon, protein[0].rect.pos.x, protein[0].rect.pos.y, color.Black)
			codonFont.drawFont(screen, protein[1].codon, protein[1].rect.pos.x, protein[1].rect.pos.y, color.Black)
		case 2:
			protein[0].draw(screen)
			protein[1].draw(screen)
			protein[2].draw(screen)
			codonFont.drawFont(screen, protein[0].codon, protein[0].rect.pos.x, protein[0].rect.pos.y, color.Black)
			codonFont.drawFont(screen, protein[1].codon, protein[1].rect.pos.x, protein[1].rect.pos.y, color.Black)
			codonFont.drawFont(screen, protein[2].codon, protein[2].rect.pos.x, protein[2].rect.pos.y, color.Black)
		case 3:
			protein[0].draw(screen)
			protein[1].draw(screen)
			protein[2].draw(screen)
			protein[3].draw(screen)
			codonFont.drawFont(screen, protein[0].codon, protein[0].rect.pos.x, protein[0].rect.pos.y, color.Black)
			codonFont.drawFont(screen, protein[1].codon, protein[1].rect.pos.x, protein[1].rect.pos.y, color.Black)
			codonFont.drawFont(screen, protein[2].codon, protein[2].rect.pos.x, protein[2].rect.pos.y, color.Black)
			codonFont.drawFont(screen, protein[3].codon, protein[3].rect.pos.x, protein[3].rect.pos.y, color.Black)
		case 4:
			protein[0].draw(screen)
			protein[1].draw(screen)
			protein[2].draw(screen)
			protein[3].draw(screen)
			codonFont.drawFont(screen, protein[0].codon, protein[0].rect.pos.x, protein[0].rect.pos.y, color.Black)
			codonFont.drawFont(screen, protein[1].codon, protein[1].rect.pos.x, protein[1].rect.pos.y, color.Black)
			codonFont.drawFont(screen, protein[2].codon, protein[2].rect.pos.x, protein[2].rect.pos.y, color.Black)
			codonFont.drawFont(screen, protein[3].codon, protein[3].rect.pos.x, protein[3].rect.pos.y, color.Black)
		default:
			break
		}

		ribosome.draw(screen)
		infoButton.draw(screen)

		otherToMenuButton.draw(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	game := &Game{}
	game.init()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

}
