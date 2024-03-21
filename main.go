package main

import (
	"bytes"
	"fmt"
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
	// GLOBAL VARIABLES
	screenWidth, screenHeight = 1250, 750
	maxWidth, maxHeight       = ebiten.ScreenSizeInFullscreen()
	widthRatio                = float64(maxWidth / baseScreenWidth)
	heightRatio               = float64(maxHeight / baseScreenHeight)

	audioContext *audio.Context

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


	menuSprites []GUI
	aboutSprites []GUI
	levSelSprites []GUI
	plasmaSprites []GUI
	cyto1Sprites []GUI
	nucleusSprites []GUI
	cyto2Sprites []GUI
)

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
	ToMenu(g)

	seedSignal = rand.Intn(4) + 1
	infoButton = newInfoPage("infoButton.png", "infoPage.png", newRect(850, 0, 165, 165), "btn")
	otherToMenuButton = newButton("menuButton.png", newRect(1000, 0, 300, 200), ToMenu)

	switch seedSignal {
	case 1:
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ACT"}
	case 2:
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATT"}
	case 3:
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATC"}
	case 4:
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATT"}
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
}

func (g *Game) Update() error {

	ebiten.SetWindowTitle("CSPS - " + scene)

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		ebiten.SetFullscreen(false)
	}

	g.stateMachine.update(g)

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

	g.stateMachine.draw(g, screen)

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
