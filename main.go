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
	seedSignal        int


	// CYTO 2 SPRITES
	cytoBg_2   StillImage
	ribosome   Ribosome
	mrna       [5]Template
	protein    [5]Transcript
	mrna_ptr   int
	rightTrna  tRNA
	wrongTrna1 tRNA
	wrongTrna2 tRNA
)

var menuSprites []GUI
var aboutSprites []GUI
var levSelSprites []GUI
var plasmaSprites []GUI
var cyto1Sprites []GUI
var nucleusSprites []GUI
var cyto2Sprites []GUI

type Game struct {
	switchedScene         bool
	stateMachine          *StateMachine
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
	//	maxWidth, maxHeight       = g.Layout(maxWidth, maxHeight)
	// 	Initialize audio context
	audioContext = audio.NewContext(44100)
	g.switchedScene = true

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
		"Main Menu": newMainMenu, "About": newAbout, "Level Selection": newLevelSelection,
		"Signal Reception": newReceptionLevel, "Signal Transduction": newTransductionLevel,
		"Transcription": newTranscriptionLevel, "Translation": newTranslationLevel,
	}

	g.stateMachine = newStateMachine(s_map)

	g.stateMachine.changeState(g, "Main Menu")

	seedSignal = rand.Intn(4) + 1
	infoButton = newInfoPage("infoButton.png", "infoPage.png", newRect(850, 0, 165, 165), "btn")
	otherToMenuButton = newButton("menuButton.png", newRect(1000, 0, 300, 200), ToMenu)

	cytoBg_2 = newStillImage("CytoBg2.png", newRect(0, 0, 1250, 750))

	switch seedSignal {
	case 1:
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ACT"}
	case 2:
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATT"}
	case 3:
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATC"}
	case 4:
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATT"}
		// PLACEHOLDER IN CASE WE DO NOT GET TIME TO CODE RANDOM CODONS
		//template = [5]string{"TAC", "GTC", "CGG", "ACA", "ACT"}
	}

	for x := 0; x < 5; x++ {
		dna[x] = newTemplate("DNA.png", newRect(-50+200*x, 500, 150, 150), template[x], x)}
	for x := 0; x < 5; x++ {
		rna[x] = newTranscript("RNA.png", newRect(0, 200, 150, 150), transcribe(template[x]))
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

	reset = false

	ribosome = newRibosome("ribosome.png", newRect(40, 300, 404, 367))

	mrna_ptr = 0

	rightTrna = newTRNA("codonButton.png", newRect(50, 150, 192, 111), translate(mrna[0].codon))
	wrongTrna1 = newTRNA("codonButton.png", newRect(350, 150, 192, 111), translate(randomRNACodon(rightTrna.bases)))
	wrongTrna2 = newTRNA("codonButton.png", newRect(650, 150, 192, 111), translate(randomRNACodon(rightTrna.bases)))

}

func (g *Game) Update() error {

	ebiten.SetWindowTitle("CSPS - " + scene)
	//g.stateMachine.update(g)

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		ebiten.SetFullscreen(false)
	}

	switch scene {
	case "Main Menu":
		g.stateMachine.update(g)
	case "About":
		g.stateMachine.update(g)
	case "Level Selection":
		g.stateMachine.update(g)
	case "Signal Reception":
		g.stateMachine.update(g)
	case "Signal Transduction":
		g.stateMachine.update(g)
	case "Transcription":
		g.stateMachine.update(g)

	case "Translation":
		otherToMenuButton.update(g)
		infoButton.update()
		info = updateInfo()

		curr := &mrna[mrna_ptr]

		if reset {
			rightTrna.bases = translate(curr.codon)
			wrongTrna1.bases = translate(randomRNACodon(rightTrna.bases))
			wrongTrna2.bases = translate(randomRNACodon(rightTrna.bases))
			reset = false
		}
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			rightTrna.update(curr)
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
		if screenWidth == baseScreenWidth && screenHeight == baseScreenHeight {
			screenWidth, screenHeight = ebiten.ScreenSizeInFullscreen()
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

	if g.switchedScene {
		g.stateMachine.Scale(screen)
		g.switchedScene = false
	}

	//g.stateMachine.draw(g, screen)
	switch scene {
	case "Main Menu":
		g.stateMachine.draw(g, screen)

	case "About":
		g.stateMachine.draw(g, screen)

	case "Level Selection":
		g.stateMachine.draw(g, screen)

	case "Signal Reception":
		g.stateMachine.draw(g, screen)

	case "Signal Transduction":
		g.stateMachine.draw(g, screen)
		
	case "Transcription":
		g.stateMachine.draw(g, screen)

	case "Translation":
		cytoBg_2.draw(screen)

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
	if ebiten.IsFullscreen() {
		return outsideWidth, outsideHeight
	} else {
		return baseScreenWidth, baseScreenHeight
	}
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
